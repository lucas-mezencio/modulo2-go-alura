package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo -= valorDoSaque
		return "saque realizado"
	}
	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito <= 0 {
		return "Valor de depósito invalido", c.saldo
	}
	c.saldo += valorDoDeposito
	return "Depósito realizado com sucesso!", c.saldo
}

// conta destino também precisa ser explicitado o uso do ponteiro
func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > c.saldo && valorTransferencia <= 0 {
		return false
	}
	c.saldo -= valorTransferencia
	contaDestino.Depositar(valorTransferencia)
	return true
}
