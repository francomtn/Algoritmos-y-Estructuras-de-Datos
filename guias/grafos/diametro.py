from collections import deque

def diametro(grafo):
    dist_max = 0
    for v in grafo.obtener_vertices():
        dist = bfs(grafo, v)
        dist_max = max(dist_max, max(dist.values()))
    return dist_max


def bfs(grafo, origen):

    dist = {v: float("inf") for v in grafo.obtener_vertices()}
    dist[origen] = 0
    cola = deque()
    cola.append(origen)
    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if dist[w] == float("inf"):
                dist[w] = dist[v] + 1
                cola.append(w)
    return dist