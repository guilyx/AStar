module astar

DIAGONALS = true
HEURISTIC = 1

mutable struct Node
    pos::Int64
    goal_position::Int64
    fCost::Float64
    gCost::Float64
    hCost::Float64
    isClosed::Bool
    isOpen::Bool
    parent
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

function get_actions(env)
    actions = [1 1; 
               env.length 1; 
              -1 1; 
              -env.length 1]
    
    if DIAGONALS
        d = [-env.length-1 sqrt(2); 
             -env.length+1 sqrt(2); 
              env.length-1 sqrt(2); 
              env.length+1 sqrt(2)]
        actions = [actions;d]
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
    elseif color == "green"
        print("\033[1;32m")
        print(string)
        print("\033[0m")
    elseif color == "cyan"
        print("\033[1;36m")
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
    return env
end

function get_random_pos(env)
    size = env.length * env.height
    n = rand(1:size)
    while env.grid[n] == 1 
        n = rand(1:size)
    end
    return n
end

function get_start_goal(env, p_max_length)
    if p_max_length > 0.95
        p_max_length = 0.95
    end

    x = 0

    start = get_random_pos(env)
    goal = get_random_pos(env)
    size = env.length * env.height
    max_h = calculate_heuristic(env.length+2, size - env.length - 1, env)

    while calculate_heuristic(start, goal, env) <= max_h*p_max_length
        start = get_random_pos(env)
        goal = get_random_pos(env)
        x += 1
        if x > 1000
            print_color("red", "Could not find start and goal respecting constraints...\n")
            break
        end
    end

    return start, goal
end

function print_grid(env) 
    for (index, val) in enumerate(env.grid)
        if index % env.length != 0
            if val == 1
                print_color("red", "█")
            elseif val == 0
                print(" ")
            elseif val == 2
                print_color("blue", "¤")
            elseif val == 3
                print_color("green", "x")
            elseif val == 4
                print_color("cyan", "o")
            end
        else
            if val == 1
                print_color("red", "█")
            elseif val == 0
                print(" ")
            elseif val == 2
                print_color("blue", "¤")
            elseif val == 3
                print_color("green", "x")
            elseif val == 4
                print_color("cyan", "o")
            end
            print("\n")
        end
    end
end

function add_elem_to_env(env, pos, val)
    env.grid[pos] = val
    return env
end

function add_path(env, path)
    for (index, val) in enumerate(path.pos)
        if val == path.pos[1] || val == last(path.pos)
            continue
        end
        env = add_elem_to_env(env, val, 2)
    end

    return env
end

function calculate_heuristic(current_pos, goal_pos, env)
    row_current = convert(Int64, floor(current_pos / env.length)) + 1
    col_current = current_pos % env.length
    row_target = convert(Int64, floor(goal_pos / env.length)) + 1
    col_target = goal_pos % env.length

    dy = abs(row_target - row_current)
    dx = abs(col_target - col_current)

    if HEURISTIC == 0
        return(sqrt(dy^2 + dx^2))
    elseif HEURISTIC == 1
        return(dx + dy)
    else
        print_color("red", "Reached unreachable statement.")
        exit()
    end
end

function find_lowest_cost_node(nodes) 
    x = 0
    currentNode = nothing
    for (_, val) in nodes
        if val.isOpen
            x += 1
            currentNode = val
            break
        end
    end

    if !(x > 0)
        return 0, nothing
    end

    for (_, val) in nodes
        if val.isOpen && val != currentNode 
            x += 1
            if val.fCost < currentNode.fCost
                currentNode = val
            end
        end
    end

    return x, currentNode
end

function get_neighbours(current_node, goal_pos, env)
    c_pos = current_node.pos
    size = env.length * env.height

    children = []
    
    if c_pos < 0 || c_pos >= env.length*env.height || env.grid[c_pos] == 1
        print_color("red", "Tried to get neighbours from unreachable tile, exiting...")
        exit()
    end

    for row in eachrow(get_actions(env))
        if !(c_pos + row[1] > 0 && c_pos + row[1] <= size)
            continue
        end

        if env.grid[Int(c_pos + row[1])] == 1
            continue
        end

        child = node_constructor(Int(c_pos + row[1]), 
                                 goal_pos, 
                                 current_node.gCost + row[2], 
                                 false, false, current_node, env)
        
        children = push!(children, child)
    end
    
    return children
end

function node_constructor(position, goal_position, gcost, isclosed, isopen, parent, env)
    h = calculate_heuristic(position, goal_position, env)
    return(Node(position, goal_position, gcost + h, gcost, h, isclosed, isopen, parent))
end

function retrieve_path(last_node)
    current_node = last_node
    path = Path([], [])
    while current_node !== nothing
        path.pos = push!(path.pos, current_node.pos)
        path.costs = push!(path.costs, current_node.gCost)
        current_node = current_node.parent
    end
    return path
end

function search(start, goal, env)
    nodes = Dict{Int64, Node}()
    start_node = node_constructor(start, goal, 0, false, true, nothing, env)
    nodes[start] = start_node

    while true
        open_nodes_cnt, currentNode = find_lowest_cost_node(nodes)
        
        if open_nodes_cnt < 1
            print_color("red", "The graph has no open nodes left...\n")
            break
        end

        if currentNode.pos == goal
            print_color("yellow", "Path found!\n")
            return(retrieve_path(currentNode))
        end
        
        nodes[currentNode.pos].isOpen = false
        nodes[currentNode.pos].isClosed = true

        children = get_neighbours(currentNode, goal, env)

        for (_, child) in enumerate(children)
            if haskey(nodes, child.pos)
                if nodes[child.pos].isClosed
                    continue
                end

                if !nodes[child.pos].isOpen
                    child.isOpen = true
                    nodes[child.pos] = child
                else
                    if child.gCost < nodes[child.pos].gCost
                        child.isOpen = true
                        nodes[child.pos] = child
                    end
                end
            else
                child.isOpen = true
                nodes[child.pos] = child
            end
        end
    end
end

function main()
    println(PROGRAM_FILE, " starting...")
    length = 50
    height = 20
    wallP = .25
    env = generate_environment(length, height, wallP)

    start, goal = get_start_goal(env, 0.5)

    env = add_elem_to_env(env, start, 3)
    env = add_elem_to_env(env, goal, 3)

    print_grid(env)

    path = search(start, goal, env)

    if path !== nothing 
        env = add_path(env, path)
        print_grid(env)
    end
end

main()

end #astar