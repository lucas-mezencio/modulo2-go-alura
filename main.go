package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaA := contas.ContaCorrente{}
	contaA.Titular = "silvia"
	contaA.Saldo = 500

	contaB := contas.ContaCorrente{
		Titular: "carlos",
		Saldo:   400,
	}

	status := contaA.Transferir(300, &contaB)
	fmt.Println(status)
	status = contaA.Transferir(300, &contaB)
	fmt.Println(status)
}
