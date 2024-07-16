from collections import deque
from tdas.grafo import Grafo

def abecedario_alien(palabras):
    res = []
    g = armar_grafo(palabras)
    cola = deque()
    grados = {v: 0 for v in g.obtener_vertices()}

    for v in g.obtener_vertices():
        for w in g.adyacentes(v):
            grados[w] += 1

    for v in g.obtener_vertices():
        if grados[v] == 0:
            cola.append(v)

    while cola:
        v = cola.popleft()
        if v not in res:
            res.append(v)
        for w in g.adyacentes(v):
            grados[w] -= 1
            if grados[w] == 0:
                cola.append(w)
    return res

def armar_grafo(palabras):
    g = Grafo(es_dirigido=True)

    for palabra in palabras:
        for carac in palabra:
            if carac not in g:
                g.agregar_vertice(carac)

    for i in range(len(palabras) - 1):
        palabra1 = palabras[i]
        palabra2 = palabras[i + 1]
        min_len = min(len(palabra1), len(palabra2))

        for j in range(min_len):
            if palabra1[j] != palabra2[j]:
                if not g.estan_unidos(palabra1[j], palabra2[j]):
                    g.agregar_arista(palabra1[j], palabra2[j])
                break

    return g
