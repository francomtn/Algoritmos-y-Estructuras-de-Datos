package diccionarios

import "tdas/diccionario"

func EsIsomorfismo(cad1, cad2 string) bool {

	if len(cad1) != len(cad2) {
		return false
	}

	transformada := diccionario.CrearHash[string, string]()
	aux := diccionario.CrearHash[string, string]()

	for i := 0; i < len(cad1); i++ {
		c1 := string(cad1[i])
		c2 := string(cad2[i])

		if !transformada.Pertenece(c1) {
			transformada.Guardar(c1, c2)
			if aux.Pertenece(c2) {
				if aux.Obtener(c2) != c1 {
					return false
				}
			} else {
				aux.Guardar(c2, c1)
			}
		} else {
			if transformada.Obtener(c1) != c2 {
				return false
			}
		}
	}

	return true
}
