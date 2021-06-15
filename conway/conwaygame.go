package main

// TODO
// implentar DEFINE

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const ALT int = 15
const LAR int = 60

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

	mapa[1][3] = 1
	mapa[2][3] = 1
	mapa[3][3] = 1
	mapa[4][3] = 1

	mapa[4][12] = 1
	mapa[4][13] = 1
	mapa[4][14] = 1
	mapa[3][14] = 1
	mapa[2][13] = 1

	mapa[9][16] = 1
	mapa[9][15] = 1
	mapa[9][14] = 1
	mapa[8][14] = 1
	mapa[7][15] = 1

	mapa[5][22] = 1
	mapa[6][22] = 1
	mapa[6][23] = 1

	mapa[9][22] = 1
	mapa[10][22] = 1
	mapa[10][23] = 1

	mapa[4][32] = 1
	mapa[4][33] = 1
	mapa[4][34] = 1
	mapa[3][34] = 1
	mapa[2][33] = 1

	mapa[5][42] = 1
	mapa[6][42] = 1
	mapa[7][42] = 1
	mapa[5][43] = 1
	mapa[6][43] = 1
	mapa[7][43] = 1
	mapa[5][44] = 1
	mapa[6][44] = 1
	mapa[7][44] = 1

	mapa[6][50] = 1
	mapa[6][50] = 1
	mapa[7][51] = 1
	mapa[6][51] = 1
	mapa[6][52] = 1
	mapa[7][52] = 1
	mapa[6][53] = 1
	mapa[7][53] = 1
	mapa[5][54] = 1
	mapa[6][54] = 1
	mapa[8][54] = 1

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
