package main

import (
	a "aula-2-oop-go/accounts"
	u "aula-2-oop-go/users"
	"bufio"
	"fmt"
	"os"
)

func main() {
	printIntro()
	userLucas := u.Owner{Name: "Lucas", Age: 28, Email: "lucasguerra123@emai.com"}
	userRhe := u.Owner{Name: "Rhe", Age: 26, Email: "rhe123@emai.com"}

	contaCorrente := a.ContaCorrente{Titular: userLucas, Agencia: 123, Conta: 123, Saldo: 100}
	contaDestinoTransferencia := a.ContaCorrente{Titular: userRhe, Agencia: 123, Conta: 123, Saldo: 100}

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

func readAccountData() *a.ContaCorrente {
	var contaCorrente *a.ContaCorrente = new(a.ContaCorrente)
	reader := bufio.NewReader(os.Stdin)
	objeto := *contaCorrente

	fmt.Println("Informe o titular da conta:")
	titular, _ := reader.ReadString('\n')
	objeto.Titular.Name = titular

	fmt.Println("Informe o número da agência:")
	fmt.Scan(&objeto.Agencia)
	fmt.Println("Informe o número da conta:")
	fmt.Scan(&objeto.Conta)
	fmt.Println("Informe o saldo da conta:")
	fmt.Scan(&objeto.Saldo)

	return &objeto

}

func makeWithdraw(contaCorrente *a.ContaCorrente) {
	fmt.Println("Informe o valor do saque:")
	var saque float64
	fmt.Scan(&saque)
	contaCorrente.Withdraw(saque)
	fmt.Println("Saldo após saque:", contaCorrente.Saldo)
}

func makeDeposit(contaCorrente *a.ContaCorrente) {
	fmt.Println("Informe o valor do depósito:")
	var deposito float64
	fmt.Scan(&deposito)
	contaCorrente.Deposit(deposito)
	fmt.Println("Saldo após depósito:", contaCorrente.Saldo)
}

func makeTransfer(contaCorrente *a.ContaCorrente, contaDestino *a.ContaCorrente) {
	fmt.Println("Informe o valor da transferência:")
	var valor float64
	fmt.Scan(&valor)
	contaCorrente.Transfer(valor, contaDestino)
	fmt.Println("Saldo após transferência:", contaCorrente.Saldo)
}
