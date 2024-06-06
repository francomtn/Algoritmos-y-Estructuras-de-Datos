from cola import Cola

def a_n_aristas(grafo, v, n):
    'Devolver una lista con los vértices que cumplen la propiedad'
    res = []
    visitados = {v:0}
    cola = Cola()
    cola.encolar((v, 0))

    while not cola.esta_vacia():
        vert, dist = cola.desencolar()
        if dist == n:
            res.append(vert)
        elif dist < n:
            for w in grafo.adyacentes(vert):
                if w not in visitados or visitados[w] > dist + 1:
                    visitados[w] = dist + 1
                    cola.encolar((w, dist+1))
    return res