package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	x := soma(1, 1)
	if x != 2 {
		t.Error("Erro ao somar")
	}
}

func TestImpressao(t *testing.T) {
	impressao_ := impressao()
	comparacao := "O resultado da da soma é 10"
	if impressao_ != comparacao {
		t.Error("String nao batem")
	}
}
