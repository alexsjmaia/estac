package veiculos

import (
	"estacionamento/impressao"
	"fmt"
	"strings"
	"time"
  	//"estacionamento/database"
)

func Entrada() {

	var pla, modelo string

	fmt.Print("DIGITE A PLACA : ")
	fmt.Scan(&pla)
	pla = strings.ToUpper(pla)

	if len(pla) != 7 {
		fmt.Println("Placa Invalida")
		time.Sleep(time.Second * 3)
		Entrada()
	} else {
		fmt.Print("DIGITE O MODELO : ")
		fmt.Scan(&modelo)
		modelo = strings.ToUpper(modelo)

		impressao.ImprimirEntrada(pla, modelo)
	}
}
