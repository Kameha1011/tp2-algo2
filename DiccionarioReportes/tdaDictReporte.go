package diccionarioreportes

import (
	TDADiccionario "tdas/diccionario"
	"time"
)

type valor struct {
	tiempo    time.Time
	cantidad  int
	intervalo time.Duration
}

type dictReporte struct {
	hash TDADiccionario.Diccionario[string, valor]
}

func CrearDiccionarioReportes() DiccionarioReportes {
	hash := TDADiccionario.CrearHash[string, valor]()
	dictReportes := new(dictReporte)
	dictReportes.hash = hash
	return dictReportes
}

func (dicc *dictReporte) Verificar(ip string, tiempo time.Time) bool {
	if dicc.hash.Pertenece(ip) {
		valor := dicc.hash.Obtener(ip)
		valor.intervalo = tiempo.Sub(valor.tiempo) + valor.intervalo
		valor.tiempo = tiempo

		if !(valor.intervalo.Seconds() >= 2) {
			valor.cantidad++
			dicc.hash.Guardar(ip, valor)
			return valor.cantidad == 5
		}

		valor.cantidad = 1
		valor.intervalo = time.Duration(0)
		dicc.hash.Guardar(ip, valor)

	} else {
		valor := new(valor)
		valor.cantidad = 1
		valor.tiempo = tiempo
		valor.intervalo = time.Duration(0)
		dicc.hash.Guardar(ip, *valor)
	}

	return false
}
