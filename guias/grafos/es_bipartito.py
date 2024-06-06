from cola import Cola

def es_bipartito(grafo):
    colores = {}
    for v in grafo.obtener_vertices():
        if v not in colores:
            if not es_bipartito_aux(grafo, v):
                    return False
    return True

def es_bipartito_aux(grafo, vert_inicial):

    colores = {}
    cola = Cola()
    colores[vert_inicial]= 0
    cola.encolar(vert_inicial)
    while not cola.esta_vacia():
        v = cola.desencolar()
        for w in grafo.adyacentes(v):
            if w in colores and colores[v] == colores[w]:
                return False
            if w not in colores:
                colores[w] = 1 - colores[v]
                cola.encolar(w)
    return True

def es_bipartito(grafo):
	colores = { }
	for vertice in grafo.obtener_vertices():
		if vertice not in colores:
			if not _es_bipartito(grafo, vertice, colores):
				return False
	return True

def _es_bipartito(grafo, v, colores):

    for w in grafo.adyacentes(v):
        if w in colores:
            if colores[w] == colores[v]:
                return False
        else:
            colores[w] = 1 - colores[v]
            if not _es_bipartito(grafo, w, colores):
                return False
    return True
