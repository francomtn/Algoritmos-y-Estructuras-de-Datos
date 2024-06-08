"""
Escribir el pseudocódigo del algoritmo de Dijkstra para obtener los caminos mínimos desde un vértice 
hacia todos los demás vértices del grafo. ¿Cómo podemos modificar el algoritmo de Dijkstra para que en 
caso de tener dos caminos mínimos se elija al de menor cantidad de aristas? 
Indicar y justificar el orden del algoritmo.
"""
def obtener_camino(g, origen):

    dist = {}
    padres = {}


"""
Implementar un algoritmo que, dado un grafo dirigido, nos devuelva un ciclo 
dentro del mismo, si es que lo tiene. Indicar el orden del algoritmo.

"""

from collections import deque

def hay_ciclo(g):

    visitados = set()
    cola = deque()
    cola.append(g.vertice_aleatorio())
    while cola:
        v = cola.popleft()
        visitados.add(v)
        for w in g.adyacentes(v):
            if w in visitados:
                return
            cola.append(w)
