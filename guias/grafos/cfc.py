from pila import Pila
def cfcs(grafo):
    res = []
    visitados = set()
    for v in grafo:
        if v not in visitados:
            dfs_cfcs(grafo,v, visitados, {}, {}, Pila(), set(), res, [0])
    return res

def dfs_cfcs(grafo, v, visitados, orden, m_b, pila, apilados, cfcs, contador):

    orden[v] = m_b[v] = contador[0]
    contador[0]+=1
    visitados.add(v)
    apilados.add(v)

    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs_cfcs(grafo, w, visitados, orden, m_b, pila, apilados, cfcs, contador)
        if w in apilados:
            m_b[w] = min(m_b[w], m_b[v])
        if orden[v] == m_b[v]:
            nueva_cfc = []
            while True:
                w = pila.desapilar()
                apilados.remove(w)
                nueva_cfc.append(w)
                if w == v:
                    break
            cfcs.append(nueva_cfc)
            