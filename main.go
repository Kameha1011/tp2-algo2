package main

import (
	"bufio"
	"fmt"
	"os"
	Diccionario "tdas/diccionario"
	aux "tp2/auxiliares"
	comandos "tp2/comandos"
)

const (
	AGREGAR_ARCHIVO   = "agregar_archivo"
	VER_MAS_VISITADOS = "ver_mas_visitados"
	VER_VISITANTES    = "ver_visitantes"
)

func main() {
	urls := Diccionario.CrearHash[string, int]()
	visitantes := Diccionario.CrearABB[aux.IPs, string](aux.FuncionComparacionAbb)
	scannerStdin := bufio.NewScanner(os.Stdin)
	for scannerStdin.Scan() {
		args := aux.Split(scannerStdin.Text(), " ")
		switch args[0] {
		case AGREGAR_ARCHIVO:

			comandos.AgregarArchivo(args, urls, visitantes)

		case VER_MAS_VISITADOS:

			comandos.VerMasVisitados(args, urls)

		case VER_VISITANTES:

			comandos.VerVisitantes(args, visitantes)

		default:
			fmt.Println("Comando no reconocido")
		}
	}
}
