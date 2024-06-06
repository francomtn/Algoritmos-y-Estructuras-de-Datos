from collections import deque

def obtener_orden(grafo):
    'Devolver una lista con un posible orden válido'

    res = []
    cola = deque()
    grados = grados_entrada(grafo)
    for v in grafo:
        if grados[v] == 0:
            cola.append(v)
    while cola:
        v = cola.pop()
        res.append(v)
        for w in grafo.adyacentes(v):
            grados[w] -= 1
            if grados[w] == 0:
                cola.append(w)
    return res

def grados_entrada(g):
    # devolver un diccionario string -> int
    dic = {v: 0 for v in g.obtener_vertices()}
    for v in g.obtener_vertices():
        for adj in g.adyacentes(v):
            dic[adj] += 1
    return dic