package main

import (
	"bytes"
	"estacionamento/veiculos"
	"fmt"
	"os"
	"os/exec"
	"time"
)

var opcao = 0

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {

	for {
		// runCmd("cmd", "/c", "cls") // Isso é para Windows
		runCmd("clear") // Isso é Para Linux

		fmt.Println("\n\t\t\t\tGOLANG ESTACIONAMENTO\n")
		fmt.Print("\t1 - ENTRADA DE VEÍCULOS")
		fmt.Println("\t\t2 - SAÍDA DE VEÍCULOS")
		//fmt.Print("\t3 - REIMPRIME ÚLTIMA ENTRADA")
		//fmt.Println("\t4 - CADASTRA OPERADOR (a)")
		//fmt.Print("\t5 - CADASTRA MARCAS")
		//fmt.Println("\t\t6 - LISTA VEÍCULOS NO PATIO")
		//fmt.Print("\t7 - CANCELA ENTRADA DE VEICULO")
		//fmt.Println("\t8 - CADASTRA MENSALISTA")
		//fmt.Print("\t9 - CADASTRA PREÇO")

		fmt.Println("\n")
		fmt.Print("ROTINA : ")
		fmt.Scan(&opcao)

		if opcao == 1 {
			veiculos.Entrada()
			time.Sleep(time.Second * 3)

		} else if opcao == 2 {
			//veiculos.Saida()
			time.Sleep(time.Second * 3)

		} else if opcao == 3 {
			var (
				cmd *exec.Cmd
				out bytes.Buffer
			)
			cmd = exec.Command("cat", "pla.txt")
			cmd.Stdout = &out
			if err := cmd.Run(); err != nil {
				fmt.Println(err) //exit status -1
				return
			}
			fmt.Println(out.String())
			time.Sleep(time.Second * 3)

		} else if opcao == 4 {
			fmt.Println("Opção 4")
			//veiculos.CadastraOperador()
			time.Sleep(time.Second * 3)

		} else if opcao == 5 {
			//veiculos.CadastraModelo()
			time.Sleep(time.Second * 3)

		} else if opcao == 6 {
			//veiculos.ListaVeiculos()
			time.Sleep(time.Second * 3)

		} else if opcao == 7 {
			//veiculos.CancelaEntrada()
			time.Sleep(time.Second * 3)

		} else if opcao == 8 {
			fmt.Println("Opção 8")
			time.Sleep(time.Second * 3)

		} else if opcao == 9 {
			fmt.Println("Opção 9")
			time.Sleep(time.Second * 3)

		} else {
			fmt.Println("Opção Inválida!")
			time.Sleep(time.Second * 3)
		}
	}
}
