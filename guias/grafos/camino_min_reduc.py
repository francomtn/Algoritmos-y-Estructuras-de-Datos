from collections import deque

# El grafo recibido tiene si o si pesos 1 o 2
def camino_minimo(grafo, origen):
    dist, padres = {v: float("inf") for v in grafo.obtener_vertices()}, {}

    dist[origen] = 0
    padres[origen] = None
    cola = deque()
    cola.append(origen)
    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            peso = grafo.peso_arista(v, w)
            if peso == 1:
                if dist[v] + 1 < dist[w]:
                    dist[w] = dist[v] + 1
                    padres[w] = v
                    cola.append(w)
            elif peso == 2:
                if dist[v] + 2 < dist[w]:
                    dist[w] = dist[v] + 2
                    padres[w] = v
                    cola.appendleft(w)
    
    return padres # devolver diccionario de padres