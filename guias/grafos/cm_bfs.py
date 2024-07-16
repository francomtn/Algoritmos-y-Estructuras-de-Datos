from cola import Cola

def camino_minimo_bfs(grafo, origen):
   
   distancia, padre, visitado = {v: float("inf") for v in grafo}, {}, {}
   distancia[origen] = 0
   padre[origen] = None
   visitado[origen] = True
   q = Cola()
   q.encolar(origen)
   while not q.esta_vacia():
       v = q.desencolar()
       for w in grafo.adyacentes(v):
           if (v not in visitado):
               distancia[w] += distancia[v] + 1
               padre[w] = v
               visitado[w] = True
               q.encolar(w)
   return padre, distancia

# O(V + E)