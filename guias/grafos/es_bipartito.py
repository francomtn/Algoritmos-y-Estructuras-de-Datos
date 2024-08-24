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
    valores = {}
    
    for vertice in grafo.obtener_vertices():
        if vertice not in valores:
            valores[vertice] = 0 
            if not _es_bipartito(grafo, vertice, valores):
                return False
    return True

def _es_bipartito(grafo, v, valores):
    for w in grafo.adyacentes(v):
        if w in valores:
            if valores[w] == valores[v]:
                return False  
        else:
            valores[w] = 1 - valores[v]  
            if not _es_bipartito(grafo, w, valores):
                return False
    return True