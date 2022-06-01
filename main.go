package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteA := clientes.Titular{
		Nome:      "Lucas",
		CPF:       "123.123.123-12",
		Profissao: "Estagiário",
	}

	contaA := contas.ContaCorrente{
		Titular:       clienteA,
		NumeroAgencia: 1234,
		NumeroConta:   123456,
		Saldo:         200,
	}

	fmt.Println(contaA)
}
