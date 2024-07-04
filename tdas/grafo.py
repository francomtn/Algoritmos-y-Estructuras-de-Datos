import random
class Grafo:

    def __init__(self, dirigido = False, vertices=[]):

        self.dirigido = dirigido
        self.vertices = {v:{} for v in vertices}

    def obtener_vertices(self):
        return self.vertices.keys()
    
    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v]= {}

    def adyacentes(self, v):
        if v not in self.vertices:
            raise ValueError("La clave " + v + " no pertenece al grafo")
        return self.vertices[v].keys()
    
    def agregar_arista(self, v, w, peso=1):
        if v not in self.vertices:
            raise ValueError("La clave " + v + " no pertenece al grafo")
        if w not in self.vertices:
            raise ValueError("La clave " + w + " no pertenece al grafo")  
        self.vertices[v][w] = peso
        if not self.dirigido:
            self.vertices[w][v] = peso
    
    def peso_arista(self, v, w):
        if v not in self.vertices:
            raise ValueError("La clave " + v + " no pertenece al grafo")
        if w not in self.vertices:
            raise ValueError("La clave " + w + " no pertenece al grafo")       
        return self.vertices[v][w]
    
    def estan_unidos(self, v, w):
        if v not in self.vertices:
            raise ValueError("La clave " + v + " no pertenece al grafo")
        if w not in self.vertices:
            raise ValueError("La clave " + w + " no pertenece al grafo")  
        return w in self.vertices[v]
    
    def borrar_vertice(self, v):
        if v not in self.vertices:
            raise ValueError("La clave " + v + " no pertenece al grafo")
        self.vertices.pop(v)
        for vert in self.vertices:
            if v in vert:
                vert.pop(v)
    
    def borrar_arista(self,v,w):
        if v not in self.vertices:
            raise ValueError("La clave " + v + " no pertenece al grafo")
        if w not in self.vertices:
            raise ValueError("La clave " + w + " no pertenece al grafo")  
        self.vertices[v].pop(w)
        if not self.dirigido:
            self.vertices[w].pop(v)

    def vertice_aleatorio(self):
        return random.choice(self.vertices)
