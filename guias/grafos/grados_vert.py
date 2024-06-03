def grados(g):
    # devolver un diccionario string -> int
    grados = {}
    for v in g:
	    grados[v] = len(g.adyacentes(v))	
    return grados

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