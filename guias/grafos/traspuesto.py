from grafo import Grafo

def grafo_traspuesto(grafo):
    'Devolver el nuevo Grafo'
    nuevo = Grafo(es_dirigido=True)

    for v in grafo:
        nuevo.agregar_vertice(v)
    for v in grafo:
        for w in grafo.adyacentes(v):
            nuevo.agregar_arista(w, v, peso=1)
    return nuevo