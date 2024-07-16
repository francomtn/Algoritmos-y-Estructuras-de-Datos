def encontrar_ciclo(grafo):
    '''
    Devuelve una lista de vertices que conforman el ciclo. 
    Si no hay ciclo, debe devolver None. 
    '''
    for v in grafo:
        visitados = set()
        padre = {}
        if v not in visitados:
            if v not in padre:
                padre[v] = None
            ciclo = dfs_ciclo(grafo, v, visitados, padre)
            if ciclo:
                return ciclo
    return None


def dfs_ciclo(grafo, v, visitados, padre):
    visitados.add(v)
    for w in grafo.adyacentes(v):
        if w in visitados:
            return reconstruir_ciclo(padre, w, v)
        else:
            padre[w] = v
            ciclo = dfs_ciclo(grafo, w, visitados, padre)
            if ciclo is not None:
                return ciclo
    return None

def reconstruir_ciclo(padre, inicio, fin):
    v = fin
    camino = []
    while v != inicio:
        camino.append(v)
        v = padre[v]
    camino.append(inicio)
    return camino[::-1]