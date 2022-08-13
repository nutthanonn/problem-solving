graph = {
    'A': ['B', 'C', 'D', 'E', 'G'],
    'B': ['F'],
    'C': ['F', 'G', 'J'],
    'D': ['G', 'K', 'H'],
    'E': ['H'],
    'F': ['I', 'M', 'J'],
    'G': ['J', 'K'],
    'H': ['K', 'O', 'L'],
    'I': ['M'],
    'J': ['M', 'N'],
    'K': ['N', 'O'],
    'L': ['O'],
    'M': ['P', 'N'],
    'N': ['P', 'Q'],
    'O': ['N', 'Q'],
    'P': ['R'],
    'Q': ['R'],
}

count = 1


def dfs(graph, start, visited):
    global count
    visited.append(start)
    for i in graph[start]:
        if i != 'R':
            dfs(graph, i, visited.copy())
        else:
            visited.append("R")
            print(f"{count}. {'-'.join(visited)}")
            count += 1


dfs(graph, "A", [])
