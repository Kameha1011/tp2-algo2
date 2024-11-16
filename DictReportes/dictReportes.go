package tp2

import "time"

type DiccionarioReportes interface {
	Verificar(ip string, tiempo time.Time) bool
}
