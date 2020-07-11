#include <stdlib.h>
#include <stdio.h>
#include <math.h>

#define clear()     printf("\033[H\033[J")

typedef struct {
    int posX;
    int posY;
} Position;

typedef struct {
    Position pos;
    float f_cost;
    float h_cost;
    float g_cost;
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

void printMap(int **map, int size) {
    clear();
    for (int i3 = 0; i3 < size; i3++) {
        for (int j3 = 0; j3 < size; j3++) {
            switch(map[i3][j3]) {
                case(0):
                    printColor('b', ".");
                    break;
                case(1):
                    printColor('r', "█");
                    break;
                case(-1):
                    printColor('y', "¤");
                    break;
            }
        }
        printf("%c", '\n');
    }
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

Position* CreatePosition(int positionX, int positionY) {
    Position *posRobot;
    posRobot = (Position*) malloc (sizeof(Position));
    posRobot->posX = positionX;
    posRobot->posY = positionY;
    return posRobot;
}

void DestroyPosition(Position *pos) {
    free(pos);
}

Position findEscape(int** map, int size) {
    Position posEscape;
    for (int i = 0 ; i < size ; ++i) {
        for (int j = 0 ; i < size ; ++j) {
            if ( (i == 0) || (j == 0) || (i == (size - 1)) || (j == (size - 1))) {
                if (map[i][j] == 2) {
                    posEscape.posX = i;
                    posEscape.posY = j;
                    return posEscape;
                }
            }
        }
    }
}

Position* getChildren(int** map, int mapsize, Position* posRobot) {
    Position children[4];
    int i = 0;
    int x = posRobot->posX;
    int y = posRobot->posY;
    if (x < 0 || y < 0 || x >= mapsize || y >= mapsize || map[x][y] == 1) {
        return 0;
    } else {
        if (map[x+1][y] != 1) {
            children[i].posX = x + 1;
            children[i].posY = y;
            if (i < 4) i++;
        } 
        if (map[x][y+1] != 1) {
            children[i].posX = x + 1;
            children[i].posY = y;
            if (i < 4) i++;
        }
        if (map[x][y-1] != 1) {
            children[i].posX = x + 1;
            children[i].posY = y;
            if (i < 4) i++;
        }
        if (map[x-1][y] != 1) {
            children[i].posX = x + 1;
            children[i].posY = y;
            if (i < 4) i++;
        }
    }

    Position fixedChildren[i + 1];
    for (int j = 0 ; j < (i+1) ; j++) {
        fixedChildren[j] = children[j];
    }

    return fixedChildren;
    
}

Node* Heuristic(Node *node, Node *goal) {
    float resDistance;
    double posPow1 = pow((goal->pos.posX - node->pos.posX), 2);
    double posPow2 = pow((goal->pos.posY - node->pos.posY), 2);
    resDistance = sqrt( posPow1 + posPow2 );
    
    
    node->h_cost = resDistance;
    node->f_cost = node->h_cost + node->g_cost;
    return node;
}

int **generateMap(int size)
{
    int **map = (int **)malloc(size * sizeof(int *));
    for (int i = 0; i < size; i++)
        map[i] = (int *)malloc(size * sizeof(int));

    for (int i2 = 0; i2 < size; i2++) {
        for (int j2 = 0; j2 < size; j2++) {
            int wallornot = rand() % 101;
            if (wallornot > 20) {
                if ((i2 == 0) || (j2 == 0) || (i2 == (size - 1)) || (j2 == (size - 1))) {
                    map[i2][j2] = 1;
                } else {
                    map[i2][j2] = 0;
                }
            } else
            {
                map[i2][j2] = 1;
            }
        }
    }

    return map;
}

int* search() {
    int* path
    return path
}

int main() {
    return 0
}