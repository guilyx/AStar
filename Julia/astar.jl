module astar

DIAGONALS = false
HEURISTIC = 1

mutable struct Node
    pos::Int64
    fCost::Float64
    gCost::Float64
    hCost::Float64
    parent::Node
end

mutable struct Path
    pos::Array{Int64}
    costs::Array{Float64}
end

mutable struct Environment
    length::Int64
    height::Int64
    wallPercentage::Float64
    grid::Array{Int8}
end

function get_actions()
    actions = [1 0 1;
               0 1 1;
              -1 0 1;
               0 -1 1]
    
    if DIAGONALS
        append!(actions, [1 1 sqrt(2); 
                         -1 -1 sqrt(2); 
                         -1 1 sqrt(2); 
                          1 -1 sqrt(2)])
    end
    return actions
end

function print_color(color, string)
    if color == "red"
        print("\033[1;31m")
        print(string)
        print("\033[0m")
    elseif color == "blue"
        print("\033[1;34m")
        print(string)
        print("\033[0m")
    elseif color == "yellow"
        print("\033[1;33m")
        print(string)
        print("\033[0m")
    else
        print(string)
    end
end

function generate_environment(length, height, wallPercentage)
    gd = zeros(Int8, height, length)
    
    for (index, _) in enumerate(gd)
        if index % length == 0 || index <= length || index > length*height - length || (index - 1) % length == 0
            gd[index] = 1
            continue
        end
        if index//height % length == 0
            gd[index] = 1
            continue
        end
        if rand() < wallPercentage
            gd[index] = 1
        end

    end

    env = Environment(length, height, wallPercentage, gd)

end

function print_grid(env) 
    for (index, val) in enumerate(env.grid)
        if index % env.length != 0
            if val == 1
                print_color("red", "█")
            elseif val == 0
                print(" ")
            else
                print_color("blue", "¤")
            end
        else
            if val == 1
                print_color("red", "█")
            elseif val == 0
                print(" ")
            else
                print_color("blue", "¤")
            end
            print("\n")
        end
    end
end

function add_path(env, path)
    for (index, val) in enumerate(path.pos)
        env.grid[val] = 3
    end

    return env
end

function main()
    println(PROGRAM_FILE, " starting...")
    length = 20
    height = 10
    wallP = .1
    env = generate_environment(length, height, wallP)
    print_grid(env)
end

main()

end #astar