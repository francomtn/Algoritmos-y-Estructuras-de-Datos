from tdas.grafo import Grafo

def cantidad_componentes_debiles(grafo):

    nuevo = Grafo(es_dirigido=False, vertices_init= grafo.obtener_vertices())

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if not nuevo.estan_unidos(v, w):
                nuevo.agregar_arista(v, w)    
    return contar_componetes_conexas(nuevo)

def contar_componetes_conexas(grafo):
    visitados = set()
    comps = 0
    resultado = []
    for v in grafo:
        if v not in visitados:
            nueva_componente = []
            resultado.append(nueva_componente)
            comps += 1
            dfs_comps(grafo, v, visitados, nueva_componente)
    return comps

def dfs_comps(grafo, v, visitados, componente):
    visitados.add(v)
    componente.append(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs_comps(grafo, w, visitados, componente)
