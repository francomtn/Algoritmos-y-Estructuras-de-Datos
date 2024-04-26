package guiaTDA

import (
	"fmt"
	"math"
)

type Fraccion struct {
	numerador   int
	denominador int
}

// CrearFraccion crea una fraccion con el numerador y denominador indicados.
// Si el denominador es 0, entra en panico.
func CrearFraccion(numerador, denominador int) Fraccion {
	if denominador == 0 {
		panic("division por 0 no definida")
	}
	return Fraccion{numerador: numerador, denominador: denominador}
}

// Sumar crea una nueva fraccion, con el resultante de hacer la suma de las fracciones originales
func (f Fraccion) Sumar(otra Fraccion) Fraccion {
	nuevoNum := f.numerador*otra.denominador + f.denominador*otra.numerador
	nuevoDenom := f.denominador * otra.denominador
	nueva := CrearFraccion(nuevoNum, nuevoDenom)
	return nueva
}

// Multiplicar crea una nueva fraccion con el resultante de multiplicar ambas fracciones originales
func (f Fraccion) Multiplicar(otra Fraccion) Fraccion {
	nuevoNum := f.numerador * otra.numerador
	nuevoDenom := f.denominador * otra.denominador
	nueva := CrearFraccion(nuevoNum, nuevoDenom)
	return nueva
}

// ParteEntera devuelve la parte entera del numero representado por la fracción.
// Por ejemplo, para "7/2" = 3.5 debe devolver 3.
func (f Fraccion) ParteEntera() int {
	return f.numerador / f.denominador
}

// Representacion devuelve una representación en cadena de la fraccion simplificada (por ejemplo, no puede devolverse
// "10/8" sino que debe ser "5/4"). Considerar que si se trata de un número entero, debe mostrarse como tal.
// Considerar tambien el caso que se trate de un número negativo.
func (f Fraccion) Representacion() string {

	f.simplificar()

	var signo string
	if f.numerador*f.denominador < 0 {
		signo = "-"
		f.numerador = int(math.Abs(float64(f.numerador)))
	}
	if f.numerador%f.denominador == 0 {
		return fmt.Sprintf("%s%d", signo, f.numerador/f.denominador)
	}

	// Construir la cadena de representación de la fracción
	return fmt.Sprintf("%s%d/%d", signo, f.numerador, f.denominador)
}

func (f *Fraccion) simplificar() {

	div := mcd(f.numerador, f.denominador)
	f.numerador = f.numerador / div
	f.denominador = f.denominador / div
}
func mcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
