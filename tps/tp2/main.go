package main

import (
	"bufio"
	"fmt"
	"os"
	lector "tp2/lector"
)

func main() {
	entrada := bufio.NewScanner(os.Stdin)
	lectorDs := lector.CrearLector()

	for entrada.Scan() {
		comando := entrada.Text()

		instruccion, resultado, err := lectorDs.Procesar(comando)
		if err != nil {
			msgError := "Error en comando " + instruccion
			fmt.Fprintln(os.Stderr, msgError)
		} else {
			switch instruccion {
			case "agregar_archivo":
				for _, ip := range resultado {
					fmt.Println("DoS: " + ip)
				}
				fmt.Println("OK")

			case "ver_visitantes":
				fmt.Println("Visitantes:")
				for _, ip := range resultado {
					fmt.Println("\t" + ip)
				}
				fmt.Println("OK")

			case "ver_mas_visitados":
				fmt.Println("Sitios más visitados:")
				for _, sitio := range resultado {
					if sitio != "" {
						fmt.Println("\t" + sitio)
					}
				}
				fmt.Println("OK")
			}
		}
	}

	if errEntrada := entrada.Err(); errEntrada != nil {
		fmt.Fprintln(os.Stderr, errEntrada)
	}
}
