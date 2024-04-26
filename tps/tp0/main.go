package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const archivo1 = "archivo1.in"
const archivo2 = "archivo2.in"

func leerArchivo(archivo string) ([]int, error) {
	arch, err := os.Open(archivo)
	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", arch, err)
		return nil, err
	}
	defer arch.Close()

	s := bufio.NewScanner(arch)
	arreglo := []int{}
	for s.Scan() {
		num, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Println(err)
		}
		arreglo = append(arreglo, num)
		err = s.Err()
		if err != nil {
			fmt.Println(err)
		}
	}
	return arreglo, nil
}

func main() {
	arreglo1, err := leerArchivo(archivo1)
	if err != nil {
		fmt.Println(err)
		return
	}
	arreglo2, err := leerArchivo(archivo2)
	if err != nil {
		fmt.Println(err)
		return
	}
	var mayorArreglo []int
	if ejercicios.Comparar(arreglo1, arreglo2) == 1 || ejercicios.Comparar(arreglo1, arreglo2) == 0 {
		mayorArreglo = arreglo1
	} else if ejercicios.Comparar(arreglo1, arreglo2) == -1 {
		mayorArreglo = arreglo2
	}
	ejercicios.Seleccion(mayorArreglo)
	for _, num := range mayorArreglo {
		fmt.Println(num)
	}
}
