package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaA := contas.ContaPoupanca{}
	contaA.Depositar(100)
	fmt.Println(contaA.GetSaldo())

	contaA.Sacar(50)
	fmt.Println(contaA.GetSaldo())

	fmt.Println(contaA)
}
