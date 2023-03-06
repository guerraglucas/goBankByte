// This is the main package of a program that prints out information about a
//bank account. It imports the packages "bufio", "fmt" and "os". The main
//function calls the printIntro() and readAccountData() functions, then prints
// out the account data (titular, agencia, conta and saldo) as well as the
//value of a pointer and variable. Finally, it prints out the address of both
//the pointer and variable. The readAccountData() function reads in data from
//the user (titular, agencia, conta and saldo) using bufio.NewReader(os.Stdin)
// and stores them in a ContaCorrente struct type.

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
	var valor int = 45
	var ptr *int = &valor

	fmt.Println("valor do ponteiro: ", *ptr)
	fmt.Println("valor da variavel: ", valor)

	fmt.Println("endereco do ponteiro: ", &ptr)
	fmt.Println("endereco da variavel: ", &valor)
	fmt.Println("Titular:", titular)
	fmt.Println("Agência:", agencia)
	fmt.Println("Conta:", conta)
	fmt.Println("Saldo:", saldo)

	makeOperation(contaCorrente)

}

func printIntro() {
	fmt.Println("Bem vindo ao Bytebank")
	fmt.Println("Informe os dados da conta")
}

func readAccountData() *ContaCorrente {
	var contaCorrente *ContaCorrente = new(ContaCorrente)
	reader := bufio.NewReader(os.Stdin)
	// Aqui eu estou passando o endereço de memória da variável contaCorrente
	objeto := *contaCorrente

	fmt.Println("Informe o titular da conta:")
	titular, _ := reader.ReadString('\n')
	objeto.titular = titular

	fmt.Println("Informe o número da agência:")
	fmt.Scan(&objeto.agencia)
	fmt.Println("Informe o número da conta:")
	fmt.Scan(&objeto.conta)
	fmt.Println("Informe o saldo da conta:")
	fmt.Scan(&objeto.saldo)

	return &objeto

}

// This function is a method of the ContaCorrente struct. It takes in a
// float64 parameter called saque and subtracts it from the saldo field
//of the struct. It then returns the new value of the saldo field.

func makeOperation(contaCorrente *ContaCorrente) {
	fmt.Println("Informe o valor do saque:")
	var saque float64
	fmt.Scan(&saque)
	contaCorrente.Withdraw(saque)
	fmt.Println("Saldo após saque:", contaCorrente.saldo)
}

type ContaCorrente struct {
	titular string
	agencia int
	conta   int
	saldo   float64
}

func (r *ContaCorrente) Withdraw(saque float64) float64 {
	if r.ValidateWithdraw(saque) {
		r.saldo = r.saldo - saque
	} else {
		fmt.Println("Não foi possível realizar o saque")
	}
	return r.saldo
}

func (r *ContaCorrente) ValidateWithdraw(saque float64) bool {
	if saque > r.saldo && saque < 0 {
		return false
	}
	return true
}

func (r *ContaCorrente) Deposit(deposito float64) float64 {
	if r.ValidateDeposit(deposito) {
		r.saldo = r.saldo + deposito
	} else {
		fmt.Println("Não foi possível realizar o depósito")
	}
	return r.saldo
}

func (r *ContaCorrente) ValidateDeposit(deposito float64) bool {
	if deposito < 0 {
		return false
	}
	return true
}

func (r *ContaCorrente) Transfer(valor float64, contaDestino *ContaCorrente) {
	if r.ValidateWithdraw(valor) {
		r.Withdraw(valor)
		contaDestino.Deposit(valor)
	} else {
		fmt.Println("Não foi possível realizar a transferência")
	}
}
