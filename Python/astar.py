import time
from math import sqrt
from typing import List, Tuple
import numpy as np
import random

# 1 for manhattan, 0 for euclidean
HEURISTIC = 1
DIAGONALS = False

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
            if i == 0:
                self.grid[elem[1]][elem[0]] = 3
            elif i == len(path) - 1:
                self.grid[elem[1]][elem[0]] = 3
            else:
                self.grid[elem[1]][elem[0]] = 2
            i += 1
        self.path_added = True

    def plot_path(self):
        if self.path_added:
            print(self.grid)
        else:
            print("No path added, can't plot path")

    def get_random_available_position(self) -> Tuple[int]:
        while 1:
            random_row = random.randint(0, self.height - 1)
            random_col = random.randint(0, self.length - 1)

            if self.grid[random_row][random_col] != 1:
                break
                
        return((random_row, random_col))
                
    def change_grid(self, position: Tuple[int], value: int):
        self.grid[position[1]][position[0]] = value
        
    def __generate_walls(self):
        for x in range(self.height):
            for y in range(self.length):
                if random.random() < self.p_walls:
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

    def __init__(self, pos_x: int, pos_y: int, goal_x: int, goal_y: int, g_cost: float, parent):
        self.pos_x = pos_x
        self.pos_y = pos_y
        self.pos = (pos_y, pos_x)
        self.goal_x = goal_x
        self.goal_y = goal_y
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
        self.start = Node(start[1], start[0], goal[1], goal[0], 0, None)
        self.target = Node(goal[1], goal[0], goal[1], goal[0], 99999, None)
        self.world = world

        self.open_nodes = [self.start]
        self.closed_nodes = []

        self.reached = False

    def search(self) -> List[Tuple[int]]:
        while self.open_nodes:
            # Open Nodes are sorted using __lt__
            self.open_nodes.sort()
            current_node = self.open_nodes.pop(0)

            if current_node.pos == self.target.pos:
                self.reached = True
                return self.retrace_path(current_node)

            self.closed_nodes.append(current_node)
            successors = self.get_successors(current_node)

            for child_node in successors:
                if child_node in self.closed_nodes:
                    continue

                if child_node not in self.open_nodes:
                    self.open_nodes.append(child_node)
                else:
                    # retrieve the best current path
                    better_node = self.open_nodes.pop(self.open_nodes.index(child_node))

                    if child_node.g_cost < better_node.g_cost:
                        self.open_nodes.append(child_node)
                    else:
                        self.open_nodes.append(better_node)

        if not (self.reached):
            return [(self.start.pos)]

    def get_successors(self, parent: Node) -> List[Node]:
        """
        Returns a list of successors (both in the world and free spaces)
        """
        successors = []
        for action in self.world.delta:
            pos_x = parent.pos_x + action[1]
            pos_y = parent.pos_y + action[0]
            if not (0 <= pos_x <= len(self.world.grid[0]) - 1 and 0 <= pos_y <= len(self.world.grid) - 1):
                continue

            if self.world.grid[pos_y][pos_x] != 0:
                continue

            successors.append(
                Node(
                    pos_x,
                    pos_y,
                    self.target.pos_y,
                    self.target.pos_x,
                    parent.g_cost + action[2],
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
    import doctest

    doctest.testmod()

    # w = World(10, 10, 0.25)
    # print(w.grid)
    
    # start = w.get_random_available_position()
    # goal = w.get_random_available_position()

    # print(start, goal)

    # astar = AStar(start, goal, w)
    # path = astar.search()

    # print(path)

    # w.add_path(path)
    # w.plot_path()

