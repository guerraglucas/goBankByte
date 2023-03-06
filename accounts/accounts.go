package accounts

import (
	u "aula-2-oop-go/users"
	"fmt"
)

type ContaCorrente struct {
	Titular u.Owner
	Agencia int
	Conta   int
	Saldo   float64
}

func (r *ContaCorrente) Withdraw(saque float64) float64 {
	if r.ValidateWithdraw(saque) {
		r.Saldo = r.Saldo - saque
	} else {
		fmt.Println("Não foi possível realizar o saque")
	}
	return r.Saldo
}

func (r *ContaCorrente) ValidateWithdraw(saque float64) bool {
	if saque > r.Saldo && saque < 0 {
		return false
	}
	return true
}

func (r *ContaCorrente) Deposit(deposito float64) float64 {
	if r.ValidateDeposit(deposito) {
		r.Saldo = r.Saldo + deposito
	} else {
		fmt.Println("Não foi possível realizar o depósito")
	}
	return r.Saldo
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
