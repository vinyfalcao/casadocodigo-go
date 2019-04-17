package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso:	conversor	<valores>	<unidade>")
		os.Exit(1)
	}

	unidadeOrigem := os.Args[len(os.Args)-1]
	valoresOrigem := os.Args[1 : len(os.Args)-1]
	converterValor(valoresOrigem, unidadeOrigem)

}

func getUnidadeDestino(unidadeOrigem string) string {
	var ret string
	if unidadeOrigem == "celsius" {
		ret := "fahrenheit"
		return ret
	} else if unidadeOrigem == "quilometros" {
		ret := "milhas"
		return ret
	} else {
		fmt.Printf("%s	não	é	uma	unidade	conhecida!",
			unidadeOrigem)
	}
	return ret
}

func converterValor(valoresOrigem []string, unidadeOrigem string) {
	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf("O valor %s ma posição %d não é um número válido!", v, i)
			os.Exit(1)
		}
		var valorDestino float64
		if unidadeOrigem == "celsius" {
			valorDestino = valorOrigem*1.8 + 32
		} else {
			valorDestino = valorOrigem / 1.60934
		}
		fmt.Printf("%.2f	%s	=	%.2f	%s\n",
			valorOrigem, unidadeOrigem,
			valorDestino, getUnidadeDestino(unidadeOrigem))
	}
}
