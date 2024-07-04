import heapq

def centralidad(grafo):
    cent = {}
    for v in grafo: cent[v] = 0
    for v in grafo:
        for w in grafo:
            if v == w:
                continue
            padres, distancias = camino_minimo(grafo, v, w)
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

def camino_minimo(grafo, origen, destino):
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