from pila import Pila
from tdas.grafo import Grafo

def puntos_de_articulacion(grafo):
    origen = grafo.vertice_aleatorio()
    ptos_artic = set()
    dfs_ptos_artic(grafo, origen, {origen}, {origen: None}, {origen: 0}, {}, ptos_artic, True)
    return ptos_artic

def dfs_ptos_artic(g, v, visitados, padres, orden, mas_bajo, ptos, es_raiz):
    hijos = 0
    mas_bajo[v] = orden[v]
    for w in g.adyacentes(v):
        if w not in visitados:
            visitados.add(v)
            hijos += 1
            orden[w] = orden[v] + 1
            padres[w] = v
            dfs_ptos_artic(g, v, visitados, padres, orden, mas_bajo, ptos, False)
            
