package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDACola "tdas/cola"
	ABB "tdas/diccionario"
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
	colas := make([]TDACola.Cola[IPs], 255)
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

func radixSort(arr []IPs) {
	for i := 3; i >= 0; i-- {
		arr = countingSort(arr, i)
	}
}

func imprimirIPs(ip []IPs) {
	for _, ip := range ip {
		fmt.Println(ip.campos[0], ".", ip.campos[1], ".", ip.campos[2], ".", ip.campos[3])
	}
}

func agregarArchivo(ruta string, urls []string, visitantes ABB.DiccionarioOrdenado[IPs, string]) {
	archivo := abrirArchivo(ruta)
	scannerArchivo := bufio.NewScanner(archivo)
	dosDetectados := make([]IPs, 0)
	dictReportes := Reportes.CrearDiccionarioReportes()

	for scannerArchivo.Scan() {
		procesarLinea(scannerArchivo.Text(), dosDetectados, dictReportes, urls, visitantes)
	}

	radixSort(dosDetectados)

	imprimirIPs(dosDetectados)

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

func procesarLinea(linea string, dosDetectados []IPs, dictReportes Reportes.DiccionarioReportes, urls []string, visitantes ABB.DiccionarioOrdenado[IPs, string]) {
	elementos := splitLinea(linea)
	ip := elementos[IP]
	tiempo, err := time.Parse(LAYOUT, elementos[TIEMPO])

	if err != nil {
		panic(err)
	}

	ipParseada := parsearIp(ip)

	if dictReportes.Verificar(ip, tiempo) {
		dosDetectados = append(dosDetectados, ipParseada)
	}

	visitantes.Guardar(ipParseada, ip)

	urls = append(urls, elementos[URL])

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

func main() {
	urls := make([]string, 0)
	visitantes := ABB.CrearABB[IPs, string](funcionComparacion)
	scannerStdin := bufio.NewScanner(os.Stdin)
	for scannerStdin.Scan() {
		args := splitStdin(scannerStdin.Text())
		switch args[0] {
		case AGREGAR_ARCHIVO:
			agregarArchivo(args[2], urls, visitantes)
		case VER_MAS_VISITADOS:
			verMasVisitados(args[2])
		case VER_VISITANTES:
			verVisitantes(args[2], args[3])
		default:
			fmt.Println("Comando no reconocido")
		}
	}
}
