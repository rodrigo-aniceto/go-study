package main

// TODO
// implentar DEFINE

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

const ALT int = 12
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

func contarVizinhos(mapa [ALT][LAR]int, i, j int) int {
	//TODO otimizar para receber apenas as casas vizinhas, em vez de enviar a matriz inteira
	nunViz := 0
	if mapa[i-1][j-1] != 0 {
		nunViz += 1
	}
	if mapa[i-1][j] != 0 {
		nunViz += 1
	}
	if mapa[i-1][j+1] != 0 {
		nunViz += 1
	}
	if mapa[i][j-1] != 0 {
		nunViz += 1
	}
	if mapa[i][j+1] != 0 {
		nunViz += 1
	}
	if mapa[i+1][j-1] != 0 {
		nunViz += 1
	}
	if mapa[i+1][j] != 0 {
		nunViz += 1
	}
	if mapa[i+1][j+1] != 0 {
		nunViz += 1
	}
	return nunViz
}

func atualizaMapa(mapa *[ALT][LAR]int) {
	var mapaAux [ALT][LAR]int = *mapa

	for i := 1; i < ALT-1; i++ {
		for j := 1; j < LAR-1; j++ {
			viz := contarVizinhos(mapaAux, i, j)
			if mapaAux[i][j] != 0 {
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

	mapa[1][1] = 2

	mapa[4][12] = 3
	mapa[4][13] = 3
	mapa[4][14] = 3
	mapa[3][14] = 3
	mapa[2][13] = 3

	mapa[5][22] = 2
	mapa[6][22] = 2
	mapa[6][23] = 2

	mapa[5][42] = 2
	mapa[6][42] = 2
	mapa[7][42] = 2
	mapa[5][43] = 2
	mapa[6][43] = 2
	mapa[7][43] = 2
	mapa[5][44] = 2
	mapa[6][44] = 2
	mapa[7][44] = 2

	for i := 0; i < 50; i++ {

		//fmt.Println("\033[2J") //limpar a tela, pesquise uma solução melhor

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		fmt.Println("CONWAY's Game of Life")
		exibeMapa(mapa)

		atualizaMapa(&mapa)

		time.Sleep(time.Second / 2)
	}
}
