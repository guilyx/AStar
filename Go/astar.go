package main

import (
	"fmt"
	"os"
	"math/rand"
	"math"
	"time"
	color "github.com/fatih/color"
)

const DIAGONALS = false
const HEURISTIC = 1

var Costs[]float64
var Actions[][]int

type Position struct {
	x int
	y int
}

type Environment struct {
	grid [][]uint8
	height int
	length int
	wallPercentage float64
}

type Path struct {
	pos []Position
	cost []float64
}

type Node struct {
	pos      Position
	fCost    float64
	gCost    float64
	hCost    float64
	isOpen   bool
	isClosed bool
	parent   *Node
}

func DeclareActions() {
	var tempActions[][]int
	for dx := -1 ; dx <= 1 ; dx++ {
		for dy := -1; dy <= 1 ; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if dx == 0 || dy == 0 {
				Actions = append(Actions, []int{dy, dx})
				Costs = append(Costs, 1.0)
			} else {
				tempActions = append(tempActions, []int{dy, dx})
			}
		}
	}
	for _, actions := range tempActions {
		Actions = append(Actions, actions)
		Costs = append(Costs, math.Sqrt(2))
	}
}

func GenerateEnvironment(height int, length int, wp float64) *Environment {
	g := make([][]uint8, height)
	for i := range g {
		g[i] = make([]uint8, length)
	}

	for j := range g {
		for k := range g[j] {
			if rand.Float64() < wp {
				g[j][k] = 1
			}
			if j == height-1 || j == 0 {
				g[j][k] = 1
			} 
			if k == length-1 || k == 0 {
				g[j][k] = 1
			}
		}
	}

	m := Environment {
		grid: g,
		height: height,
		length: length,
		wallPercentage: wp,
	}
	
	return &m
}

func PrintEnvironment(m *Environment) {
	if m == nil {
		color.Red("Empty environment passed for printing.\n")
		os.Exit(1)
	}
	for k := range m.grid {
		for j := range m.grid[k] {
			switch m.grid[k][j] {
			case 0:
				color.Set(color.FgWhite)
				fmt.Printf("%d", 0)
				color.Unset()
			case 1:
				color.Set(color.FgRed)
				fmt.Printf("%d", 1)
				color.Unset()
			case 2:
				color.Set(color.FgCyan)
				fmt.Printf("X")
				color.Unset()
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

func AddPath(p *Path, m *Environment) {
	if p != nil {
		for _, val := range p.pos {
			AddElemToEnvironment(m, &val)
		}
	}
}

func AddElemToEnvironment(m *Environment, p *Position) {
	m.grid[p.y][p.x] = 2
}

func FindAvailablePosition(m *Environment) *Position {
	l := m.length - 2
	h := m.height - 2
	x := rand.Intn(l) + 1
	y := rand.Intn(h) + 1
	val := m.grid[h][l]

	for val != 0 {
		x = rand.Intn(l) + 1
		y = rand.Intn(h) + 1
		val = m.grid[h][l]
	}

	pos := createPosition(x, y)

	return pos
}

func calculateHeuristic(currentPos *Position, goal *Position) float64 {
	var h float64
	dy := goal.y - currentPos.y
	dx := goal.x - currentPos.x
	if HEURISTIC == 0 {
		h = math.Abs(float64(dy)) + math.Abs(float64(dx))
	} else if HEURISTIC == 1 {
		p1 := math.Pow(float64(dy), 2)
		p2 := math.Pow(float64(dx), 2)
		h = math.Sqrt(p1 + p2)
	} else {
		return 0.0
	}
	return h
}

func createPosition(x int, y int) *Position {
	p := Position {
		x: x,
		y: y,
	}
	return &p
}

func createNode(nodes map[Position]*Node, pos *Position, goal *Position, open bool, closed bool, g float64, parent *Node) {
	nodeTemp := Node {
		pos: *pos,
		isOpen: open,
		isClosed: closed,
		gCost: g,
		hCost: calculateHeuristic(pos, goal),
		parent: parent,
	}
	nodeTemp.fCost = nodeTemp.gCost + nodeTemp.hCost
	nodes[*pos] = &nodeTemp
}

func retrievePath(lastNode *Node) *Path{
	path := Path {}
	for lastNode != nil {
		path.pos = append(path.pos, lastNode.pos)
		path.cost = append(path.cost, lastNode.gCost)
		lastNode = lastNode.parent
	}
	return &path
}

func getNeighbours(n *Node, goal *Position, env *Environment) []Node {
	children := []Node{}
	//fmt.Println("Node : ", *n)
	for i, action := range Actions {
		if !( (0 <= n.pos.y + action[0] && n.pos.y + action[0] < env.length) && 
		      (0 <= n.pos.x + action[1] && n.pos.x + action[1] < env.height) ) {
			continue
		}
		if env.grid[n.pos.y + action[0]][n.pos.x + action[1]] == 1 {
			continue
		}
		
		if DIAGONALS == false && action[0] != 0 && action[1] != 0 {
			continue
		}
		position := Position {
			x: n.pos.x + action[1],
			y: n.pos.y + action[0],
		}
		child := Node {
			pos: position,
			gCost: n.gCost + Costs[i],
			hCost: calculateHeuristic(&position, goal),
			parent: n,
		}
		child.fCost = child.gCost + child.hCost
		//fmt.Println("Child : ", child)
		children = append(children, child)
	}
	return children
}

func findLowestFCost(nodes map[Position]*Node) *Node {
	var currentNode *Node

	for _, rdm := range nodes {
		if rdm.isOpen == true {
			currentNode = rdm
			break
		}
	}

	for _, node := range nodes {
		if (node.isOpen == true && node.fCost <= currentNode.fCost) {
			currentNode = node
		}
	}

	if currentNode != nil {
		return currentNode
	}

	fmt.Printf("Current node is nil !!\n")
	os.Exit(-1)
	return nil
}

func Search(start *Position, goal *Position, m *Environment) *Path {
	// Initialize A*
	nodes := make(map[Position]*Node)
	createNode(nodes, start, goal, true, false, 0.0, nil)
	// createNode(nodes, goal, goal, false, false, 999999, nil)

	currentPosition := Position {
		x: start.x,
		y: start.y,
	}

	currentNode := nodes[*start]

	// While there is an open node
	for currentNode != nil {
		// Extract lowest F in Open Nodes
		currentNode = findLowestFCost(nodes)
		currentPosition = currentNode.pos
		
		nodes[currentPosition].isOpen = false
		nodes[currentPosition].isClosed = true

		if currentPosition == *goal {
			p := retrievePath(nodes[currentPosition])
			fmt.Printf("Found path.\n")
			return p
		}

		children := getNeighbours(nodes[currentPosition], goal, m)

		for _, child := range children {
			if previousNode, ok := nodes[child.pos]; ok {
				if previousNode.isClosed {
					continue
				}
				if previousNode.isOpen {
					if previousNode.gCost > child.gCost {
						nodes[child.pos] = &child 
						nodes[child.pos].isOpen = true
						nodes[child.pos].isClosed = false
					} else {
						nodes[child.pos].isOpen = true
						nodes[child.pos].isClosed = false
					}
				} else {
					nodes[child.pos] = &child 
					nodes[child.pos].isOpen = true
					nodes[child.pos].isClosed = false
				}
			} else {
				nodes[child.pos] = &child
				nodes[child.pos].isOpen = true
				nodes[child.pos].isClosed = false
			}
		}
	}

	fmt.Printf("Path not found.\n")
	return nil
}

func main() {
	DeclareActions()
	rand.Seed(time.Now().UnixNano())
	height := 10
	length := 10
	var wp float64 = .1
	env := GenerateEnvironment(height, length, wp)

	start := FindAvailablePosition(env)
	goal := FindAvailablePosition(env)

	fmt.Printf("Start [%d ; %d]\nGoal [%d ; %d]\n", start.x, start.y, goal.x, goal.y)

	AddElemToEnvironment(env, start)
	AddElemToEnvironment(env, goal)

	PrintEnvironment(env)

	path := Search(start, goal, env)

	if path != nil {
		AddPath(path, env)
	}

	PrintEnvironment(env)
}