#!/usr/bin/env python3

from grafo import Grafo
import biblioteca as b
import sys

mensajes = "./data/mensajes.tsv"
mensajes_minimos = "./data/minimo.tsv"

def construir_grafo(archivo_datos):
    grafo_mensajes = Grafo(True)
    with open(archivo_datos, "r") as archivo:
        for linea in archivo:
            v, w = linea.strip().split("\t")
            if v not in grafo_mensajes.obtener_vertices():
                grafo_mensajes.agregar_vertice(v)
            if w not in grafo_mensajes.obtener_vertices():
                grafo_mensajes.agregar_vertice(w)
            grafo_mensajes.agregar_arista(v, w, 1)
    return grafo_mensajes

def parsear_comando(entrada):
    elementos = entrada.split(" ")
    return elementos

def min_seguimientos(grafo, origen, destino):
    camino_arr = b.camino_minimo(grafo, origen, destino)

    if not camino_arr:
        return "Seguimiento imposible"

    return " -> ".join(camino_arr)

def divulgar(grafo, vertice, n):
    res = b.divulgar(grafo, vertice, int(n))
    return ", ".join(res)

def mas_imp(grafo, cantidad):
    res = b.obtenerNMasCentrales(grafo, cantidad)
    return ", ".join(res)

def persecucion(grafo, vertices, k):
    res = b.caminos_mas_rapidos(grafo, vertices, k)
    return " -> ".join(res)

def divulgar_ciclo(grafo, vertice):
    res = b.ciclo_mas_corto(grafo, vertice)

    if not res:
        return"No se encontro recorrido"
    
    return " -> ".join(res)

def main():
    if len(sys.argv) < 1:
        sys.exit("Error: No se pasan todos los parámetros")

    archivo_datos = sys.argv[1]

    grafo = construir_grafo(archivo_datos)

    for entrada in sys.stdin:
        procesar_entrada(grafo, entrada.strip())

def procesar_entrada(grafo, entrada):

    comando, *elementos = parsear_comando(entrada)
    resultado = ""

    if comando == "min_seguimientos":
        origen, destino = elementos[0], elementos[1]
        resultado = min_seguimientos(grafo, origen, destino)
    
    elif comando == "cfc":
        cfcs = b.cfcs_grafo(grafo)
        for i, cfc in enumerate(cfcs):
            respuesta = f"CFC {i+1}: "
            vertices = ", ".join(cfc)
            print(respuesta + vertices)
    
    elif comando == "divulgar":
        resultado = divulgar(grafo, elementos[0], elementos[1])
    
    elif comando == "divulgar_ciclo":
        resultado = divulgar_ciclo(grafo, elementos[0])
    
    elif comando == "comunidades":
        comunidades = b.obtener_comunidades(grafo, elementos[0])
        for i, c in enumerate(comunidades):
            respuesta = f"Comunidad {i+1}: "
            vertices = ", ".join(c)
            print(respuesta + vertices)
    
    elif comando == "mas_imp":
        cantidad = int(elementos[0])
        resultado = mas_imp(grafo, cantidad)
    
    elif comando == "persecucion":
        k = int(elementos.pop())
        vertices = elementos[0].split(",")
        resultado = persecucion(grafo, vertices, k)
    
    else:
        print("Comando no valido.")
    
    if resultado:
        print(resultado)

if __name__ == "__main__":
    main()


