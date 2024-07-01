from tdas.grafo import Grafo
import heapq

def minimas_inversiones(grafo, s, t):

    nuevo = Grafo(es_dirigido=True, vertices_init = [v for v in grafo.obtener_vertices()])
    for v in grafo:
        for w in grafo.adyacentes(v):
            nuevo.agregar_arista(v, w, 0)
            nuevo.agregar_arista(w, v, 1)
    
    return dijkstra(nuevo, s, t)

def dijkstra(g, origen, destino):

    dist = { v: float("inf") for v in g}
    dist[origen] = 0
    heap = [(origen, 0)]
    while heap:
        v, dist_actual = heapq.heappop(heap)
        for w in g.adyacentes(v):
            if dist[v] + g.peso_arista(v, w) < dist[w]:
                dist[w] = dist[v] + g.peso_arista(v, w)
                heapq.heappush(heap, (w, dist[w]))
    return dist[destino]


