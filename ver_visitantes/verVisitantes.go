package vervisitantes

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	aux "tp2/auxiliares"
)

const (
	_ARIDAD        = 2
	_MENSAJE_ERROR = "Error en comando ver_visitantes"
)

func VerVisitantes(args []string, visitantes TDADiccionario.DiccionarioOrdenado[aux.IPs, string]) {
	if aux.VerificarArgs(args, _ARIDAD, _MENSAJE_ERROR) {
		return

	}

	inicio, fin := args[0], args[1]
	comenzar := aux.parsearIp(inicio)
	final := aux.parsearIp(fin)
	iterRango := visitantes.IteradorRango(&comenzar, &final)
	for iterRango.HaySiguiente() {
		_, ip := iterRango.VerActual()
		fmt.Printf("\t%s\n", ip)
		iterRango.Siguiente()
	}

	fmt.Println("OK")
}
