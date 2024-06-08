def convertir_a_matriz(grafo):
    'Devolver la Matriz construida'
    'El arreglo de mapeo debe contener a todos los vértices, en el mismo orden en el que están representados en la matriz.'

    arreglo_mapeo = grafo.obtener_vertices()

    matriz = [[0 for _ in range(len(arreglo_mapeo))] for _ in range(len(arreglo_mapeo))]
    
    indices = {v: i for i, v in enumerate(arreglo_mapeo)}
    for v in grafo:
        for w in grafo.adyacentes(v):
            peso = grafo.peso_arista(v, w)
            i, j = indices[v], indices[w]
            matriz[i][j] = peso
            matriz[j][i] = peso

    return matriz, arreglo_mapeo    