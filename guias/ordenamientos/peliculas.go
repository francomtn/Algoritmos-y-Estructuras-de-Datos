package ordenamientos

type Pelicula struct {
	nombre string
	genero string
}

func OrdernarPelis(arr []Pelicula) []Pelicula {

	buckets := make(map[string][]Pelicula)

	for _, peli := range arr {
		gen := peli.genero
		buckets[gen] = append(buckets[gen], peli)
	}
	mergedPelis := make([]Pelicula, len(arr))
	for _, bucket := range buckets {
		mergedPelis = append(mergedPelis, bucket...)
	}
	return mergedPelis
}

/*
map funciona parecido a un diccionario en python, es una estructura de datos que asocia claves con valores.
La clave del map es el género de la película, y el valor es una lista (slice) de películas que pertenecen a ese género
*/
