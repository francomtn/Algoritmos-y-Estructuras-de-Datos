from collections import deque
"""
5. Implementar un algoritmo que reciba un grafo dirigido, acíclico y pesado, un vértice v y otro w, y devuelva la longitud
del camino máximo. Indicar y justificar la complejidad del algoritmo implementado.
"""

def camino_maximo(g, v, w):

    dist = {v: float("-inf") for v in g}
    padres = {v: None}
    dist[v] = 0
    cola = deque()
    cola.append(v)
    while cola:
        v = cola.popleft()
        for x in g.adyacentes(v):   
            peso = g.peso_arista(v,x)
            if dist[v] + peso > dist[x]:
                dist[x] = dist[v]+ peso
                cola.append(x)
    return dist[w]
            

'''
def grados_entrada(grafo):
    g_ent = {}
    for v in grafo:
        g_ent[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            g_ent[w] += 1
    return g_ent

def topologico_grados(grafo):
    g_ent = grados_entrada(grafo)
    q = Cola()
    resultado = []
    for v in grafo:  # O(V)
        if g_ent[v] == 0:
            q.encolar(v)

    while not q.esta_vacia():
        v = q.desencolar()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                q.encolar(w)
    return resultado

def camino_maximo(g, v, w):
    order = topologico_grados(g)
    dist = {vertex: float('-inf') for vertex in g.obtener_vertices()}
    dist[v] = 0

    for u in order:
        if dist[u] != float('-inf'):
            for x in g.adyacentes(u):
                peso = g.peso_arista(u, x)
                if dist[u] + peso > dist[x]:
                    dist[x] = dist[u] + peso

    return dist[w] if dist[w] != float('-inf') else None
'''