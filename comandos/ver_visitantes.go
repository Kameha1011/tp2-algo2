package comandos

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	aux "tp2/auxiliares"
)

const (
	_INICIO                       = 1
	_FIN                          = 2
	_ARIDAD_VER_VISITANTES        = 3
	_MENSAJE_ERROR_VER_VISITANTES = "Error en comando ver_visitantes\n"
)

func VerVisitantes(args []string, visitantes TDADiccionario.DiccionarioOrdenado[aux.IPs, string]) {
	if aux.VerificarArgs(args, _ARIDAD_VER_VISITANTES, _MENSAJE_ERROR_VER_VISITANTES) {
		return

	}

	fmt.Println("Visitantes:")

	inicio, fin := args[_INICIO], args[_FIN]
	comenzar := aux.ParsearIp(inicio)
	final := aux.ParsearIp(fin)
	iterRango := visitantes.IteradorRango(&comenzar, &final)
	for iterRango.HaySiguiente() {
		_, ip := iterRango.VerActual()
		fmt.Printf("\t%s\n", ip)
		iterRango.Siguiente()
	}

	aux.PrintOK()
}
