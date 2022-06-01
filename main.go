package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaA := contas.ContaCorrente{}
	contaA.Depositar(100)
	fmt.Println(contaA.GetSaldo())
}
