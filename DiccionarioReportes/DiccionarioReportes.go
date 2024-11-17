package diccionarioreportes

import "time"

type DiccionarioReportes interface {
	Verificar(ip string, tiempo time.Time) bool
}
