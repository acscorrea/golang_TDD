package main

import "testing"

func TestOla(t *testing.T) {

	checkMessage := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado %s, esperado %s", resultado, esperado)
		}
	}

	t.Run("espanhol", func(t *testing.T) {
		resultado := Ola("Augusto", "es")
		esperado := "Hola, Augusto"
		checkMessage(t, resultado, esperado)
	})

	t.Run("frances", func(t *testing.T) {
		resultado := Ola("Augusto", "fr")
		esperado := "Bonjour, Augusto"
		checkMessage(t, resultado, esperado)
	})

	t.Run("diz olá para as pessoas", func(t *testing.T) {

		resultado := Ola("Chris", "")
		esperado := "Olá, Chris"
		checkMessage(t, resultado, esperado)
	})

	t.Run("diz olá mundo quando recebe string vazia", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"
		checkMessage(t, resultado, esperado)
	})
}
