package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAPila "tdas/pila"
	calcu "tp1/calculadora"
)

func identificarCaracteres(pilaNumeros TDAPila.Pila[int64], elementos []string) (TDAPila.Pila[int64], error) {
	for _, elem := range elementos {
		if esOperador(elem) {
			pilaNumeros, err := calcu.Calcular(pilaNumeros, elem)
			if err != nil {
				return pilaNumeros, err
			}
		} else {
			num, err := strconv.Atoi(elem)
			if err != nil {
				return pilaNumeros, err
			}
			pilaNumeros.Apilar(int64(num))
		}
	}
	return pilaNumeros, nil
}

func esOperador(cadena string) bool {
	operadores := [8]string{"+", "-", "*", "/", "sqrt", "^", "log", "?"}
	for _, elem := range operadores {
		if cadena == elem {
			return true
		}
	}
	return false
}

func vaciarPila(pila TDAPila.Pila[int64]) TDAPila.Pila[int64] {
	for !pila.EstaVacia() {
		pila.Desapilar()
	}
	return pila
}

func imprimirResultado(pila TDAPila.Pila[int64]) (int64, error) {
	if pila.EstaVacia() {
		return 0, fmt.Errorf("ERROR")
	}
	resultado := pila.Desapilar()
	if !pila.EstaVacia() {
		return 0, fmt.Errorf("ERROR") // Sobran operandos
	}
	return resultado, nil
}

func main() {
	pilaNumeros := TDAPila.CrearPilaDinamica[int64]()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		pilaNumeros = vaciarPila(pilaNumeros)
		linea := s.Text()
		elementos := strings.Fields(linea)
		pilaNumeros, err := identificarCaracteres(pilaNumeros, elementos)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			resultado, err2 := imprimirResultado(pilaNumeros)
			if err2 != nil {
				fmt.Println("ERROR")
			} else {
				fmt.Println(resultado)
			}
		}
	}
	if err := s.Err(); err != nil {
		fmt.Println("ERROR")
	}
}
