# Bellman-Ford
def camino_minimo_bf(grafo, origen):
    distancia = {}
    padre = {}
    for v in grafo:                     # O(V)
        distancia[v] = float("inf")
    distancia[origen] = 0
    padre[origen] = None
    aristas = obtener_aristas(grafo) # O(V + E)
    for i in range(len(grafo)): # for i:= 0; i < len(grafo); i++
        cambio = False
        for origen, destino, peso in aristas:
            if distancia[origen] + peso < distancia[destino]:
                cambio = True
                padre[destino] = origen
                distancia[destino] = distancia[origen] + peso
        if not cambio:
            return padre, distancia

    for v, w, peso in aristas:
        if distancia[v] + peso < distancia[w]:
            return None  # Hay un ciclo negativo (lanzar excep) panic(AAAAAAAAA)
    return padre, distancia

def obtener_aristas(g):

    res = []
    for v in g:
        for w in g.adyacentes(v):
            res.append((v, w, g.peso_arista(v, w)))
    return res