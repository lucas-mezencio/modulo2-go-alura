package main

import (
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	contaA := contas.ContaPoupanca{}
	contaA.Depositar(100)
	pagarBoleto(&contaA, 60)

	fmt.Println(contaA.GetSaldo())

	contaB := contas.ContaCorrente{}
	contaB.Depositar(500)
	pagarBoleto(&contaB, 150)
	fmt.Println(contaB.GetSaldo())
}
