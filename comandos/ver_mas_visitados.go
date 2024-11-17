package comandos

import (
	"fmt"
	"strconv"
	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	aux "tp2/auxiliares"
)

const (
	_CANTIDAD                    = 1
	_ARIDAD_MAS_VISITADOS        = 2
	_MENSAJE_ERROR_MAS_VISITADOS = "Error en comando ver_mas_visitados\n"
)

func VerMasVisitados(args []string, urls TDADiccionario.Diccionario[string, int]) {

	if aux.VerificarArgs(args, _ARIDAD_MAS_VISITADOS, _MENSAJE_ERROR_MAS_VISITADOS) {

		return
	}

	fmt.Println("Sitios más visitados:")

	arr := aux.DictUrlsToArrElemHeap(urls)

	cant_visitados, _ := strconv.Atoi(args[_CANTIDAD])

	heap := TDAColaPrioridad.CrearHeapArr(arr, aux.FuncionComparacionHeap)
	if cant_visitados > heap.Cantidad() {
		cant_visitados = heap.Cantidad()
	}

	for i := 0; i < cant_visitados; i++ {
		elemento := heap.Desencolar()
		fmt.Printf("\t%s - %d\n", elemento.Url, elemento.Cant)
	}

	aux.PrintOK()
}
