#include <vector>
#include <map>
#include <random>
#include <stdlib.h>
#include <stdio.h>
#include <math.h>

#define clear()     printf("\033[H\033[J")


#define diagonals true
#define heuristic 0

typedef struct {
    int posX;
    int posY;
} Position;

typedef struct {
    int movX;
    int movY;
    float cost;
} Actions;

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
        World(int height, int length, float wallPercentage);
        void printGrid();
        void addPath(std::vector<Position> path);
    private:
        int m_height;
        int m_length;
        float m_wallPercentage;
        int** m_grid;
        bool m_hasPath;
        void generateGrid();
};

World::World(int height, int length, float wallPercentage) {
    m_height = height;
    m_length = length;
    m_wallPercentage = wallPercentage;

    generateGrid();
}

void World::generateGrid() {
    m_grid = (int **)malloc(m_height * sizeof(int *));
    for (int i = 0; i < m_height; i++)
        m_grid[i] = (int *)malloc(m_length * sizeof(int));

    for (int i2 = 0; i2 < m_height; i2++) {
        for (int j2 = 0; j2 < m_length; j2++) {
            float wallornot = rand();
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
    for (int i3 = 0; i3 < m_height; i3++) {
        for (int j3 = 0; j3 < m_length; j3++) {
            for (auto &i : path) {
                m_grid[i.posY][i.posX] = 3;
            }
        }
    }
}

class Node {
    public:
        Node(Position pos, Position targetPos, Node parent, float gCost);
        void printNode();
    private:
        float m_hCost;
        float m_gCost;
        float m_fCost;
        Position m_pos;
        Position m_targetPos;
        void calculateHeuristic();
};

class AStar {
    public:
        AStar(Position start, Position goal, World w);
        std::vector<Position> search();
    private:
        std::vector<Node> getNeighbours();
        std::vector<Position> retrievePath(Node lastNode);
        Position m_start;
        Position m_goal;
        World m_world;
};

/*
for (auto it = openSet.begin(); it != openSet.end(); it++) {
            auto node = *it;
            if (node->getScore() <= current->getScore()) {
                current = node;
                current_it = it;
            }
        }
*/

int main(int argc, char const *argv[])
{
    World w(10, 30, 0.1);
    w.printGrid();
    return 0;
}
