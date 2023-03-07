package accounts

import (
	"fmt"

	u "github.com/guerraglucas/aula-2-oop-go/users"
)

type ContaPoupanca struct {
	Titular                  u.Owner
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (r *ContaPoupanca) Withdraw(saque float64) float64 {
	if r.ValidateWithdraw(saque) {
		r.saldo = r.saldo - saque
	} else {
		fmt.Println("Não foi possível realizar o saque")
	}
	return r.saldo
}

func (r *ContaPoupanca) ValidateWithdraw(saque float64) bool {
	if saque > r.saldo && saque < 0 {
		return false
	}
	return true
}

func (r *ContaPoupanca) Deposit(deposito float64) float64 {
	if r.ValidateDeposit(deposito) {
		r.saldo = r.saldo + deposito
	} else {
		fmt.Println("Não foi possível realizar o depósito")
	}
	return r.saldo
}

func (r *ContaPoupanca) ValidateDeposit(deposito float64) bool {
	return deposito >= 0
}

func (r *ContaPoupanca) GetBalance() float64 {
	return r.saldo
}
