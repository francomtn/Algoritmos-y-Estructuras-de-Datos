def encontrar_ciclo(g):
    visitados = set()
    padres = {}

    for v in g:
        if v not in visitados:
            ciclo = dfs_ciclo(g, v, visitados, padres)

            if ciclo is not None:
                return ciclo

    return None


def dfs_ciclo(g, v, visitados, padres):
    visitados.add(v)

    for w in g.adyacentes(v):
        if w in visitados and w != padres[v]:
            return recons_ciclo(padres, w, v)

        if w not in visitados: 
            padres[w] = v
            ciclo = dfs_ciclo(g, w, visitados, padres)

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
