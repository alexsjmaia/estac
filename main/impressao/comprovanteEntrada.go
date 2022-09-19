package impressao

import (
	"bytes"
	"os"
	"time"
	"estacionamento/database"
	"fmt"
	"log"
	"os/exec"
)

var (
	cmd *exec.Cmd
	out bytes.Buffer
	pla string
	modelo string
)

func ImprimirEntrada(pla, modelo string) {

	db, erro := database.Conectar()
	if erro != nil {
		fmt.Println("Erro ao Conectar no banco de dados")
		time.Sleep(time.Second * 600) // Parar o Programa aqui
	}
	defer db.Close()
	dataHoraEntrada := time.Now()
	dataHoraSaida := time.Now()

	statement, erro := db.Prepare("insert into entrada (placa, modelo, dataHoraEntrada, dataHoraSaida) values (?, ?, ?, ?)")

	if erro != nil {
		fmt.Println("Erro ao Criar o statement", erro)
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(pla, modelo, dataHoraEntrada.Format("02/01/2006 15:04:05"), dataHoraSaida.Format("02/01/2006 15:04:05"))
	if erro != nil {
		fmt.Println("Erro ao Montar a Inserção", insercao)
		fmt.Println(erro.Error())
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro  != nil {
		fmt.Println("Erro ao Obter o Numero da Entrada")
	}	

	file, erro := os.Create("pla.txt")
	if erro != nil {
		log.Fatalf("Falha ao Criar Arquivo Para Impressão", erro)
	}

	defer file.Close()

	len, erro := file.WriteString("===================================\n" +
		"\n===================================\n" +
		"\n * * * ESTACIONAMENTO SE MEGA PARK * * *\n\n" +
		"CNPJ : 10.809.909/0001-61\n" +
		"Rua Rio Bonito, 845 CEP : 03023-000\n\n" +
		"COMPROVANTE DE ENTRADA DE VEICULO\n\n" +
		"SAO PAULO " + time.Now().Format("02/01/2006 15:04:05\n\n") +
		//"ENTRADA : " + int64(idInserido) + "\n" +
		"ENTRADA : " + fmt.Sprintf("%d",idInserido) + "\n" +
		"DATA E HORA DE ENTRADA : " + time.Now().Format("02/01/2006 15:04:05\n\n") +
		"\tPLACA : " + pla + "\n" +
		//"\tMARCA : " + RegistraVeiculo.Marca + "" +
		"\tMODELO : " + modelo + "\n\n\n" +
		"****** AVISO IMPORTANTE ******\n" +
		"EM CASO DE CONVENIO, NAO ESQUECA\n" +
		"DE VALIDAR ESTE TICKET NO CAIXA\n" +
		"DA LOJA. ATT SE MEGA PARK\n\n" +
		"===================================\n" +
		"===================================\n\n\n\n\n\n" +
		"\n\n\n\n\n\n.")

	if erro != nil {
		log.Fatalf("Falha ao Criar Arquivo: %s", erro)
	}

	cmd = exec.Command("bash", "/usr/local/estac/main/pla.sh")
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		///////////////////////////////
		// Verificar o Retorno de Erros
		///////////////////////////////

		//fmt.Println(err, cmd) //exit status -1
		//log.Fatal(err)
		//time.Sleep(time.Second * 60)
		fmt.Println("Imprimindo...") // Esta Linha Não será necessária após verificar o retorno de erros
		return
	}

	fmt.Printf("\nTamanho do Arqivo : %d bytes", len)
	//time.Sleep(time.Second * 60)
}