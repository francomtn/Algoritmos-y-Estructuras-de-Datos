package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	Heap "tdas/cola_prioridad"
	Dic "tdas/diccionario"
	"time"
)

var arrInstrucciones = []string{"agregar_archivo", "ver_visitantes", "ver_mas_visitados"}

type sitiosVisitados struct {
	nombre   string
	cantidad int
}

type lector struct {
	instrucciones Dic.Diccionario[string, bool]
	ips           Dic.DiccionarioOrdenado[string, bool]
	sitios        Dic.Diccionario[string, int]
}

func CrearLector() lector {
	instrucciones := Dic.CrearHash[string, bool]()
	ips := Dic.CrearABB[string, bool](compararIps)
	sitios := Dic.CrearHash[string, int]()

	for _, instruccion := range arrInstrucciones {
		instrucciones.Guardar(instruccion, true)
	}

	return lector{
		instrucciones: instrucciones,
		ips:           ips,
		sitios:        sitios,
	}
}

func (l *lector) Procesar(comando string) (string, []string, error) {

	var resultado []string
	elementos := strings.Fields(comando)
	instruccion := elementos[0]

	if !l.instrucciones.Pertenece(instruccion) {
		return instruccion, resultado, fmt.Errorf("comando no valido")
	}

	if len(elementos) == 1 {
		return instruccion, resultado, fmt.Errorf("parámetros invalidos")
	}

	switch instruccion {
	case "agregar_archivo":
		nombreArchivo := elementos[1]

		//En caso que se agregue un archivo nuevo, se reinician los valores del struct
		ips := Dic.CrearABB[string, bool](compararIps)
		sitios := Dic.CrearHash[string, int]()
		l.ips, l.sitios = ips, sitios

		res, err := l.agregarArchivo(nombreArchivo)

		if err != nil {
			return instruccion, resultado, err
		}

		resultado = res

	case "ver_visitantes":
		desde, hasta := elementos[1], elementos[2]
		resultado = l.verVisitantes(desde, hasta)

	case "ver_mas_visitados":
		n, err := strconv.Atoi(elementos[1])
		if err != nil {
			return instruccion, resultado, err
		}
		resultado = l.verMasVisitados(n)
	}

	return instruccion, resultado, nil
}

func (l *lector) agregarArchivo(ruta string) ([]string, error) {

	var resDesordenado []string
	hashAuxiliar := Dic.CrearHash[string, bool]()

	entradas := Dic.CrearHash[string, Heap.ColaPrioridad[string]]()
	archivo, err := os.Open(ruta)

	if err != nil {
		return resDesordenado, err
	}

	defer archivo.Close()

	s := bufio.NewScanner(archivo)

	for s.Scan() {
		linea := strings.Split(s.Text(), "\t")
		ip, fecha, visitado := linea[0], linea[1], linea[3]

		//Guardo cantidad de veces que se visitó un sitio para sitios mas visitados
		l.guardarSitios(visitado)

		//Guardo Ips para Ver Visitantes
		if !l.ips.Pertenece(ip) {
			l.ips.Guardar(ip, true)
		}

		if !entradas.Pertenece(ip) {
			heapActual := Heap.CrearHeap(func(f1, f2 string) int {
				return obtenerDiferencia(f1, f2)
			})
			heapActual.Encolar(fecha)
			entradas.Guardar(ip, heapActual)
			continue
		}

		heapActual := entradas.Obtener(ip)
		tope := heapActual.VerMax()

		diferencia := obtenerDiferencia(tope, fecha)

		for diferencia >= 2 && !heapActual.EstaVacia() {
			tope = heapActual.VerMax()
			heapActual.Desencolar()
			diferencia = obtenerDiferencia(tope, fecha)
		}

		heapActual.Encolar(fecha)
		entradas.Guardar(ip, heapActual)

		if heapActual.Cantidad() == 5 {
			hashAuxiliar.Guardar(ip, true)
			entradas.Borrar(ip)
		}
	}

	for iter := hashAuxiliar.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		actual, _ := iter.VerActual()
		resDesordenado = append(resDesordenado, actual)
	}

	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}

	res := OrdenarArregloIps(resDesordenado)

	return res, nil
}

func (l *lector) guardarSitios(visitado string) {
	if !l.sitios.Pertenece(visitado) {
		l.sitios.Guardar(visitado, 1)
	} else {
		valorActual := l.sitios.Obtener(visitado)
		l.sitios.Guardar(visitado, valorActual+1)
	}
}

func (l *lector) verVisitantes(desde string, hasta string) []string {

	var resultado []string
	for iter := l.ips.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		resultado = append(resultado, clave)
	}
	return resultado
}

func (l *lector) verMasVisitados(n int) []string {
	return l.TopKStream(n)
}

func obtenerDiferencia(anterior, actual string) int {

	var layout = "2006-01-02T15:04:05-07:00"

	ant, _ := time.Parse(layout, anterior)
	act, _ := time.Parse(layout, actual)
	dif := act.Sub(ant)

	return int(dif.Seconds())
}

func compararIps(a, b string) int {

	arr1 := strings.Split(a, ".")
	arr2 := strings.Split(b, ".")
	nuevo1 := make([]int, len(arr1))
	nuevo2 := make([]int, len(arr2))

	for i := range arr1 {
		num, _ := strconv.Atoi(arr1[i])
		nuevo1[i] = num
		num, _ = strconv.Atoi(arr2[i])
		nuevo2[i] = num
	}
	for i := 0; i < 4; i++ {
		if nuevo1[i] < nuevo2[i] {
			return -1
		} else if nuevo1[i] > nuevo2[i] {
			return 1
		}
	}

	return 0
}

func (l *lector) TopKStream(k int) []string {

	sitiosArr := make([]sitiosVisitados, k)

	for iter := l.sitios.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		sitiosArr = append(sitiosArr, sitiosVisitados{nombre: clave, cantidad: valor})
	}

	cp := Heap.CrearHeapArr(sitiosArr[:k], func(s1, s2 sitiosVisitados) int {
		c1, c2 := s1.cantidad, s2.cantidad
		return c2 - c1
	})

	for _, elem := range sitiosArr[k:] {
		if elem.cantidad > cp.VerMax().cantidad {
			cp.Desencolar()
			cp.Encolar(elem)
		}
	}

	top := make([]string, k)

	for i := 0; !cp.EstaVacia(); i++ {
		tope := cp.Desencolar()
		sitio := tope.nombre
		cant := tope.cantidad

		if cant != 0 {
			top[k-i-1] = sitio + " - " + strconv.Itoa(cant)
		}

	}
	return top
}
func OrdenarArregloIps(arr []string) []string {
	countingSort(arr, 256, 3)
	countingSort(arr, 256, 2)
	countingSort(arr, 256, 1)
	countingSort(arr, 256, 0)
	return arr
}

func countingSort(elementos []string, rango int, digito int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]string, len(elementos))

	for _, elem := range elementos {
		datos := strings.Split(elem, ".")
		valor, _ := strconv.Atoi(datos[digito])
		frecuencias[valor]++
	}

	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}

	for _, elem := range elementos {
		datos := strings.Split(elem, ".")
		valor, _ := strconv.Atoi(datos[digito])
		pos := sumasAcumuladas[valor]
		resultado[pos] = elem
		sumasAcumuladas[valor]++
	}

	for i := 0; i < len(resultado); i++ {
		elementos[i] = resultado[i]
	}
}
