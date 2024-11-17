package vermasvisitados

import (
	"fmt"
	"strconv"
	TDAColaPrioridad "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	aux "tp2/auxiliares"
)

const (
	_ARIDAD        = 1
	_MENSAJE_ERROR = "Error en comando ver_mas_visitados"
)

type urlVisitas struct {
	url  string
	cant int
}

func funcionComparacion(a, b urlVisitas) int {
	return a.cant - b.cant
}

func verMasVisitados(args []string, urls TDADiccionario.Diccionario[string, int]) {
	if aux.VerificarArgs(args, _ARIDAD, _MENSAJE_ERROR) {
		return
	}

	arr, cant_visitados := aux.arrADict(urls), args[0]
	k, _ := strconv.Atoi(cant_visitados)
	heap := TDAColaPrioridad.CrearHeapArr(arr, funcionComparacion)
	if k > heap.Cantidad() {
		k = heap.Cantidad()
	}

	for i := 0; i < k; i++ {
		elemento := heap.Desencolar()
		fmt.Printf("\t%s - %d\n", elemento.url, elemento.cant)
	}
	fmt.Println("OK")
}
