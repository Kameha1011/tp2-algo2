package auxiliares

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	Diccionario "tdas/diccionario"
)

/* Auxiliares generales */

type IPs struct {
	Campos [4]int
}

func ParsearIp(ip string) IPs {
	var ipParseada IPs
	elementos := strings.Split(ip, ".")
	for i := 0; i < 4; i++ {
		num, _ := strconv.Atoi(elementos[i])
		ipParseada.Campos[i] = num
	}
	return ipParseada
}

func VerificarArgs(args []string, aridad int, mensajeError string) bool {
	if len(args) != aridad {
		printError(mensajeError)
		return true
	}
	return false
}

func VerificarError(err error, mensajeError string) bool {
	if err != nil {
		printError(mensajeError)
		return true
	}
	return false
}

func printError(err string) {
	fmt.Fprint(os.Stderr, err)
}

func PrintOK() {
	fmt.Println("OK")
}

func Split(linea string, separador string) []string {
	return strings.Split(linea, separador)
}

/* Auxiliares de main */

func FuncionComparacionAbb(a, b IPs) int {
	for i := 0; i < 4; i++ {
		if a.Campos[i] == b.Campos[i] {
			continue
		}
		return a.Campos[i] - b.Campos[i]
	}
	return 0
}

/* Auxiliares de agregar_archivo */

func ImprimirIPs(ips []IPs) {
	for _, ip := range ips {
		fmt.Printf("DoS: %d.%d.%d.%d\n", ip.Campos[0], ip.Campos[1], ip.Campos[2], ip.Campos[3])
	}
}

func RadixSort(arr []IPs) []IPs {
	for i := 3; i >= 0; i-- {
		arr = countingSort(arr, i)
	}
	return arr
}

func countingSort(arr []IPs, campo int) []IPs {
	colas := make([]TDACola.Cola[IPs], 256)
	for i := range colas {
		colas[i] = TDACola.CrearColaEnlazada[IPs]()
	}

	for _, ip := range arr {
		colas[ip.Campos[campo]].Encolar(ip)
	}

	ordenadas := make([]IPs, len(arr))
	indice := 0
	for _, cola := range colas {
		for !cola.EstaVacia() {
			ordenadas[indice] = cola.Desencolar()
			indice++
		}
	}

	return ordenadas
}

func HashIpToArr(hash Diccionario.Diccionario[string, bool]) []IPs {
	arr := make([]IPs, 0)
	for iter := hash.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		ip := ParsearIp(clave)
		arr = append(arr, ip)
	}
	return arr
}

/* Auxiliares de ver_mas_visitados */

type UrlVisitadas struct {
	Url  string
	Cant int
}

func FuncionComparacionHeap(a, b UrlVisitadas) int {
	return a.Cant - b.Cant
}

func DictUrlsToArrElemHeap(urls Diccionario.Diccionario[string, int]) []UrlVisitadas {
	arr := make([]UrlVisitadas, 0)
	for iter := urls.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		arr = append(arr, UrlVisitadas{Url: clave, Cant: valor})
	}
	return arr
}

/* Auxiliares de ver_visitantes */
