def kruskal(graph):
    node = []; weight = 0; path = []
    cicleNode = []
    newGraph = {k: v for k, v in sorted(graph.items(), key=lambda item: item[1])}
    for (i, j), w_ij in newGraph.items():
        if ord(i) > ord(j):
            i, j = j, i

        if j not in node:
           weight += w_ij
           node.append(j)
           cicleNode.append(i)
           path.append(i + " ---> " + j)

        else:
            if ord(i) < ord(j):
                i, j = j, i
            count = cicleNode.count(j)
            for (fNode, sNode), w_fs in newGraph.items():
                if j not in node:
                    if sNode == j and w_ij < w_fs and count == 0:
                        node.append(j)
                        path.append(i + " ---> " + j)
                        weight += w_ij

    return path, weight

# i = I
# j = K

graph1 = {
    ('a', 'b'): 4,
    ('a', 'c'): 7,
    ('a', 'f'): 12,
    ('b', 'd'): 7,
    ('b', 'c'): 8,
    ('c', 'e'): 4,
    ('c', 'f'): 8,
    ('e', 'd'): 7,
    ('e', 'h'): 5,
    ('e', 'g'): 10,
    ('e', 'g'): 10,
    ('g', 'f'): 4,
    ('g', 'h'): 10,
    ('d', 'h'): 6,

}

graph2 = {
    ('a', 'b'): 1,
    ('a', 'c'): 2,
    ('a', 'd'): 3,
    ('b', 'c'): 5,
    ('b', 'e'): 3,
    ('c', 'e'): 5,
    ('c', 'f'): 2,
    ('c', 'g'): 4,
    ('c', 'd'): 5,
    ('d', 'g'): 4,
    ('e', 'f'): 3,
    ('e', 'h'): 2,
    ('f', 'h'): 1,
    ('f', 'g'): 4,
    ('h', 'i'): 5,
    ('h', 'j'): 5,
    ('h', 'g'): 2,
    ('h', 'k'): 3,
    ('i', 'k'): 4,
    ('j', 'k'): 3,
}


graph3 = {
    ('1', '2'):4,
    ('1', '4'):6,
    ('2', '4'):8,
    ('2', '3'):2,
    ('4', '5'):1,
    ('3', '5'):7,
    ('2', '5'):9,
}

path, weight = kruskal(graph3)
path.sort()
for i in path:
    print(i)
print("Weight : " + str(weight))