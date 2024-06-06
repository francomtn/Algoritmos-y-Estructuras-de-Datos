def encontrar_ciclo(g):
    '''
    Devuelve una lista de vertices que conforman el ciclo. En el segundo ejemplo, 
    debería devolver [A, B, C] (o [B, C, A], etc...). 
    Si no hay ciclo, debe devolver None. 
    '''
    visitados = set()
    padres = {}
    for v in g:
        if v not in visitados:
            ciclo = dfs(g, v, visitados, padres)
            if ciclo is not None:
                return ciclo
    return  None

def dfs(grafo, v, visitados, padres):
    visitados.add(v)
    for w in grafo.adyacentes(w):
        if w in visitados:
            if w != padres[v]:
                return True
            else:
                padres[w]= v
                ciclo = dfs(grafo, w, visitados, padres)
                if ciclo:
                    return True
    return False


"""
def encontrar_ciclo(g):
    visitados = set()
    padres = {}

    for v in g:
        if v not in visitados:
            ciclo = _dfs_ciclo(g, v, visitados, padres)

            if ciclo is not None:
                return ciclo

    return None


def _dfs_ciclo(g, v, visitados, padres):
    visitados.add(v)

    for w in g.adyacentes(v):
        if w in visitados and w != padres[v]:
            return recons_ciclo(padres, w, v)

        if w not in visitados: 
            padres[w] = v
            ciclo = _dfs_ciclo(g, w, visitados, padres)

            if ciclo is not None:
                return ciclo

    return None


def recons_ciclo(padres, ini, fin):
    v = fin
    camino = []

    while v != ini:
        camino.append(v)
        v = padres[v]

    camino.append(ini)

    return camino
"""