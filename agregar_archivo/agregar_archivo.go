package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	Diccionario "tdas/diccionario"
	"time"
	Reportes "tp2/DiccionarioReportes"
	aux "tp2/auxiliares"
)

const (
	ARIDAD = 2
	ERROR  = "Error en comando agregar_archivo\n"
)

func verificarArgs(args []string, aridad int, mensajeError string) bool {
	if len(args) != aridad {
		printError(mensajeError)
		return true
	}
	return false
}

func verificarArchivo(err error, mensajeError string) bool {
	if err != nil {
		printError(mensajeError)
		return true
	}
	return false
}

func printError(err string) {
	fmt.Fprint(os.Stderr, err)
}

func abrirArchivo(ruta string) (*os.File, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	return archivo, nil
}

func hashToArr(hash Diccionario.Diccionario[string, bool]) []aux.IPs {
	arr := make([]aux.IPs, 0)
	for iter := hash.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		ip := aux.parsearIp(clave)
		arr = append(arr, ip)
	}
	return arr
}

func imprimirIPs(ips []aux.IPs) {
	for _, ip := range ips {
		fmt.Printf("DoS: %d.%d.%d.%d\n", ip.campos[0], ip.campos[1], ip.campos[2], ip.campos[3])
	}
}

func procesarArchivo(archivo *os.File, urls Diccionario.Diccionario[string, int], visitantes Diccionario.DiccionarioOrdenado[aux.IPs, string]) {
	scanner := bufio.NewScanner(archivo)

	dosDetectados := Diccionario.CrearHash[string, bool]()
	dictReportes := Reportes.CrearDiccionarioReportes()

	for scanner.Scan() {
		linea := scanner.Text()
		procesarLinea(linea, dosDetectados, dictReportes, urls, visitantes)
	}

	ipsOrdenadas := aux.RadixSort(hashToArr(dosDetectados))
	imprimirIPs(ipsOrdenadas)

}

func splitLinea(linea string) []string {
	return strings.Split(linea, "\t")
}

func procesarLinea(linea string, dosDetectados Diccionario.Diccionario[string, bool], dictReportes Reportes.DiccionarioReportes, urls Diccionario.Diccionario[string, int], visitantes Diccionario.DiccionarioOrdenado[aux.IPs, string]) {
	elementos := splitLinea(linea)
	ip := elementos[aux.IP]
	tiempo, err := time.Parse(aux.LAYOUT, elementos[aux.TIEMPO])
	if err != nil {
		panic(err)
	}

	ipParseada := aux.parsearIp(ip)

	if dictReportes.Verificar(ip, tiempo) {
		if !dosDetectados.Pertenece(ip) {
			dosDetectados.Guardar(ip, true)
		}
	}

	visitantes.Guardar(ipParseada, ip)

	if urls.Pertenece(elementos[aux.URL]) {
		urls.Guardar(elementos[aux.URL], urls.Obtener(elementos[aux.URL])+1)
	} else {
		urls.Guardar(elementos[aux.URL], 1)
	}

}

func AgregarArchivo(args []string, urls Diccionario.Diccionario[string, int], visitantes Diccionario.DiccionarioOrdenado[aux.IPs, string]) {

	if verificarArgs(args, ARIDAD, ERROR) {
		return
	}

	archivo, err := abrirArchivo(args[1])

	if verificarArchivo(err, ERROR) {
		return
	}

	defer archivo.Close()

}
