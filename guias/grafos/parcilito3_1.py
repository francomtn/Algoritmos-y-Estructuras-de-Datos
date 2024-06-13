

"""
Escribir una función que dado un grafo dirigido pasado por parámetro, encuentre y retorne todas las componentes
débilmente conexas del grafo, que consten de dos o más vértices. Indicar y justificar la complejidad de la función.
"""
def comp_debilmente_conex(g):

    nuevo = Grafo(False, g.obtener_vertices())
    for v in g.obtener_vertices():
        for w in g.adyacentes(v):
            if not nuevo.estan_unidos():
                nuevo.agregar_arista(v, w)
    componentes =  contar_componentes(g)
    resultado = [comp for comp in componentes if len(comp)> 1]
    return resultado

def contar_componentes(g):

    visitados = set()
    res = []
    for v in g:
        if v not in visitados:
            nueva_comp = []
            visitados.add(v)
            res.append(nueva_comp)
            dfs(g, v, visitados, nueva_comp)
    return res

def dfs(g, v, visitados, componentes):

    visitados.add(v)
    componentes.append(v)
    for w in g.adyacentes(v):
        if w not in visitados:
            dfs(g, w, visitados, componentes)

