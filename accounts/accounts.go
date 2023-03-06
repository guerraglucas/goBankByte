package accounts

import (
	u "aula-2-oop-go/users"
	"fmt"
)

type ContaCorrente struct {
	Titular u.Owner
	Agencia int
	Conta   int
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

func (r *ContaCorrente) GetBalance() float64 {
	return r.saldo
}
