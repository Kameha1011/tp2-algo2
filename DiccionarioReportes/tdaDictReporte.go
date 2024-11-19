package diccionarioreportes

import (
	TDADiccionario "tdas/diccionario"
	"time"
)

type registroIp struct {
	tiempo    time.Time
	cantidad  int
	intervalo time.Duration
}

type dictReporte struct {
	hash TDADiccionario.Diccionario[string, registroIp]
}

func CrearDiccionarioReportes() DiccionarioReportes {
	hash := TDADiccionario.CrearHash[string, registroIp]()
	dictReportes := new(dictReporte)
	dictReportes.hash = hash
	return dictReportes
}

func (dicc *dictReporte) Verificar(ip string, tiempo time.Time) bool {
	if dicc.hash.Pertenece(ip) {
		registroIp := dicc.hash.Obtener(ip)
		registroIp.intervalo = tiempo.Sub(registroIp.tiempo) + registroIp.intervalo
		registroIp.tiempo = tiempo

		if !(registroIp.intervalo.Seconds() >= 2) {
			registroIp.cantidad++
			dicc.hash.Guardar(ip, registroIp)
			return registroIp.cantidad == 5
		}

		registroIp.cantidad = 1
		registroIp.intervalo = time.Duration(0)
		dicc.hash.Guardar(ip, registroIp)

	} else {
		registroIp := new(registroIp)
		registroIp.cantidad = 1
		registroIp.tiempo = tiempo
		registroIp.intervalo = time.Duration(0)
		dicc.hash.Guardar(ip, *registroIp)
	}

	return false
}
