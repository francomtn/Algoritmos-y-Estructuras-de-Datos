from tdas.grafo import Grafo

def grafo_traspuesto(grafo):
    'Devolver el nuevo Grafo'
    nuevo = Grafo(es_dirigido=True, vertices_init= [v for v in grafo])

    for v in grafo:
        for w in grafo.adyacentes(v):
            nuevo.agregar_arista(w, v, peso=1)
    return nuevo