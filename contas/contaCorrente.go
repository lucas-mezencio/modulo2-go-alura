package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "saque realizado"
	}
	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito <= 0 {
		return "Valor de depósito invalido", c.Saldo
	}
	c.Saldo += valorDoDeposito
	return "Depósito realizado com sucesso!", c.Saldo
}

// conta destino também precisa ser explicitado o uso do ponteiro
func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > c.Saldo && valorTransferencia <= 0 {
		return false
	}
	c.Saldo -= valorTransferencia
	contaDestino.Depositar(valorTransferencia)
	return true
}
