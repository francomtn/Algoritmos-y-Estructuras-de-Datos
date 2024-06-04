from collections import deque


def es_arbol(g):

    return es_conexo(g) and cant_aristas(g) == len(g)-1

def es_conexo(g):

    if len(g)== 0:
        return True
    visitados = set()
    cola = deque()
    datos = [k for k in g]
    cola.append(datos[0])
    while len(cola) > 0:
        v = cola.pop()
        for w in g.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                cola.append(w)
    return len(visitados) == len(g)

def cant_aristas(g):

    contador = 0
    for v in g:
        for w in g.adyacentes(v):
            contador += 1
    return contador // 2