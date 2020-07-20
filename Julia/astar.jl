module astar

DIAGONALS = false
HEURISTIC = 1

mutable struct Position
    x::Int64
    y::Int64
end

mutable struct Node
    pos::Position
    fCost::Float64
    gCost::Float64
    hCost::Float64
    parent::Node
end

mutable struct Path
    pos::Array{Position}
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

function generate_environment(length, height, wallPercentage)
    gd = zeros(Int8, height, length)
    
    for (index, _) in enumerate(gd)
        if rand() < wallPercentage
            gd[index] = 1
        end

    end

    env = Environment(length, height, wallPercentage, gd)

end

function print_grid(env) 
    for (index, val) in enumerate(env.grid)
        if index % env.length != 0
            print(val, " ")
        else
            print("\n")
        end
    end
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