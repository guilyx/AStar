package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	color "github.com/fatih/color"
)

const DIAGONALS = true
const HEURISTIC = 0
const STEP_DISPLAY = false

var Costs []float64
var Actions [][]int

type Position struct {
	x int
	y int
}

type Environment struct {
	grid           [][]uint8
	height         int
	length         int
	wallPercentage float64
}

type Path struct {
	pos  []Position
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

func clearTerminal() {
	fmt.Printf("\033[H\033[2J")
}

func DeclareActions() {
	var tempActions [][]int
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
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

	m := Environment{
		grid:           g,
		height:         height,
		length:         length,
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
			case 3:
				color.Set(color.FgGreen)
				fmt.Printf("o")
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
			AddElemToEnvironment(m, &val, 2)
		}
	}
}

func AddElemToEnvironment(m *Environment, p *Position, val uint8) {
	m.grid[p.y][p.x] = val
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

func calculateHeuristic(currentPos Position, goal Position) float64 {
	dy := goal.y - currentPos.y
	dx := goal.x - currentPos.x

	if HEURISTIC == 0 {
		return(math.Abs(float64(dy)) + math.Abs(float64(dx)))
	}

	p1 := math.Pow(float64(dy), 2)
	p2 := math.Pow(float64(dx), 2)
	return(math.Sqrt(p1 + p2))
}

func createPosition(x int, y int) *Position {
	p := Position{
		x: x,
		y: y,
	}
	return &p
}

func createNode(nodes map[Position]*Node, pos *Position, goal *Position, open bool, closed bool, g float64, parent *Node) {
	nodeTemp := Node{
		pos:      *pos,
		isOpen:   open,
		isClosed: closed,
		gCost:    g,
		hCost:    calculateHeuristic(*pos, *goal),
		parent:   parent,
	}
	nodeTemp.fCost = nodeTemp.gCost + nodeTemp.hCost
	nodes[*pos] = &nodeTemp
}

func retrievePath(lastNode *Node) *Path {
	path := Path{}
	for lastNode != nil {
		path.pos = append(path.pos, lastNode.pos)
		path.cost = append(path.cost, lastNode.gCost)
		lastNode = lastNode.parent
	}
	return &path
}

func getNeighbours(n *Node, goal *Position, env *Environment) []Node {
	children := []Node{}
	for i, action := range Actions {
		if !((0 <= n.pos.y+action[0] && n.pos.y+action[0] < env.height) &&
			(0 <= n.pos.x+action[1] && n.pos.x+action[1] < env.length)) {
			continue
		}
		if env.grid[n.pos.y+action[0]][n.pos.x+action[1]] == 1 {
			continue
		}

		if DIAGONALS == false && action[0] != 0 && action[1] != 0 {
			continue
		}
		position := Position{
			x: n.pos.x + action[1],
			y: n.pos.y + action[0],
		}
		child := Node{
			pos:    position,
			gCost:  n.gCost + Costs[i],
			hCost:  calculateHeuristic(position, *goal),
			parent: n,
		}
		child.fCost = child.gCost + child.hCost
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
		if node.isOpen == true && node.fCost < currentNode.fCost {
			currentNode = node
		}
	}

	if currentNode != nil {
		return currentNode
	}

	fmt.Printf("Current node is nil !!\n")

	return nil
}

func countOpenNodes(nodes map[Position]*Node) int {
	var cnt int
	for _, val := range nodes {
		if val.isOpen == true {
			cnt += 1
		}
	}
	return cnt
}

func Search(start *Position, goal *Position, m *Environment) *Path {
	// Initialize A*
	nodes := make(map[Position]*Node)
	createNode(nodes, start, goal, true, false, 0.0, nil)
	// createNode(nodes, goal, goal, false, false, 999999, nil)

	currentPosition := Position{
		x: start.x,
		y: start.y,
	}

	for countOpenNodes(nodes) != 0 {
		// Extract lowest F in Open Nodes
		currentNode := findLowestFCost(nodes)
		currentPosition = currentNode.pos

		if currentPosition == *goal {
			p := retrievePath(nodes[currentPosition])
			fmt.Printf("Found path.\n")
			return p
		}

		if STEP_DISPLAY == true && currentPosition != *start && currentPosition != *goal{
			AddElemToEnvironment(m, &currentPosition, 3)
			fmt.Printf("Start [%d ; %d]\nGoal [%d ; %d]\n", start.x, start.y, goal.x, goal.y)
			PrintEnvironment(m)
			time.Sleep(500 * time.Millisecond)
			clearTerminal()
		}

		nodes[currentPosition].isOpen = false
		nodes[currentPosition].isClosed = true

		children := getNeighbours(nodes[currentPosition], goal, m)

		// if len(children) == 0 {
		// 	fmt.Printf("No children found\n")
		// }

		for _, child := range children {
			if previousNode, ok := nodes[child.pos]; ok {
				if previousNode.isClosed == true {
					continue
				}

				if previousNode.isOpen == false {
					child.isOpen = true
					nodes[child.pos] = &Node{
						pos: child.pos,
						fCost: child.fCost,
						gCost: child.gCost,
						hCost: child.hCost,
						isClosed: child.isClosed,
						isOpen: child.isOpen,
						parent: child.parent,
					}
				} else {
					if child.gCost < previousNode.gCost {
						child.isOpen = true
						nodes[child.pos] = &Node{
							pos: child.pos,
							fCost: child.fCost,
							gCost: child.gCost,
							hCost: child.hCost,
							isClosed: child.isClosed,
							isOpen: child.isOpen,
							parent: child.parent,
						}
					} else {
						previousNode.isOpen = true
					}
				}
			} else {
				child.isOpen = true
				nodes[child.pos] = &Node{
					pos: child.pos,
					fCost: child.fCost,
					gCost: child.gCost,
					hCost: child.hCost,
					isClosed: child.isClosed,
					isOpen: child.isOpen,
					parent: child.parent,
				}
			}
			// fmt.Println("Child : ", child, "---- Parent : ", *child.parent)
		}
		// fmt.Printf("--------------\n")
	}

	fmt.Printf("Path not found")
	return nil
}

func main() {
	DeclareActions()
	rand.Seed(time.Now().UnixNano())
	height := 20
	length := 50
	var wp float64 = .2
	env := GenerateEnvironment(height, length, wp)

	start := FindAvailablePosition(env)
	goal := FindAvailablePosition(env)

	fmt.Printf("Start [%d ; %d]\nGoal [%d ; %d]\n", start.x, start.y, goal.x, goal.y)

	AddElemToEnvironment(env, start, 2)
	AddElemToEnvironment(env, goal, 2)

	PrintEnvironment(env)

	path := Search(start, goal, env)

	if path != nil {
		fmt.Println(path)

		AddPath(path, env)
		PrintEnvironment(env)
	}
}