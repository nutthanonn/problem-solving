visited = []
def dfs(graph, node):
    if node not in visited:
        visited.append(node)
        for n in graph[node]:
            dfs(graph,n)

graph = {
    'A' : ['B', 'C', 'D'],
    'B' : ['A', 'C', 'D'],
    'C' : ['A', 'B', 'E'],
    'D' : ['A', 'B', 'E', 'F'],
    'E' : ['C', 'D', 'F'],
    'F' : ['D', 'E'],
}


dfs(graph,'A')
print(*visited)


# Cr : https://stackoverflow.com/users/157726/juan-leni