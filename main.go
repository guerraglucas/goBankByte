package main

import (
	"aula-2-oop-go/accounts"
	"bufio"
	"fmt"
	"os"
)

func main() {
	printIntro()
	contaCorrente := accounts.ContaCorrente{Titular: "Lucas", Agencia: 123, Conta: 123, Saldo: 100}
	contaDestinoTransferencia := accounts.ContaCorrente{Titular: "Rhe", Agencia: 123, Conta: 123, Saldo: 100}

	printMenu()
	operationSelected := readOperationSelected()
	switch operationSelected {
	case 1:
		makeWithdraw(&contaCorrente)
	case 2:
		makeDeposit(&contaCorrente)
	case 3:
		makeTransfer(&contaCorrente, &contaDestinoTransferencia)
	case 4:
		fmt.Println("Obrigado por utilizar nossos serviços!")
		os.Exit(0)
	default:
		fmt.Println("Operação inválida!")
		os.Exit(-1)
	}

}

func printIntro() {
	fmt.Println("Bem vindo ao Bytebank")
	fmt.Println("Selecione a operação desejada:")
}
func printMenu() {
	fmt.Println("1 - Saque")
	fmt.Println("2 - Depósito")
	fmt.Println("3 - Transferência")
	fmt.Println("4 - Sair")
}

func readOperationSelected() int {
	var operationSelected int
	fmt.Scan(&operationSelected)
	return operationSelected
}

func readAccountData() *accounts.ContaCorrente {
	var contaCorrente *accounts.ContaCorrente = new(accounts.ContaCorrente)
	reader := bufio.NewReader(os.Stdin)
	objeto := *contaCorrente

	fmt.Println("Informe o titular da conta:")
	titular, _ := reader.ReadString('\n')
	objeto.Titular = titular

	fmt.Println("Informe o número da agência:")
	fmt.Scan(&objeto.Agencia)
	fmt.Println("Informe o número da conta:")
	fmt.Scan(&objeto.Conta)
	fmt.Println("Informe o saldo da conta:")
	fmt.Scan(&objeto.Saldo)

	return &objeto

}

func makeWithdraw(contaCorrente *accounts.ContaCorrente) {
	fmt.Println("Informe o valor do saque:")
	var saque float64
	fmt.Scan(&saque)
	contaCorrente.Withdraw(saque)
	fmt.Println("Saldo após saque:", contaCorrente.Saldo)
}

func makeDeposit(contaCorrente *accounts.ContaCorrente) {
	fmt.Println("Informe o valor do depósito:")
	var deposito float64
	fmt.Scan(&deposito)
	contaCorrente.Deposit(deposito)
	fmt.Println("Saldo após depósito:", contaCorrente.Saldo)
}

func makeTransfer(contaCorrente *accounts.ContaCorrente, contaDestino *accounts.ContaCorrente) {
	fmt.Println("Informe o valor da transferência:")
	var valor float64
	fmt.Scan(&valor)
	contaCorrente.Transfer(valor, contaDestino)
	fmt.Println("Saldo após transferência:", contaCorrente.Saldo)
}
