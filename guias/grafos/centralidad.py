import heapq

def centralidad(grafo):
    cent = {}
    for v in grafo: cent[v] = 0
    for v in grafo:
        for w in grafo:
            if v == w:
                continue
            padres, distancias = camino_minimo_dijkstra(grafo, v, w)
            if padres[w] is None:
                continue
            actual = padres[w]
            while actual != v:
                cent[actual] += 1
                actual = padres[actual]
    return cent

def mas_centrales(grafo):
    ciudades = centralidad(grafo)
    return sorted(ciudades)

def camino_minimo_dijkstra(grafo, origen, destino):
    dist = {}
    padre = {}
    for v in grafo:
        dist[v] = float("inf")
    dist[origen] = 0
    padre[origen] = None
    q = [] #heap
    heapq.heappush(q, (0, origen))
    while q:
        _, v = heapq.heappop()
        if v == destino:
            return padre, dist
        for w in grafo.adyacentes(v):
            distancia_por_aca = dist[v] + grafo.peso(v, w)
            if distancia_por_aca < dist[w]:
                dist[w] = distancia_por_aca
                padre[w] = v
                heapq.heappush(q,(dist[w], w))
    return padre, dist

# para grafos no pesados
from cola import Cola

def camino_minimo_bfs(grafo, origen, destino):
   
   distancia, padre, visitado = {v: float("inf") for v in grafo}, {}, {}
   distancia[origen] = 0
   padre[origen] = None
   visitado[origen] = True
   q = Cola()
   q.encolar((0, origen))
   while not q.esta_vacia():
       dist, v = q.desencolar()
       if v == destino:
            return padre, dist
       for w in grafo.adyacentes(v):
           if (v not in visitado):
               distancia[w] += distancia[v] + 1
               padre[w] = v
               visitado[w] = True
               q.encolar((distancia[w], w))
   return padre, distancia
