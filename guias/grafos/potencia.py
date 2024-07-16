from tdas.grafo import Grafo
from cola import Cola


"""
1. Implementar la función grafo_potencia(grafo, k), que reciba un grafo no dirigido y un entero k, y nos devuelva un
nuevo grafo, que tenga los mismos vértice que el grafo recibido por parámetro, y cuyas aristas (v, w) indican que v
está a lo sumo a k aristas de w en el grafo original. Indicar y justificar la complejidad del algoritmo implementado.
"""


def grafo_potencia(g, k):
    nuevo = Grafo(False, g.obtener_vertices())

    for v in g.obtener_vertices():
        distancias = obtener_distancias(g, v)
        for w, d in distancias.items():
            if d <= k and v != w:
                nuevo.agregar_arista(v, w)
    
    return nuevo

def obtener_distancias(g, v):
    dist = {v: 0}
    visitados = set(v)
    cola = Cola()
    cola.encolar(v)    
    while cola:
        actual = cola.desencolar()
        for w in g.adyacentes(actual):
            if w not in visitados:
                visitados.add(w)
                dist[w] = dist[actual] + 1
                cola.encolar(w)
    
    return dist

# O(V * (O + E))