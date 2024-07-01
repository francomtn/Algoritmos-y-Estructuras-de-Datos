package guia0

import (
	"bufio"
	"os"
)

func esCaracterValido(caracter int) bool {
	return caracter >= 'A' && caracter <= 'Z' || caracter >= 'a' && caracter <= 'z' || caracter >= '0' && caracter <= '9'
}

func AlfanumericoEnArchivo(ruta string) int {
	archivo, err := os.Open(ruta)
	if err != nil {
		return 0
	}
	defer archivo.Close()

	var contador int
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		for _, caracter := range linea {
			if esCaracterValido(int(caracter)) {
				contador++
			}
		}
	}

	err = s.Err()
	if err != nil {
		return 0
	}

	return contador
}
