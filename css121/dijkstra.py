graph = {
    'a': {'b': 4, 'd': 2},
    'b': {'a': 4, 'e': 3, 'c':3},
    'c': {'b': 3, 'z': 2},
    'd': {'a': 2, 'e': 3 },
    'e': {'b': 3, 'z': 1, 'd': 3},
    'z': {'c': 2, 'e': 1}
}


def dijkstra(graph, start, goal):
    shortest_distance = {}
    track_predecessor = {}
    unseenNodes = graph
    infinity = 999999
    track_path = []
    
    for node in unseenNodes:
        shortest_distance[node] = infinity

    shortest_distance[start] = 0

    while unseenNodes: # object
        min_distance_node = None
        for node in unseenNodes:
            if min_distance_node is None:
                min_distance_node = node
            elif shortest_distance[node] < shortest_distance[min_distance_node]:
                min_distance_node = node  # a

        path_options = graph[min_distance_node].items() # {'b': 4, 'd': 2}

        for child_node, weight in path_options:
            if weight + shortest_distance[min_distance_node] < shortest_distance[child_node]:
                shortest_distance[child_node] = weight + shortest_distance[min_distance_node]
                track_predecessor[child_node] = min_distance_node
                
        unseenNodes.pop(min_distance_node)

    currentNode = goal

    while currentNode != start:
        try:
            track_path.insert(0, currentNode)
            currentNode = track_predecessor[currentNode]

        except KeyError:
            print("Error")
            break
        
    track_path.insert(0, start)


    if shortest_distance[goal] != infinity:
        print("Width : " + str(shortest_distance[goal]))
        print("Path : " + str(track_path))

dijkstra(graph, 'a', 'z')




# https://github.com/Salt-N-Pepa/dijkstra-implementation