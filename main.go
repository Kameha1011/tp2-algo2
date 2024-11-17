package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	TDAColaPrioridad "tdas/cola_prioridad"
	Dict "tdas/diccionario"
	"time"
	Reportes "tp2/DictReportes"
)

const (
	AGREGAR_ARCHIVO   = "agregar_archivo"
	VER_MAS_VISITADOS = "ver_mas_visitados"
	VER_VISITANTES    = "ver_visitantes"
	LAYOUT            = "2006-01-02T15:04:05-07:00"
	IP                = 0
	TIEMPO            = 1
	METODO            = 2
	URL               = 3
)

type IPs struct {
	campos [4]int
}

func countingSort(arr []IPs, campo int) []IPs {
	colas := make([]TDACola.Cola[IPs], 256)
	for i := range colas {
		colas[i] = TDACola.CrearColaEnlazada[IPs]()
	}

	for _, ip := range arr {
		colas[ip.campos[campo]].Encolar(ip)
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

func radixSort(arr []IPs) []IPs {
	for i := 3; i >= 0; i-- {
		arr = countingSort(arr, i)
	}
	return arr
}

func imprimirIPs(ip []IPs) {
	for _, ip := range ip {
		fmt.Printf("DoS: %d.%d.%d.%d\n", ip.campos[0], ip.campos[1], ip.campos[2], ip.campos[3])
	}
}

func hashAArray(hash Dict.Diccionario[string, bool]) []IPs {
	arr := make([]IPs, 0)
	for iter := hash.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		ip := parsearIp(clave)
		arr = append(arr, ip)
	}
	return arr
}

func agregarArchivo(ruta string, urls Dict.Diccionario[string, int], visitantes Dict.DiccionarioOrdenado[IPs, string]) {
	archivo := abrirArchivo(ruta)
	scannerArchivo := bufio.NewScanner(archivo)
	dosDetectados := Dict.CrearHash[string, bool]()
	dictReportes := Reportes.CrearDiccionarioReportes()

	for scannerArchivo.Scan() {
		procesarLinea(scannerArchivo.Text(), dosDetectados, dictReportes, urls, visitantes)
	}

	ipsOrdenadas := radixSort(hashAArray(dosDetectados))

	imprimirIPs(ipsOrdenadas)

}

func parsearIp(ip string) IPs {
	var ipParseada IPs
	elementos := strings.Split(ip, ".")
	for i := 0; i < 4; i++ {
		num, _ := strconv.Atoi(elementos[i])
		ipParseada.campos[i] = num
	}
	return ipParseada
}

func procesarLinea(linea string, dosDetectados Dict.Diccionario[string, bool], dictReportes Reportes.DiccionarioReportes, urls Dict.Diccionario[string, int], visitantes Dict.DiccionarioOrdenado[IPs, string]) {
	elementos := splitLinea(linea)
	ip := elementos[IP]
	tiempo, err := time.Parse(LAYOUT, elementos[TIEMPO])
	if err != nil {
		panic(err)
	}

	ipParseada := parsearIp(ip)

	if dictReportes.Verificar(ip, tiempo) {
		if !dosDetectados.Pertenece(ip) {
			dosDetectados.Guardar(ip, true)
		}
	}

	visitantes.Guardar(ipParseada, ip)

	if urls.Pertenece(elementos[URL]) {
		urls.Guardar(elementos[URL], urls.Obtener(elementos[URL])+1)
	} else {
		urls.Guardar(elementos[URL], 1)
	}

}

func abrirArchivo(ruta string) *os.File {
	archivo, err := os.Open(ruta)
	if err != nil {
		panic(err)
	}
	return archivo
}

func splitLinea(linea string) []string {
	return strings.Split(linea, "\t")
}

func splitStdin(linea string) []string {
	return strings.Split(linea, " ")
}

var funcionComparacion = func(a, b IPs) int {
	for i := 0; i < 4; i++ {
		if a.campos[i] == b.campos[i] {
			continue
		}
		return a.campos[i] - b.campos[i]
	}
	return 0
}

type elementoHeap struct {
	url  string
	cant int
}

var funcionComparacionHeap = func(a, b elementoHeap) int {
	return a.cant - b.cant
}

func arrADict(urls Dict.Diccionario[string, int]) []elementoHeap {
	arr := make([]elementoHeap, 0)
	for iter := urls.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		arr = append(arr, elementoHeap{url: clave, cant: valor})
	}
	return arr
}

func verMasVisitados(cant_visitados string, urls Dict.Diccionario[string, int]) {
	arr := arrADict(urls)
	k, _ := strconv.Atoi(cant_visitados)
	heap := TDAColaPrioridad.CrearHeapArr(arr, funcionComparacionHeap)
	if k > heap.Cantidad() {
		k = heap.Cantidad()
	}
	for i := 0; i < k; i++ {
		elemento := heap.Desencolar()
		fmt.Printf("\t%s - %d\n", elemento.url, elemento.cant)
	}

}

func verVisitantes(inicio string, fin string, visitantes Dict.DiccionarioOrdenado[IPs, string]) {
	comenzar := parsearIp(inicio)
	final := parsearIp(fin)
	iterRango := visitantes.IteradorRango(&comenzar, &final)
	for iterRango.HaySiguiente() {
		_, ip := iterRango.VerActual()
		fmt.Printf("\t%s\n", ip)
		iterRango.Siguiente()
	}
}

func main() {
	urls := Dict.CrearHash[string, int]()
	visitantes := Dict.CrearABB[IPs, string](funcionComparacion)
	scannerStdin := bufio.NewScanner(os.Stdin)
	for scannerStdin.Scan() {
		args := splitStdin(scannerStdin.Text())
		switch args[0] {
		case AGREGAR_ARCHIVO:
			if len(args) != 2 {
				fmt.Fprint(os.Stderr, "Error en comando agregar_archivo\n")
				break
			}
			_, err := os.Open(args[1])
			if err != nil {
				fmt.Fprint(os.Stderr, "Error en comando agregar_archivo\n")
				break
			}
			agregarArchivo(args[1], urls, visitantes)
			fmt.Println("OK")
		case VER_MAS_VISITADOS:
			if len(args) != 2 {
				fmt.Fprint(os.Stderr, "Error en comando ver_mas_visitados\n")
				break
			}
			fmt.Println("Sitios más visitados:")
			verMasVisitados(args[1], urls)
			fmt.Println("OK")
		case VER_VISITANTES:
			if len(args) != 3 {
				fmt.Fprint(os.Stderr, "Error en comando ver_visitantes\n")
				break
			}
			fmt.Println("Visitantes:")
			verVisitantes(args[1], args[2], visitantes)
			fmt.Println("OK")
		default:
			fmt.Println("Comando no reconocido")
		}
	}
}
