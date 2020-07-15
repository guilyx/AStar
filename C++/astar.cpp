#include <vector>
#include <map>
#include <random>
#include <stdlib.h>
#include <algorithm>
#include <stdio.h>
#include <cmath>
#include <iostream>
#include <set>

#define clear()     printf("\033[H\033[J")


#define DIAGONALS            true
#define HEURISTIC_FUNCTION   1

typedef struct {
    int posX;
    int posY;
} Position;

typedef struct {
    int movX;
    int movY;
    float cost;
} Actions;

typedef struct {
    float manhattan(Position a, Position b) {
        int dy = abs(b.posY - a.posY);
        int dx = abs(b.posX - a.posX);
        float h = dx + dy;
        return dx + dy;
    }
    float euclidian(Position a, Position b) {
        int dy = b.posY - a.posY;
        int dx = b.posX - a.posX;
        float h = sqrt(pow(dx, 2) + pow(dy, 2));
        return h;
    }
} Heuristics;

void printColor(char color, char* printNumb) {
    switch(color) {
        case('r'):
            printf("\033[1;31m");
            printf("%s", printNumb);
            printf("\033[0m");
            break;
        case('b'):
            printf("\033[1;34m");
            printf("%s", printNumb);
            printf("\033[0m");
            break;
        case('y'):
            printf("\033[0;33m");
            printf("%s", printNumb);
            printf("\033[0m");
            break;

    }
}

class World {
    public:
        World(int height = 10, int length = 10, float wallPercentage = .1);
        void printGrid();
        void addPath(std::vector<Position> path);
        std::vector<Actions> GetActions();
        int** GetGrid();
        int GetHeight();
        int GetLength();
        Position getRandomFreePosition();
    private:
        int m_height;
        int m_length;
        float m_wallPercentage;
        std::vector<Actions> m_actions;
        int** m_grid;
        bool m_hasPath;
        void generateGrid();
};

World::World(int height, int length, float wallPercentage) {
    m_height = height;
    m_length = length;
    m_wallPercentage = wallPercentage;
    if (DIAGONALS) {
        m_actions.push_back(Actions{1, 0, 1});
        m_actions.push_back(Actions{0, 1, 1});
        m_actions.push_back(Actions{-1, 0, 1});
        m_actions.push_back(Actions{0, -1, 1});
        m_actions.push_back(Actions{1, 1, sqrt(2)});
        m_actions.push_back(Actions{-1, 1, sqrt(2)});
        m_actions.push_back(Actions{-1, -1, sqrt(2)});
        m_actions.push_back(Actions{1, -1, sqrt(2)});
    } else {
        m_actions.push_back(Actions{1, 0, 1});
        m_actions.push_back(Actions{0, 1, 1});
        m_actions.push_back(Actions{-1, 0, 1});
        m_actions.push_back(Actions{0, -1, 1});
    }
    generateGrid();
}

std::vector<Actions> World::GetActions() {
    std::vector<Actions> actions = this->m_actions;
    return actions;
}

int** World::GetGrid() {
    int** grid = this->m_grid;
    return grid;
}

int World::GetLength() {
    int length = this->m_length;
    return length;
}

int World::GetHeight() {
    int height = this->m_height;
    return height;
}

Position World::getRandomFreePosition() {
    int rand_h = rand() % (m_height - 1);
    int rand_l = rand() % (m_length - 1);

    while (m_grid[rand_h][rand_l] != 0) {
        rand_h = rand() % (m_height - 1);
        rand_l = rand() % (m_length - 1);
    }
    
    return(Position{rand_h, rand_l});

}

void World::generateGrid() {
    m_grid = (int **)malloc(m_height * sizeof(int *));
    for (int i = 0; i < m_height; i++)
        m_grid[i] = (int *)malloc(m_length * sizeof(int));

    for (int i2 = 0; i2 < m_height; i2++) {
        for (int j2 = 0; j2 < m_length; j2++) {
            float wallornot = (float) rand() / RAND_MAX;
            if (wallornot > m_wallPercentage) {
                if ((i2 == 0) || (j2 == 0) || (i2 == (m_height - 1)) || (j2 == (m_length - 1))) {
                    m_grid[i2][j2] = 1;
                } else {
                    m_grid[i2][j2] = 0;
                }
            } else
            {
                m_grid[i2][j2] = 1;
            }
        }
    }
}

void World::printGrid() {
    clear();
    for (int i3 = 0; i3 < m_height; i3++) {
        for (int j3 = 0; j3 < m_length; j3++) {
            switch(m_grid[i3][j3]) {
                case(0):
                    printColor('b', ".");
                    break;
                case(1):
                    printColor('r', "█");
                    break;
                case(2):
                    printColor('y', "¤");
                    break;
            }
        }
        printf("%c", '\n');
    }
}

void World::addPath(std::vector<Position> path) {
    for (auto &i : path) {
        m_grid[i.posX][i.posY] = 2;
    }
    this->m_hasPath = true;
}

class Node {
    public:
        Node(Position pos = {-1, -1}, Position targetPos = {-1, -1}, Node* parent = nullptr, float gCost = sizeof(float));
        void printNode();
        Node* GetParent();
        Position GetPosition();
        Position GetTargetPosition();
        float GetGCost();
        float GetHCost();
        float GetFCost();
        void SetPosition(Position pos);
        void SetTargetPosition(Position targetPos);
        void SetParent(Node* parent);
        void SetGCost(float gCost);
    private:
        float m_hCost;
        float m_gCost;
        float m_fCost;
        Node* m_parent;
        Position m_pos;
        Position m_targetPos;
        void calculateHeuristic();
};

Node::Node(Position pos, Position targetPos, Node* parent, float gCost) {
    this->m_pos = pos;
    this->m_targetPos = targetPos;
    this->m_parent = parent;
    this->m_gCost = gCost;
    calculateHeuristic();
    this->m_fCost = this->m_gCost + this->m_hCost;
}

Node* Node::GetParent() {
    Node* parent = this->m_parent;
    return parent;
}

Position Node::GetPosition() {
    Position pos = this->m_pos;
    return pos;
}

Position Node::GetTargetPosition() {
    Position pos = this->m_targetPos;
    return pos;
}

float Node::GetGCost() {
    float cost = this->m_gCost;
    return cost;
}

float Node::GetHCost() {
    float cost = this->m_hCost;
    return cost;
}

float Node::GetFCost() {
    float cost = this->m_fCost;
    return cost;
}

void Node::SetPosition(Position pos) {
    this->m_pos = pos;
}
void Node::SetTargetPosition(Position targetPos) {
    this->m_targetPos = targetPos;
}
void Node::SetParent(Node* parent) {
    this->m_parent = parent;
}
void Node::SetGCost(float gCost) {
    this->m_gCost = gCost;
}

void Node::printNode() {
    std::cout << "Node [ Position : ( " << this->m_pos.posX << " ; " << this->m_pos.posY << 
                 " ) || Costs : ( G : " << this->m_gCost << " ; H : " << this->m_hCost << 
                 " ; F : " << this->m_fCost << " ) ]" << std::endl;
}
void Node::calculateHeuristic() {
    Heuristics hFunctions;
    if (HEURISTIC_FUNCTION == 0) this->m_hCost = hFunctions.euclidian(this->m_pos, this->m_targetPos);
    else if (HEURISTIC_FUNCTION == 1) this->m_hCost = hFunctions.manhattan(this->m_pos, this->m_targetPos);
    else {
        std::cerr << "Heuristic function not found !!" << std::endl;
        exit(EXIT_FAILURE);
    }
}

bool operator<(const Node &node, const Node &otherNode) {
    Node x = node;
    Node y = otherNode;
    return x.GetFCost() < y.GetFCost();
}

/*
bool operator==(const Node &node, const Node &otherNode) {
    Node x = node;
    Node y = otherNode;

    return ((x.GetPosition().posY == x.GetPosition().posY) && (x.GetPosition().posX == y.GetPosition().posX) && (x.GetParent() == y.GetParent()) && (x.GetFCost() == y.GetFCost()));
}*/


class AStar {
    public:
        AStar(Position start, Position goal, World w);
        std::vector<Position> search();
    private:
        std::vector<Node> getNeighbours(Node node);
        std::vector<Position> retrievePath(Node* lastNode);
        Node* findNodeInVector(Node node, std::vector<Node> vector);
        Node m_nStart;
        Node m_nGoal;
        bool m_reached;
        std::vector<Node> m_openNodes;
        std::vector<Node> m_closedNodes;
        World m_world;
};

AStar::AStar(Position start, Position goal, World w) {
    this->m_nStart = Node(start, goal, nullptr, 0);
    this->m_nGoal = Node(goal, goal, nullptr, sizeof(float));
    this->m_world = w; 
    this->m_openNodes.push_back(this->m_nStart);
    this->m_reached = false;
}

std::vector<Node> AStar::getNeighbours(Node node) {
    std::vector<Node> children;
    for (auto &i : this->m_world.GetActions()) {
        int posy = node.GetPosition().posY + i.movY;
        int posx = node.GetPosition().posX + i.movX;
        
        Node* nodeptr = new Node(node);

        if (!((0 <= posx < this->m_world.GetLength()) && (0 <= posy < this->m_world.GetHeight()))) continue;
        if (this->m_world.GetGrid()[posx][posy] != 0) continue;

        Node child(Position{posx, posy}, this->m_nStart.GetTargetPosition(), nodeptr, node.GetGCost() + i.cost);

        children.push_back(child);
    }
    return children;
}

std::vector<Position> AStar::retrievePath(Node* node) {
    Node* current_node = node;
    std::vector<Position> path;
    while (current_node != nullptr) {
        path.push_back(current_node->GetPosition());
        current_node = current_node->GetParent();
    }
    std::reverse(path.begin(), path.end());
    return path;
}

Node* AStar::findNodeInVector(Node mNode, std::vector<Node> vector) {
    for (auto &node : vector) {
        if ((mNode.GetPosition().posX == node.GetPosition().posX) && 
            (mNode.GetPosition().posY == node.GetPosition().posY)) {
            return(&mNode);
        }
    }
    return nullptr;
}
std::vector<Position> AStar::search() {
    while (!this->m_openNodes.empty() || !m_reached) {
        Node best_node = *m_openNodes.begin();
        auto toDelete = &(*m_openNodes.begin());
        for (auto &i : m_openNodes) {
            if (i.GetFCost() <= best_node.GetFCost()) {
                best_node = i;
                toDelete = &i;
            }
        }
        if ((best_node.GetPosition().posX == this->m_nGoal.GetPosition().posX) && 
            (best_node.GetPosition().posY == this->m_nGoal.GetPosition().posY)) {

            std::vector<Position> path = retrievePath(&best_node);
            return path;
        }
        // We now have the best node, remove it from the open nodes
        this->m_openNodes.erase(this->m_openNodes.begin() + (toDelete - &this->m_openNodes[0]));
        this->m_closedNodes.push_back(best_node);

        std::vector<Node> childrenNodes = this->getNeighbours(best_node);

        for (auto &child : childrenNodes) {
            if (findNodeInVector(child, m_closedNodes) != nullptr) continue;

            Node* cNode = findNodeInVector(child, m_openNodes);

            if (cNode == nullptr) {
                this->m_openNodes.push_back(child);
            } else {
                Node betterNode = *cNode;

                if (betterNode.GetGCost() < child.GetGCost()) this->m_openNodes.push_back(betterNode);
                else this->m_openNodes.push_back(child);
            }
        }
    }

    if (!this->m_reached) {
        std::vector<Position> path;
        path.push_back(this->m_nStart.GetPosition());
        return path;
    }
}

int main(int argc, char const *argv[])
{
    World w(20, 30, 0.1);
    w.printGrid();
    Position start = w.getRandomFreePosition();
    Position goal = w.getRandomFreePosition();
    std::cout << "START ( " << start.posX << " ; " << start.posY << " )" << std::endl;
    std::cout << "GOAL ( " << goal.posX << " ; " << goal.posY << " )" << std::endl;
    AStar astar(start, goal, w);
    std::vector<Position> path = astar.search();
    std::cout << "Path Found." << std::endl;
    w.addPath(path);
    w.printGrid();
    return 0;
}
