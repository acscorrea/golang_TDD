package main

import "fmt"

const espanhol = "es"
const frances = "fr"

const olaPT = "Olá, "
const olaES = "Hola, "
const olaFR = "Bonjour, "

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = olaFR
	case espanhol:
		prefixo = olaES
	default:
		prefixo = olaPT
	}
	return
}

func main() {
	fmt.Println(Ola("mundo", ""))
}
