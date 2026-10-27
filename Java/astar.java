package astar;

class Environment {
    int length;
    int height;
    float p_walls;
    int[][] grid;
    float[][] actions;
    
    private static final boolean DIAGONALS = false;

    private Environment(int l_, int h_, float p_) {
        length = l_;
        height = h_;
        p_walls = p_;
        grid = new int[height][length];
        
        if (DIAGONALS == true) {
            length = 0;
        } else {
            actions = new float[][] {
                {1, 0, 1},
                {0, 1, 1},
                {-1, 0, 1},
                {0, -1, 1},
            };
        }
    }
}

class Position {
    int x;
    int y;

    public Position(int x_, int y_) {
        x = x_;
        y = y_;
    }
}

class Node {
    
}

public class astar {

}