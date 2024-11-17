package comandos

import (
	"bufio"
	"os"
	Diccionario "tdas/diccionario"
	"time"
	Reportes "tp2/DiccionarioReportes"
	aux "tp2/auxiliares"
)

const (
	_IP                            = 0
	_TIEMPO                        = 1
	_URL                           = 3
	_LAYOUT                        = "2006-01-02T15:04:05-07:00"
	_ARIDAD_AGREGAR_ARCHIVO        = 2
	_MENSAJE_ERROR_AGREGAR_ARCHIVO = "Error en comando agregar_archivo\n"
)

func abrirArchivo(ruta string) (*os.File, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	return archivo, nil
}

func procesarArchivo(archivo *os.File, urls Diccionario.Diccionario[string, int], visitantes Diccionario.DiccionarioOrdenado[aux.IPs, string]) {
	scanner := bufio.NewScanner(archivo)

	dosDetectados := Diccionario.CrearHash[string, bool]()
	dictReportes := Reportes.CrearDiccionarioReportes()

	for scanner.Scan() {
		linea := scanner.Text()
		procesarLinea(linea, dosDetectados, dictReportes, urls, visitantes)
	}

	ipsOrdenadas := aux.RadixSort(aux.HashIpToArr(dosDetectados))
	aux.ImprimirIPs(ipsOrdenadas)

}

func procesarLinea(linea string, dosDetectados Diccionario.Diccionario[string, bool], dictReportes Reportes.DiccionarioReportes, urls Diccionario.Diccionario[string, int], visitantes Diccionario.DiccionarioOrdenado[aux.IPs, string]) {
	elementos := aux.Split(linea, "\t")
	ip := elementos[_IP]

	tiempo, err := time.Parse(_LAYOUT, elementos[_TIEMPO])

	if aux.VerificarError(err, _MENSAJE_ERROR_AGREGAR_ARCHIVO) {
		return
	}

	ipParseada := aux.ParsearIp(ip)

	if dictReportes.Verificar(ip, tiempo) {
		if !dosDetectados.Pertenece(ip) {
			dosDetectados.Guardar(ip, true)
		}
	}

	visitantes.Guardar(ipParseada, ip)

	if urls.Pertenece(elementos[_URL]) {
		urls.Guardar(elementos[_URL], urls.Obtener(elementos[_URL])+1)
	} else {
		urls.Guardar(elementos[_URL], 1)
	}

}

func AgregarArchivo(args []string, urls Diccionario.Diccionario[string, int], visitantes Diccionario.DiccionarioOrdenado[aux.IPs, string]) {

	if aux.VerificarArgs(args, _ARIDAD_AGREGAR_ARCHIVO, _MENSAJE_ERROR_AGREGAR_ARCHIVO) {
		return
	}

	archivo, err := abrirArchivo(args[1])

	if aux.VerificarError(err, _MENSAJE_ERROR_AGREGAR_ARCHIVO) {
		return
	}

	defer archivo.Close()

	procesarArchivo(archivo, urls, visitantes)

	aux.PrintOK()

}
