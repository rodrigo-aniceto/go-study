package main

// TODO
// implentar DEFINE

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const ALT int = 15
const LAR int = 60

func preencheMapa(mapa *[ALT][LAR]int) {

	for i := 1; i < ALT-1; i++ {
		for j := 1; j < LAR-1; j++ {
			if rand.Intn(100) > 70 {
				mapa[i][j] = 1
			}
		}
	}

}

func exibeMapa(mapa [ALT][LAR]int) {
	//fmt.Println(mapa)

	fmt.Print("+")
	for i := 0; i < LAR-2; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	for i := 1; i < ALT-1; i++ {
		fmt.Print("|")
		for j := 1; j < LAR-1; j++ {
			if mapa[i][j] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Print("|\n")
	}

	fmt.Print("+")
	for i := 0; i < LAR-2; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")

}

func atualizaMapa(mapa *[ALT][LAR]int) {
	var mapaAux [ALT][LAR]int = *mapa

	for i := 1; i < ALT-1; i++ {
		for j := 1; j < LAR-1; j++ {

			//contar os vizinhos
			viz := mapaAux[i-1][j-1] + mapaAux[i-1][j] + mapaAux[i-1][j+1] + mapaAux[i][j-1] +
				mapaAux[i][j+1] + mapaAux[i+1][j-1] + mapaAux[i+1][j] + mapaAux[i+1][j+1]

			if mapaAux[i][j] == 1 {
				if viz < 2 {
					mapa[i][j] = 0
				} else if viz == 2 || viz == 3 {
					mapa[i][j] = 1
				} else {
					mapa[i][j] = 0
				}

			} else if mapaAux[i][j] == 0 {
				if viz == 3 {
					mapa[i][j] = 1
				} else {
					mapa[i][j] = 0
				}
			}
		}
	}

}

func main() {

	var mapa [ALT][LAR]int

	preencheMapa(&mapa)

	for i := 0; i < 100; i++ {

		//fmt.Println("\033[2J") //limpar a tela, pesquise uma solução melhor

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		fmt.Println("CONWAY's Game of Life")
		exibeMapa(mapa)

		atualizaMapa(&mapa)

		time.Sleep(time.Second / 4)
	}
}
