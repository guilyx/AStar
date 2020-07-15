cmake_minimum_required(VERSION 3.10)
project(c_astar)

set(SOURCE_FILES astar.c)

add_executable(c_astar ${SOURCE_FILES})
target_link_libraries(c_astar m)