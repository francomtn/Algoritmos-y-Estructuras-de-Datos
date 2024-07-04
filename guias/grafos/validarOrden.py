from tdas.grafo import Grafo
from collections import deque


'''
5. Adam es un historiador que está analizando los datos de entrevistas realizadas a miembros de una aldea cercana al
mar Báltico. Con estas recopilaciones busca aprender sobre personas que vivieron allí por los últimos 200 años. De las
entrevistas escuchó cosas sobre n personas (todas ellas ya fallecidas), a quienes llamaremos P1, P2, · · · , Pn. También
pudo obtener información sobre cuándo estas personas vivieron, en relación a otras. Se tienen datos de la forma:
• Para algún i y algún j, Pi falleció antes que Pj haya nacido; o
• Para algún i y algún j, hubo algún momento en el cual Pi y Pj estuvieron vivos al mismo tiempo.
Como es de esperar, dado que todo esto está basado en memorias, y a veces son historias que alguien le contó a
otra persona, pueden haber algunas inconsistencias. Adam necesita implementar un algoritmo que determine si la
información recolectada es al menos consistente (es decir, no hay contradicciones, y realmente pueden haber existido
estas n personas que cumplan con todo lo obtenido).
Implementar un algoritmo que determine si dados m datos recopilados (como los indicados), estos son consistentes.
Para esto, modelar antes el problema con grafos. Indicar y justificar la complejidad del algoritmo implementado (en
función de las variables del problema).'''

def orden_topologico(grafo):

    res = []
    g_entrada = grados_entrada(grafo)
    cola = deque()
    v = grafo.vertice_aleatorio()
    cola.append(v)
    while cola:
        v = cola.popleft()
        if g_entrada[v] == 0:
            res.append(v)
        for w in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                cola.append(w)
    return res

def grados_entrada(grafo):

    res = {}
    for v in grafo:
        for w in grafo.adyacentes(v):
            res[w] = res.get(res[w], 0) + 1
    return res

