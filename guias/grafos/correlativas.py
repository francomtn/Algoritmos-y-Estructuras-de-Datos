from collections import deque
"""
3. Tenemos un plan de estudios que nos indica las correlatividades de las materias que debemos realizar. Suponer que no
hay electivas, ni correlativas por cantidad de créditos. Tenemos un alumno, al que llamaremos agus9900, que quiere
recibirse lo antes posible (es decir, en la mínima cantidad de cuatrimestres). Modelar este problema con grafos, e
implementar una función que reciba dicho grafo y devuelva una lista de listas, donde en la lista i diga qué materias hay
que cursar en el i-ésimo cuatrimestre, de tal manera de tomar la menor cantidad de cuatrimestres en recibirse. Por
supuesto, siempre debe suceder que para toda materia de la lista i, todas sus correlativas deben haberse cursado en
cuatrimestres anteriores. Pueden asumir que agus9900 es tan genio que aprobó todas las cursadas y todos los finales
(en el mismo cuatrimestre de haberlas cursado). Indicar y justificar la complejidad del algoritmo implementado en
función de la cantidad de materias del plan de estudios, y la cantidad de correlatividades.
"""

def correlativas(g):
    pass


def topologico_grados(grafo):
    g_entrada = grados_entrada(grafo)
    res = []
    cola = deque()
    for v in grafo:
        if g_entrada[v] == 0:
            cola.append(v)
    while cola:
        v = cola.popleft()
        res.append(v)
        for w in grafo.adyacentes(v):
            g_entrada[w] -= 1
            if g_entrada[w] == 0:
                cola.append(w)
    return res

def grados_entrada(grafo):
    grados = {v:0 for v in grafo}
    for v in grafo:
        for w in grafo.adyacentes(v):
            grados[w] += 1
    return grados
