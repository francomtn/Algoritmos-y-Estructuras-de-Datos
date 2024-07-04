from unionFind  import UnionFind
from tdas.grafo  import Grafo

def kruskal(grafo):

    conjuntos = UnionFind(grafo.obtener_vertices())
    aristas = sorted(obtener_aristas(grafo))
    arbol = Grafo(False, grafo.obtener_vertices())
    for a in aristas:
        v,w, peso = a
        if conjuntos.find(v) == conjuntos.find(w):
            continue
        arbol.agregar_arista(v, w, peso)
        conjuntos.union(v, w)
    return arbol


def obtener_aristas(grafo):

    visitados = set()
    aristas = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            if w not in visitados:
                aristas.append((v,w, grafo.peso_arista(v,w)))
        visitados.add(v)
    return aristas