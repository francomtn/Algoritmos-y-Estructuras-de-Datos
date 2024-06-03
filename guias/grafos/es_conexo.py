from cola import Cola


def es_conexo(grafo):

    if len(grafo) == 0:
        return True

    visitados = set()
    cola = Cola()
    datos = list(k for k in grafo)
    cola.encolar(datos[0])
    while not cola.esta_vacia():
        v = cola.desencolar()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                cola.encolar(w)
                visitados.add(w)
    return len(visitados) == len(grafo)

