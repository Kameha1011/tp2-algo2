package tp2_test

import (
	"testing"
	"time"
	TDADiccionarioReoporte "tp2/DictReportes"
)

func TestVerificar(t *testing.T) {
	dictReportes := TDADiccionarioReoporte.CrearDiccionarioReportes()
	tiempo0, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:00+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo0)
	tiempo1, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:00+00:00")
	dictReportes.Verificar("66.249.73.185", tiempo1)
	tiempo2, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:03+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo2)
	tiempo3, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:03+00:00")
	dictReportes.Verificar("46.105.14.53", tiempo3)
	tiempo4, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:03+00:00")
	dictReportes.Verificar("110.136.166.128", tiempo4)
	tiempo5, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:04+00:00")
	dictReportes.Verificar("93.114.45.13", tiempo5)
	tiempo6, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:06+00:00")
	dictReportes.Verificar("110.136.166.128", tiempo6)
	tiempo7, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:07+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo7)
	tiempo8, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:08+00:00")
	dictReportes.Verificar("110.136.166.128", tiempo8)
	tiempo9, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:10+00:00")
	dictReportes.Verificar("50.16.19.13", tiempo9)
	tiempo10, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:11+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo10)
	tiempo11, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:11+00:00")
	dictReportes.Verificar("200.49.190.101", tiempo11)
	tiempo12, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:12+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo12)
	tiempo13, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:13+00:00")
	dictReportes.Verificar("81.220.24.207", tiempo13)
	tiempo14, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:14+00:00")
	dictReportes.Verificar("93.114.45.13", tiempo14)
	tiempo15, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:14+00:00")
	dictReportes.Verificar("93.114.45.13", tiempo15)
	tiempo16, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:15+00:00")
	dictReportes.Verificar("209.85.238.199", tiempo16)
	tiempo17, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:16+00:00")
	dictReportes.Verificar("66.249.73.135", tiempo17)
	tiempo18, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:17+00:00")
	dictReportes.Verificar("93.114.45.13", tiempo18)
	tiempo19, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:17+00:00")
	dictReportes.Verificar("66.249.73.135", tiempo19)
	tiempo20, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:18+00:00")
	dictReportes.Verificar("67.214.178.190", tiempo20)
	tiempo21, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:19+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo21)
	tiempo22, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:21+00:00")
	dictReportes.Verificar("93.114.45.13", tiempo22)
	tiempo23, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:21+00:00")
	dictReportes.Verificar("81.220.24.207", tiempo23)
	tiempo24, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:22+00:00")
	dictReportes.Verificar("91.177.205.119", tiempo24)
	tiempo25, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:22+00:00")
	dictReportes.Verificar("66.249.73.185", tiempo25)
	tiempo26, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:24+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo26)
	tiempo27, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:24+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo27)
	tiempo28, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:25+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo28)
	tiempo29, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:26+00:00")
	dictReportes.Verificar("81.220.24.207", tiempo29)
	tiempo30, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:28+00:00")
	dictReportes.Verificar("207.241.237.220", tiempo30)
	tiempo31, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:30+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo31)
	tiempo32, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:30+00:00")
	dictReportes.Verificar("209.85.238.199", tiempo32)
	tiempo33, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:31+00:00")
	dictReportes.Verificar("91.177.205.119", tiempo33)
	tiempo34, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:32+00:00")
	dictReportes.Verificar("110.136.166.128", tiempo34)
	tiempo35, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:32+00:00")
	dictReportes.Verificar("91.177.205.119", tiempo35)
	tiempo36, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:33+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo36)
	tiempo37, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:33+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo37)
	tiempo38, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:33+00:00")
	dictReportes.Verificar("66.249.73.135", tiempo38)
	tiempo39, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:34+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo39)
	tiempo40, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:34+00:00")
	dictReportes.Verificar("91.177.205.119", tiempo40)
	tiempo41, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:35+00:00")
	dictReportes.Verificar("110.136.166.128", tiempo41)
	tiempo42, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:36+00:00")
	dictReportes.Verificar("200.49.190.101", tiempo42)
	tiempo43, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:37+00:00")
	dictReportes.Verificar("66.249.73.185", tiempo43)
	tiempo44, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:37+00:00")
	dictReportes.Verificar("200.49.190.101", tiempo44)
	tiempo45, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:37+00:00")
	dictReportes.Verificar("91.177.205.119", tiempo45)
	tiempo46, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:38+00:00")
	dictReportes.Verificar("200.49.190.100", tiempo46)
	tiempo47, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:39+00:00")
	dictReportes.Verificar("81.220.24.207", tiempo47)
	tiempo48, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:40+00:00")
	dictReportes.Verificar("24.236.252.67", tiempo48)
	tiempo49, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:40+00:00")
	dictReportes.Verificar("66.249.73.135", tiempo49)
	tiempo50, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:40+00:00")
	dictReportes.Verificar("207.241.237.228", tiempo50)
	tiempo51, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:41+00:00")
	dictReportes.Verificar("110.136.166.128", tiempo51)
	tiempo52, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:43+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo52)
	tiempo53, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:44+00:00")
	dictReportes.Verificar("46.105.14.53", tiempo53)
	tiempo54, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:44+00:00")
	dictReportes.Verificar("81.220.24.207", tiempo54)
	tiempo55, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:45+00:00")
	dictReportes.Verificar("93.114.45.13", tiempo55)
	tiempo56, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:46+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo56)
	tiempo57, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:46+00:00")
	dictReportes.Verificar("123.125.71.35", tiempo57)
	tiempo58, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:46+00:00")
	dictReportes.Verificar("50.150.204.184", tiempo58)
	tiempo59, _ := time.Parse("2006-01-02T15:04:05-07:00", "2015-05-17T10:05:47+00:00")
	dictReportes.Verificar("83.149.9.216", tiempo59)

}
