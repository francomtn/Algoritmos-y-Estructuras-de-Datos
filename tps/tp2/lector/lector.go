package tp2

type Lector interface {

	//Procesa de forma completa un archivo de log.
	Procesar(string) (string, error)
}
