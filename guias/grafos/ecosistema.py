def amenazados(grafo):
    'Devolver una lista con los vértices que cumplen la condición'

    g_ent = grados_entrada(grafo)
    g_sal = grados_salida(grafo)
    especies = set()

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            if g_ent[w] == 1:
                especies.add(v)
            if g_sal[v] == 1:
                especies.add(w)            
    return list(especies)


def grados_entrada(g):
    # devolver un diccionario string -> int
    dic = {v: 0 for v in g.obtener_vertices()}
    for v in g.obtener_vertices():
        for adj in g.adyacentes(v):
            dic[adj] += 1
    return dic

def grados_salida(g):
    # devolver un diccionario string -> int
    dic = {}
    for v in g:
        dic[v] = len(g.adyacentes(v))
    return dic 
    