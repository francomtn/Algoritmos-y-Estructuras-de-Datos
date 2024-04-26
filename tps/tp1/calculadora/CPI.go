package calculadora

import (
	"fmt"
	"math"
	TDAPila "tdas/pila"
)

const (
	MINIMO_OPERANDO int64 = 0
	MINIMO_BASE     int64 = 2
)

func Calcular(numeros TDAPila.Pila[int64], operador string) (TDAPila.Pila[int64], error) {

	if numeros.EstaVacia() {
		return numeros, fmt.Errorf("ERROR: Faltan operandos")
	}
	switch operador {
	case "+", "-", "*", "/", "^", "log":
		return operarDosNumeros(numeros, operador)
	case "sqrt":
		return operadorRaiz(numeros)
	case "?":
		return operadorTernario(numeros)
	default:
		return numeros, nil
	}
}

func operarDosNumeros(pila TDAPila.Pila[int64], signo string) (TDAPila.Pila[int64], error) {

	if pila.EstaVacia() {
		return pila, fmt.Errorf("ERROR: Faltan operandos")
	}
	operando1 := pila.Desapilar()
	if pila.EstaVacia() {
		return pila, fmt.Errorf("ERROR: Faltan operandos")
	}
	operando2 := pila.Desapilar()
	switch signo {
	case "+":
		pila.Apilar(operando2 + operando1)

	case "-":
		pila.Apilar(operando2 - operando1)

	case "*":
		pila.Apilar(operando2 * operando1)

	case "/":
		if operando1 == MINIMO_OPERANDO {
			return pila, fmt.Errorf("ERROR: Division por 0")
		} else {
			pila.Apilar(operando2 / operando1)
		}

	case "^":
		if operando1 < MINIMO_OPERANDO {
			return pila, fmt.Errorf("ERROR: Exponente negativo")
		}
		pila.Apilar(int64(math.Pow(float64(operando2), float64(operando1))))

	case "log":
		if operando1 < MINIMO_BASE || operando2 < MINIMO_OPERANDO {
			return pila, fmt.Errorf("ERROR: La base debe ser mayor o igual a 2 y el argumento positivo")
		} else {
			aplicarLog := math.Log(float64(operando2)) / math.Log(float64(operando1))
			pila.Apilar(int64(aplicarLog))
		}
	}
	return pila, nil
}

func operadorRaiz(pila TDAPila.Pila[int64]) (TDAPila.Pila[int64], error) {

	if pila.EstaVacia() {
		return pila, fmt.Errorf("ERROR: Faltan operandos")
	}
	operando1 := pila.Desapilar()
	if operando1 < MINIMO_OPERANDO {
		return pila, fmt.Errorf("ERROR: Raiz menor que 0")
	}
	pila.Apilar(int64(math.Sqrt(float64(operando1))))
	return pila, nil
}

func operadorTernario(pila TDAPila.Pila[int64]) (TDAPila.Pila[int64], error) {

	if pila.EstaVacia() {
		return pila, fmt.Errorf("ERROR: Faltan operandos")
	}
	operando1 := pila.Desapilar()
	if pila.EstaVacia() {
		return pila, fmt.Errorf("ERROR: Faltan operandos")
	}
	operando2 := pila.Desapilar()
	if pila.EstaVacia() {
		return pila, fmt.Errorf("ERROR: Faltan operandos")
	}
	operando3 := pila.Desapilar()
	if operando3 == MINIMO_OPERANDO {
		pila.Apilar(operando1)
	} else {
		pila.Apilar(operando2)
	}
	return pila, nil
}
