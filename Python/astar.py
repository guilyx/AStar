import time
from math import sqrt
from typing import List, Tuple
from sys import stdout
import numpy as np
import random
from termcolor import colored

# 1 for manhattan, 0 for euclidean
HEURISTIC = 0
DIAGONALS = True
TIME = True
PLOT = True

class World:
    def __init__(self, length: int, height: int, p_walls: float):
        self.length = length
        self.height = height
        self.p_walls = p_walls
        self.path_added = False

        self.grid = np.zeros((height, length), int)
        self.__generate_walls()

        if not DIAGONALS:
            # up, left, down, right
            self.delta = [[-1, 0, 1], 
                          [0, -1, 1], 
                          [1, 0, 1], 
                          [0, 1, 1]]  
        else:
            # up, left, down, right 
            # upleft, upright, downleft, downright
            self.delta = [[-1, 0, 1], 
                          [0, -1, 1], 
                          [1, 0, 1], 
                          [0, 1, 1],
                          [-1, -1, sqrt(2)],
                          [-1, 1, sqrt(2)],
                          [1, -1, sqrt(2)],
                          [1, 1, sqrt(2)]]

    def add_path(self, path: List[Tuple[int]]):
        i = 0
        for elem in path:
            if i == 0 or i == len(path) - 1:
                self.grid[elem[0]][elem[1]] = 3
            else:
                self.grid[elem[0]][elem[1]] = 2
            i += 1
        self.path_added = True

    def plot_grid(self):
        for rows in self.grid:
            for elem in rows:
                if elem == 1:
                    stdout.write(colored('█', 'red'))
                elif elem == 0:
                    stdout.write(' ')
                elif elem == 2:
                    stdout.write(colored('¤', 'blue'))
                elif elem == 3:
                    stdout.write(colored('x', 'green'))
                else:
                    stdout.write('o', 'yellow')
            print()

    def get_random_available_position(self) -> Tuple[int]:
        i = 0
        while 1:
            i += 1
            random_row = random.randint(0, self.height - 1)
            random_col = random.randint(0, self.length - 1)

            if self.grid[random_row][random_col] != 1:
                break
            if i > self.length*self.height*100:
                print(colored("Couldn't find random available position.", "red"))
                return None
                
        return((random_row, random_col))

    def get_start_goal(self, pHeuristic) -> Tuple[Tuple[int]]:
        i = 0

        # Initialization
        start = self.get_random_available_position()
        goal = self.get_random_available_position()
        n = Node(start[1], start[0], goal[1], goal[0], 0, False, False, None)
        h = n.calculate_heuristic()

        # Set goal value
        nTarget = Node(0, 0, self.length-1, self.height-1, 0, False, False, None)
        hTarget = nTarget.calculate_heuristic()

        while h < hTarget*pHeuristic:
            i += 1
            start = self.get_random_available_position()
            goal = self.get_random_available_position()
            n = Node(start[1], start[0], goal[1], goal[0], 0, False, False, None)
            h = n.calculate_heuristic()
            if i > 100*self.length*self.height:
                print(colored("Couldn't find start and goal respecting the constraints.", "red"))
                return None
        return start, goal
                
    def change_grid(self, position: Tuple[int], value: int):
        self.grid[position[1]][position[0]] = value
        
    def __generate_walls(self):
        for x in range(self.height):
            for y in range(self.length):
                if random.random() < self.p_walls:
                    self.grid[x][y] = 1
                if x == 0 or x == self.height - 1:
                    self.grid[x][y] = 1
                if y == 0 or y == self.length - 1:
                    self.grid[x][y] = 1

class Node:
    """
    >>> k = Node(0, 0, 4, 3, 0, None)
    >>> k.calculate_heuristic()
    5.0
    >>> n = Node(1, 4, 3, 4, 2, None)
    >>> n.calculate_heuristic()
    2.0
    >>> l = [k, n]
    >>> n == l[0]
    False
    >>> l.sort()
    >>> n == l[0]
    True
    """

    def __init__(self, pos_x: int, pos_y: int, goal_x: int, goal_y: int, g_cost: float, is_closed:bool, is_open:bool, parent):
        self.pos_x = pos_x
        self.pos_y = pos_y
        self.pos = (pos_y, pos_x)
        self.goal_x = goal_x
        self.goal_y = goal_y
        self.is_closed = is_closed
        self.is_open = is_open
        self.g_cost = g_cost
        self.parent = parent
        self.h_cost = self.calculate_heuristic()
        self.f_cost = self.g_cost + self.h_cost

    def calculate_heuristic(self) -> float:
        """
        Heuristic for the A*
        """
        dy = self.pos_x - self.goal_x
        dx = self.pos_y - self.goal_y
        if HEURISTIC == 1:
            return abs(dx) + abs(dy)
        else:
            return sqrt(dy ** 2 + dx ** 2)

    def print(self):
        print("Node : <Pos : ", self.pos, " >")

    def __lt__(self, other) -> bool:
        return self.f_cost <= other.f_cost


class AStar:
    """
    >>> wd = World(10, 10, 0.2)
    >>> astar = AStar((0, 0), (len(wd.grid) - 1, len(wd.grid[0]) - 1), wd)
    >>> (astar.start.pos_y + wd.delta[3][0], astar.start.pos_x + wd.delta[3][1])
    (0, 1)
    >>> [x.pos for x in astar.get_successors(astar.start)]
    [(1, 0), (0, 1)]
    >>> (astar.start.pos_y + wd.delta[2][0], astar.start.pos_x + wd.delta[2][1])
    (1, 0)
    >>> astar.retrace_path(astar.start)
    [(0, 0)]
    >>> astar.search()  # doctest: +NORMALIZE_WHITESPACE
    [(0, 0), (1, 0), (2, 0), (2, 1), (2, 2), (2, 3), (3, 3),
     (4, 3), (4, 4), (5, 4), (5, 5), (6, 5), (6, 6)]
    """

    def __init__(self, start: Tuple[int], goal: Tuple[int], world: World):
        self.start = Node(start[1], start[0], goal[1], goal[0], 0, False, True, None)
        self.target = Node(goal[1], goal[0], goal[1], goal[0], 99999, False, False, None)
        self.world = world

        self.nodes = dict()
        self.nodes[start] = self.start

        self.reached = False

    def search(self) -> List[Tuple[int]]:
        while self.count_open_nodes() > 0:
            # Open Nodes are sorted using __lt__
            current_key = min([n for n in self.nodes if self.nodes[n].is_open], key=(lambda k: self.nodes[k].f_cost))
            current_node = self.nodes[current_key]

            if current_node.pos == self.target.pos:
                print(colored("Found path.", "green"))
                return self.retrace_path(current_node)

            current_node.is_closed = True
            current_node.is_open = False

            successors = self.get_successors(current_node)

            for child_node in successors:
                if child_node.pos in self.nodes:
                    if self.nodes[child_node.pos].is_closed:
                        continue
                    if not self.nodes[child_node.pos].is_open:
                        self.nodes[child_node.pos] = child_node
                    else:
                        if child_node.g_cost < self.nodes[child_node.pos].g_cost:
                            self.nodes[child_node.pos] = child_node
                else:
                    self.nodes[child_node.pos] = child_node
        print(colored("Path not found", "red"))
        return [self.start.pos]

    def count_open_nodes(self) -> int:
        i = 0
        for _, val in self.nodes.items():
            if val.is_open:
                i += 1
                break
        return i

    def get_successors(self, parent: Node) -> List[Node]:
        """
        Returns a list of successors (both in the world and free spaces)
        """
        successors = []
        for action in self.world.delta:
            pos_x = parent.pos_x + action[0]
            pos_y = parent.pos_y + action[1]
            if not (0 <= pos_x < self.world.length and 0 <= pos_y < self.world.height):
                continue

            if self.world.grid[pos_y][pos_x] != 0:
                continue

            successors.append(
                Node(
                    pos_x,
                    pos_y,
                    self.target.pos_x,
                    self.target.pos_y,
                    parent.g_cost + action[2],
                    False,
                    True,
                    parent,
                )
            )
        return successors

    def retrace_path(self, node: Node) -> List[Tuple[int]]:
        """
        Retrace the path from parents to parents until start node
        """
        current_node = node
        path = []
        while current_node is not None:
            path.append((current_node.pos_y, current_node.pos_x))
            current_node = current_node.parent
        path.reverse()
        return path

if __name__ == "__main__":
    # all coordinates are given in format [y,x]
    # import doctest

    # doctest.testmod()

    w = World(70, 20, 0.2)
    w.plot_grid()
    
    # Generated
    start, goal = w.get_start_goal(0.8)

    # Hardcoded
    # start = (5, 5)
    # goal = (0, 9)
    
    print(start, goal)

    astar = AStar(start, goal, w)
    if TIME:
        st = time.process_time()
    path = astar.search()
    if TIME:
        print("Search() Execution Time : ", time.process_time() - st, "seconds")

    w.add_path(path)
    w.plot_grid()

