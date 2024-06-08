from grafo import Grafo
from collections import deque


def abecedario_alien(palabras):
    '''
    palabras: lista de strings (palabras).
    Devolver lista de "letras"/caracteres en un orden válido. 
    '''

    res = []
    g = armar_grafo(palabras)
    cola = deque()
    grados = {}
    for v in g:
        for w in g.adyacentes(v):       
            grados[w] = grados.get(w, 0) + 1
    for v in g:
        if v not in grados:
            cola.append(v)
    while cola:
        v = cola.popleft()
        res.append(v)
        for w in g.adyacentes(v):
            grados[w] = grados[w] - 1
            if grados[w] == 0:
               cola.append(w)

    return res

def armar_grafo(palabras):

    g = Grafo(es_dirigido=True)

    for pal in palabras:
        for carac in pal:
            if carac not in g:
                g.agregar_vertice(carac)

    for i in range(len(palabras) - 1):
       pal1 = palabras[i]
       pal2 = palabras[i+1]

       for j in range(len(pal1)):
           if pal1[j] != pal2[j]:
               g.agregar_arista(pal1[j], pal2[j], 1)
               break     
    return g

    
