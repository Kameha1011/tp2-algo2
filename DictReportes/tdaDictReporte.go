package tp2

import (
	TDADiccionario "tdas/diccionario"
	"time"
)

type valor struct {
	tiempo_0 time.Time
	cant_0   int
	tiempo_1 time.Time
	cant_1   int
}

type DictReporte struct {
	hash TDADiccionario.Diccionario[string, valor]
}

// [192.168.3.2, 192.168.3.2, 192.168.3.2, 192.168.3.2,192.168.3.2, 192.168.3.2]

func CrearDiccionarioReportes() DiccionarioReportes {
	hash := TDADiccionario.CrearHash[string, valor]()
	dictReportes := new(DictReporte)
	dictReportes.hash = hash
	return dictReportes
}

func (dicc *DictReporte) Verificar(ip string, tiempo time.Time) bool {
	t := time.Time{}
	if dicc.hash.Pertenece(ip) {
		valor := dicc.hash.Obtener(ip)
		if valor.tiempo_0 == tiempo {
			valor.cant_0++
			dicc.hash.Guardar(ip, valor)
			if valor.cant_0 == 5 {
				return true
			}
		} else if valor.tiempo_1 == tiempo {
			valor.cant_1++
			dicc.hash.Guardar(ip, valor)
			if valor.cant_1 == 5 || (valor.cant_1+valor.cant_0 >= 5 && valor.tiempo_1.Sub(valor.tiempo_0).Abs().Seconds() < 2) {
				return true
			}
		} else if valor.tiempo_1 == t {
			valor.tiempo_1 = tiempo
			valor.cant_1 = 1
			if valor.cant_0+valor.cant_1 >= 5 && valor.tiempo_1.Sub(valor.tiempo_0).Abs().Seconds() < 2 {
				return true
			}
			dicc.hash.Guardar(ip, valor)
			return false
		} else {
			valor.tiempo_0 = valor.tiempo_1
			valor.cant_0 = valor.cant_1
			valor.tiempo_1 = tiempo
			valor.cant_1 = 1
			dicc.hash.Guardar(ip, valor)
			return false
		}
	} else {
		valor := new(valor)
		valor.cant_0 = 1
		valor.tiempo_0 = tiempo
		valor.cant_1 = 0
		valor.tiempo_1 = time.Time{}
		dicc.hash.Guardar(ip, *valor)
	}
	return false
}
