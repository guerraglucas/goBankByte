package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	printIntro()
	contaCorrente := readAccountData()
	titular, agencia, conta, saldo := contaCorrente.titular, contaCorrente.agencia, contaCorrente.conta, contaCorrente.saldo

	fmt.Println("Titular:", titular)
	fmt.Println("Agência:", agencia)
	fmt.Println("Conta:", conta)
	fmt.Println("Saldo:", saldo)

}

func printIntro() {
	fmt.Println("Bem vindo ao Bytebank")
	fmt.Println("Informe os dados da conta")
}

func readAccountData() ContaCorrente {
	var contaCorrente *ContaCorrente = new(ContaCorrente)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Informe o titular da conta:")
	titular, _ := reader.ReadString('\n')
	contaCorrente.titular = titular
	fmt.Println("Informe o número da agência:")
	fmt.Scan(&contaCorrente.agencia)
	fmt.Println("Informe o número da conta:")
	fmt.Scan(&contaCorrente.conta)
	fmt.Println("Informe o saldo da conta:")
	fmt.Scan(&contaCorrente.saldo)

	return *contaCorrente

}

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}
