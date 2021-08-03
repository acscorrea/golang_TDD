package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Augusto")
	esperado := "Olá, Augusto"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
