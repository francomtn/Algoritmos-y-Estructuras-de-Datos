import heapq
from tdas.grafo  import Grafo

def prim(grafo):

    visitados = set()
    cola = []
    v = grafo.vertice_aleatorio()
    for w in grafo.adyacentes(v):
        heapq.heappush(cola, (v,w, grafo.peso_arista(v, w)))

    arbol = Grafo(False, grafo.obetener_vertices())
    while cola:
        (v, w), peso = heapq.heappop()
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, peso)
        visitados.add(w)
        for x in grafo.adyacentes(w):
            if x not in visitados:
                heapq.heappush((w, x, grafo.peso_arista(w, x)))
    return arbol
