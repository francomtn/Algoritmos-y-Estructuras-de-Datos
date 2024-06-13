from grafos.cola import Cola
import heapq

"""
1. Implementar una función que reciba un grafo dirigido y acíclico y determine si dicho grafo admite un único orden
topológico. Indicar y justificar la complejidad de la función. Pista: pensar qué condición puede darse para que exista
más de un posible orden topológico.
"""

def obtener_orden_topologico(g):

    res = []
    grad = grados_entrada(g)
    cola = Cola()
    cola.encolar(g.vertice_aleatorio())
    while not cola.esta_vacia():
        v = cola.desencolar()
        res.append(v)
        for w in g.adyacentes(v):
            grad[w] -= 1
            if grad[w] == 0:
                cola.encolar(w)        
    return res

def grados_entrada(grafo):

    grados = {v : 0 for v in grafo.obtener_vertices()}

    for v in grafo.obtener_vertices():
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados

'''
Para que el el orden topologico sea unico, no debe haber mas de un vertice con grado de entrada 0
en cada iteracion.
la complejidad de las dos funciones en O(V + E) porque para cada vertice vemos sus adyacentes, por lo tanto 
la complejida total del algoritmos es 2 * O(V+E) -> O(V + E)
'''

"""
2. José Zubiría vive en una ciudad montañosa, es por eso que todos de los movimientos que realiza dentro de la ciudad
están en bajada o en subida. José desea encontrar el camino más rápido que le lleve desde su casa a la embajada
de su ciudad, sabiendo que le toma 3 minutos recorrer cada tramo en bajada, y le toma exactamente el doble, 6
minutos recorrerlo en subida. Actualmente cuenta con un detallado mapa de la ciudad, donde para cada tramo entre
dos esquinas tiene identificado si se encuentra en bajada o en subida. Asegurarse que el algoritmo implementado tenga
un comportamiento lineal respecto a la cantidad de tramos y esquinas que existan en la ciudad.
a. Modelar el problema usando Grafos especificando de forma completa todos sus elementos, para ello, describir
detalladamente con palabras de qué manera convertimos la información que tenemos en un grafo. Incluir algún
dibujo esquematizando el modelo.
b. Implementar un algoritmo que, dado el grafo modelado y la restricción temporal, solucione el problema de José,
indicando y justificando la complejidad final. 
"""

def camino_minimo_dijkstra(g, origen, destino):

    padre , distancia = {}, {v: float("inf") for v in g}
    distancia[origen] = 0
    cola = [(origen, 0)]
    while cola:
        v, _ = heapq.heappop(cola)
        if v == destino:
            return distancia[destino]
        for w in g.adyacentes(v):
            peso = g.peso_arista(v, w)
            if distancia[v] + peso < distancia[w]:
                distancia[w] = distancia[v] + peso
                padre[w] = v
                heapq.heappush(cola, (distancia[w], w))
    return distancia[destino]


"""
3. Se tiene un Grafo no dirigido G, y un árbol de tendido mínimo T, de G. Se obtiene una de las aristas de G, que no
se encuentra en T, y se le reduce el peso. Dar un algoritmo lineal que permita determinar el nuevo árbol de tendido
mínimo T'(notar que T' podría ser igual a T). Justificar la complejidad del algoritmo.
"""

def modificar_arbol(g, t):

    for v in g:
        for w in g.adyacentes(v):
            if not t.estan_unidos(v, w):
                peso = g.peso_arista(v, w)
                t.agregar_arista(v, w, peso - 1)
                return prim(t)

def prim(g):
    v = g.vertice_aleatorio()
    visitados = set()
    visitados.add(v)
    for w in g.adyacentes(v):
        cola = [((v, w), g.peso_arista(v, w))]
    arbol = Grafo(es_dirigido = False, vertices_init = g.obtener_vertices())
    while cola:

        (v, w), peso = heapq.heappop(cola)
        if w in visitados:
            continue
        arbol.agregar_arista(v, w, peso)
        visitados.add(w)
        for x in g.adyacentes(w):
            heapq.heapush(cola,((w, x), peso))
    return arbol
   
