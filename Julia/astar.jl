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
    grid::Array{Array{Int64}}
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


function main()
    println(PROGRAM_FILE, " starting...")
    m = get_actions()
    println(m)
end

main()

end #astar