package guiaTDA

import TDAPila "tdas/pila"

func InvertirArreglo[T any](arr []T) {

	pila := TDAPila.CrearPilaDinamica[T]()

	for i := 0; i < len(arr); i++ {
		pila.Apilar(arr[i])
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = pila.Desapilar()
	}

}

//el orden ejecucion es O(n²) porque a cada elemento del arreglo se lo apila
//en un ciclo y se lo desapila en otro, sobreescribiendo el valor del arreglo
//al que se lo asigna
