#include <vector>
#include <map>
#include <random>

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

class World {
    public:
        World(int height, int length, float wallPercentage);
        void printGrid();
        void addPath();
    private:
        int m_height;
        int m_length;
        float m_wallPercentage;
        int** m_grid;
        bool m_hasPath;
        void generateGrid();
};

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
        std::vector<Node> retrievePath(Node lastNode);
        Position m_start;
        Position m_goal;
        World m_world;
};

'''
for (auto it = openSet.begin(); it != openSet.end(); it++) {
            auto node = *it;
            if (node->getScore() <= current->getScore()) {
                current = node;
                current_it = it;
            }
        }
'''