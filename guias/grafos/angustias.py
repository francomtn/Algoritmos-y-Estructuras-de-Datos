from tdas.grafo import Grafo
import heapq
"""
2. La señora Angustias busca peces en el lago para su pescadería. Para ello tiene marcadas las posibles ubicaciones de
cardúmenes de peces; y sigue el siguiente procedimiento:
(1) Tiene un papel en donde va a ir anotando la lista de ubicaciones ya visitadas, inicialmente vacía.
(2) se fija en el mapa cual es la ubicación más cercana a donde se encuentra que no esté en la lista y va a ella. Por
ejemplo, si desde donde está tiene posibles cardúmenes de peces restantes a 1.5km y a 1.132km, elige este último
(3) Si en la ubicación seleccionada se confirma la existencia del cardúmen, completa su búsqueda con éxito.
(4) En caso contrario, anota esa ubicación en la lista, y repite el proceso desde (2).
Dado el mapa del lago con las posibles ubicaciones de cardúmenes de peces:
a. Modelar el problema usando Grafos especificando de forma completa todos sus elementos, para ello, describir
detalladamente con palabras de qué manera convertimos la información que tenemos en un grafo.
b. Implementar un algoritmo que, dado el grafo modelado, replique el procedimiento de Angustias, indicando y
justificando la complejidad final.
"""

def angustias(g: Grafo, origen):

    visitados = set()

    while True:
        if origen == "cardumen":
            return origen
        
        padre, dist = dijkstra(g, origen) #posibles puntos
        no_visitados = {v: d for v, d in dist.items()}

        if not no_visitados:
            return "No se encontraron cardúmenes"
        
        mas_cercano = min(no_visitados, key=no_visitados.get)
        if "cardumen" == mas_cercano:
            return mas_cercano
        visitados.add(mas_cercano)
        origen = mas_cercano



def dijkstra(g, origen):

    cardumen = []
    padre, dist = {}, {v: float("inf") for v in g}
    padre[origen], dist[origen] = None, 0
    cola = [(0, origen)] # heap
    while cola:
        dist_v, v = heapq.heappop()
        if v == "cardumen":
            cardumen.append((dist_v, v))
        for w in g.adyacentes(v):
            dist_aux = dist_v + g.peso_arista(v, w)
            if dist_aux < dist[w]:
                dist[w] = dist_aux
                padre[w] = v
                heapq.heappush(cola, (dist[w], w))
    return padre, dist