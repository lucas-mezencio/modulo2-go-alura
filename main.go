package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaA := ContaCorrente{}
	contaA.titular = "silvia"
	contaA.saldo = 500

	fmt.Println(contaA.saldo)

	valorDoSaque := 200.
	fmt.Println(contaA.Sacar(valorDoSaque))

	fmt.Println(contaA.saldo)
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "saque realizado"
	}
	return "saldo insuficiente"
}
