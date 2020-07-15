#include <stdlib.h>
#include <stdio.h>
#include <math.h>

#define clear()     printf("\033[H\033[J")
#define DIAGONALS   0
#define HEURISTIC   0

#define TRUE        1
#define FALSE       0

typedef struct {
    int posX;
    int posY;
} Position;

typedef struct {
    int** grid;
    int height;
    int length;
    float wallPercentage;
} Map;

typedef struct {
    Position pos;
    float f;
} Path;

typedef struct {
    Position pos;
    float f_cost;
    float h_cost;
    float g_cost;
    int isOpen;
    int isClosed;
    Position parent;
} Node;

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

void printMap(Map map) {
    for (int i3 = 0; i3 < map.height; i3++) {
        for (int j3 = 0; j3 < map.length; j3++) {
            switch(map.grid[i3][j3]) {
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

Map* addPath(Path* path, Map* map) {
    for (int i = 0 ; i < (sizeof(path) * sizeof(Path)) ; i++) {
        int x = path[i].pos.posX;
        int y = path[i].pos.posY;
        if (x == 0 && y == 0) continue;
        map->grid[x][y] = 2;
    }
    return map;
}

Position* CreatePosition(int positionX, int positionY) {
    Position *posRobot;
    posRobot = (Position*) malloc (sizeof(Position));
    posRobot->posX = positionX;
    posRobot->posY = positionY;
    return posRobot;
}

Position* placeRobot(int **map, int size) {
    Position *posRobot;
    int rand1 = rand() % size;
    int rand2 = rand() % size;
    while (map[rand1][rand2] == 1) {
        rand1 = rand() % size;
        rand2 = rand() % size;
    }
    map[rand1][rand2] = -1;
    posRobot = CreatePosition(rand1, rand2);
    return posRobot;
}

void DestroyPosition(Position *pos) {
    free(pos);
}

Position* getChildren(Map map, Position* posRobot, Node** nodes) {
    Position children[8];
    int i = 0;
    int x = posRobot->posX;
    int y = posRobot->posY;

    if (x < 0 || y < 0 || x >= map.length || y >= map.height || map.grid[x][y] == 1) return NULL;
    
    for (int dx = -1; dx <= 1; dx++) {
        for (int dy = -1; dx <= 1; dx++) {
            if (DIAGONALS == TRUE) {
                if(!(0 <= x+dx < map.height && 0 <= y+dy < map.length)) continue;
                if (map.grid[x+dx][y+dy] != 0) continue;
                children[i].posX = x+dx;
                children[i].posY = y+dy;
                if (i<8) i++;
            } else {
                if (dx == 0 || dy == 0) {
                    if(!(0 <= x+dx < map.height && 0 <= y+dy < map.length)) continue;
                    if (map.grid[x+dx][y+dy] != 0) continue;
                    children[i].posX = x+dx;
                    children[i].posY = y+dy;
                    if (i<8) i++;
                }
            }
        }
    }

    Position fixedChildren[i + 1];
    for (int j = 0 ; j < (i+1) ; j++) {
        fixedChildren[j] = children[j];
    }

    return fixedChildren;
}

float calculateHeuristic(Position nodePos, Position goalPos) {
    float h;
    if (HEURISTIC == 1) {
        double posPow1 = pow((goalPos.posX - nodePos.posX), 2);
        double posPow2 = pow((goalPos.posY - nodePos.posY), 2);
        h = sqrt( posPow1 + posPow2 );
    } else if (HEURISTIC == 0) {
        double dy = goalPos.posY - nodePos.posY;
        double dx = goalPos.posX - nodePos.posX;
        h = abs(dy) + abs(dx);
    } else {
        return 0.0;
    }
    
    return h;
}

Node* createNode(Position *pos) {
    Node *node;
    node = (Node*) malloc (sizeof(Node));
    node->pos = *pos;
    node->isOpen = FALSE;
    node->isClosed = FALSE;
    node->h_cost = 0;
    node->f_cost = 0;
    node->g_cost = 0;
    return node;
}

Map* generateMap(int height, int length, int wallPercentage)
{
    Map *map;
    map = (Map*) malloc (sizeof(Map));
    int **grid = (int **)malloc(height * sizeof(int *));
    for (int i = 0; i < height; i++)
        grid[i] = (int *)malloc(length * sizeof(int));

    for (int i2 = 0; i2 < height; i2++) {
        for (int j2 = 0; j2 < length; j2++) {
            float wallornot = (float) rand() / RAND_MAX;
            if (wallornot > wallPercentage) {
                if ((i2 == 0) || (j2 == 0) || (i2 == (height - 1)) || (j2 == (length - 1))) {
                    grid[i2][j2] = 1;
                } else {
                    grid[i2][j2] = 0;
                }
            } else
            {
                grid[i2][j2] = 1;
            }
        }
    }

    map->grid = grid;
    map->height = height;
    map->length = length;
    map->wallPercentage = wallPercentage;

    return map;
}

int compare (const Path *a, const Path *b) {
    return (a->f < b->f);
}

Path* retrievePath(Node lastNode, Node** nodes, Position start, Path* path, Path* cleanPath, Map map) {
    int i = 0;
    
    while(lastNode.pos.posX != start.posX && lastNode.pos.posY != start.posY) {
        if (lastNode.pos.posX == 0 && lastNode.pos.posY == 0) {
            int y = 100;
        }
        path[i].pos = lastNode.pos;
        lastNode.pos = lastNode.parent;
        lastNode = nodes[lastNode.pos.posX][lastNode.pos.posY];
        i++;
    }
    
    path[i].pos = lastNode.pos;

    cleanPath = (Path *)malloc(i+1 * sizeof(Path *));

    for (int x = 0 ; x < i+1 ; x++) {
        cleanPath[x] = path[x];
    }

    return cleanPath;
}

Node** initNodes(Map map, Position goal) {
    Node **nodes = (Node **)malloc(map.height * sizeof(Node *));
    for (int i = 0; i < map.height; i++)
        nodes[i] = (Node *)malloc(map.length * sizeof(Node));

    for (int x = 0; x < map.height ; x++) {
        for (int y = 0; y < map.length ; y++) {
            nodes[x][y].isClosed = FALSE;
            nodes[x][y].isOpen = FALSE;
            nodes[x][y].pos.posX = x;
            nodes[x][y].pos.posY = y;
            nodes[x][y].h_cost = calculateHeuristic(nodes[x][y].pos, goal);
            nodes[x][y].g_cost = 0;
            nodes[x][y].f_cost = 0;
        }
    }

    return nodes;
}

Path* search(Position start, Position goal, Map map, Path *cleanPath) {
    Node** nodes = initNodes(map, goal);
    Path path[map.height*map.length];
    float lowestf = 999999999;
    int currentX = start.posX;
    int currentY = start.posY;
    nodes[start.posX][start.posY].isOpen = TRUE;
    Node lastNode;

    while (nodes[goal.posX][goal.posY].isClosed == FALSE) {

        for (int x = 0; x < map.height ; x++) {
            for (int y = 0; y < map.length ; y++) {
                nodes[x][y].f_cost = nodes[x][y].g_cost + nodes[x][y].h_cost;
                if (nodes[x][y].isOpen == TRUE) {
                    if(nodes[x][y].f_cost <= lowestf) {
                        currentX = x;
                        currentY = y;
                        lowestf = nodes[x][y].f_cost;
                    }
                }
            }
        }

        lastNode = nodes[currentX][currentY];
        nodes[currentX][currentY].isOpen = FALSE;
        nodes[currentX][currentY].isClosed = TRUE;

        Position* currentPosition = CreatePosition(currentX, currentY);

        for (int dx = -1 ; dx <= 1 ; dx++) {
            for (int dy = -1 ; dy <= 1 ; dy++) {
                if(!(0 <= currentX+dx < map.height && 0 <= currentY+dy < map.length)) continue;
                if (map.grid[currentX+dx][currentY+dy] != 0) continue;

                if(nodes[currentX+dx][currentY+dy].isClosed == TRUE) continue;

                if (nodes[currentX+dx][currentY+dy].isOpen == FALSE) {
                    nodes[currentX+dx][currentY+dy].isOpen = TRUE;
                    if (dx == 0 || dy == 0) nodes[currentX+dx][currentY+dy].g_cost = nodes[currentX][currentY].g_cost + 1;
                    else nodes[currentX+dx][currentY+dy].g_cost = nodes[currentX][currentY].g_cost + sqrt(2);
                    nodes[currentX+dx][currentY+dy].f_cost = nodes[currentX][currentY].g_cost + nodes[currentX][currentY].h_cost;
                    nodes[currentX+dx][currentY+dy].parent = nodes[currentX][currentY].pos;
                } else {
                    if (dx == 0 || dy == 0) {
                        if (nodes[currentX+dx][currentY+dy].g_cost > nodes[currentX][currentY].g_cost + 1) {
                            nodes[currentX+dx][currentY+dy].parent = nodes[currentX][currentY].pos;
                            nodes[currentX+dx][currentY+dy].g_cost = nodes[currentX][currentY].g_cost + 1;
                            nodes[currentX+dx][currentY+dy].f_cost = nodes[currentX][currentY].g_cost + nodes[currentX][currentY].h_cost;
                        } else {
                            nodes[currentX+dx][currentY+dy].isOpen = TRUE;
                        }
                    } else {
                        if (nodes[currentX+dx][currentY+dy].g_cost > nodes[currentX][currentY].g_cost + sqrt(2)) {
                            nodes[currentX+dx][currentY+dy].parent = nodes[currentX][currentY].pos;
                            nodes[currentX+dx][currentY+dy].g_cost = nodes[currentX][currentY].g_cost + sqrt(2);
                            nodes[currentX+dx][currentY+dy].f_cost = nodes[currentX][currentY].g_cost + nodes[currentX][currentY].h_cost;
                        } else {
                            nodes[currentX+dx][currentY+dy].isOpen = TRUE;
                        }
                  }
              }
            }
        }
    }

    cleanPath = retrievePath(lastNode, nodes, start, path, cleanPath, map);

    return cleanPath;
}

int main() {
    int height = 10;
    int length = 20;
    Path *path;
    Map* map = generateMap(height, length, 0);
    printMap(*map);
    Position *start = CreatePosition(1, 1);
    Position *goal = CreatePosition(8, 18);
    path = search(*start, *goal, *map, path);
    map = addPath(path, map);
    printMap(*map);
    return 0;
}