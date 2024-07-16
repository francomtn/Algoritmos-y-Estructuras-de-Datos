from collections import deque

def seis_grados(grafo):

    for v in grafo:
        distancias = bfs(grafo, v)
        for d in distancias.values():
            if d > 6:
                return False
    return True

def bfs(g, vert):
    visitados = set()
    dist = {v: float("inf") for v in g}
    dist[vert] = 0
    cola = deque([vert])
    while cola:
        v = cola.popleft()
        for w in g.adyacentes(v):
            if w not in visitados:
                dist[w] = dist[v] + 1
                visitados.add(w)
                cola.append(w)
    return dist