package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const ALT int = 15
const LAR int = 60

type comida struct {
	posY int
	posX int
}

type rato struct {
	posY   int
	posX   int
	destY  int
	destX  int
	passos int
}

func (r *rato) inicializaRato() {
	r.posX = 5
	r.posY = 5
	r.destX = 5
	r.destY = 5
	r.passos = 0
}

func medirDistancia(ax, bx, ay, by int) int {
	n := ax - bx
	m := ay - by
	if n < 0 {
		n = n * -1
	}
	if m < 0 {
		m = m * -1
	}
	return n + m
}

func atualizaMapa(mapa *[ALT][LAR]int, listaComida *[]comida, mickey *rato) {

	// se já encontrou comida, procurar a mais proxima
	if mickey.posX == mickey.destX && mickey.posY == mickey.destY {
		for i := 0; i < len(*listaComida); i++ {
			// remover da lista
			if (*listaComida)[i].posX == mickey.posX && (*listaComida)[i].posY == mickey.posY {
				*listaComida = append((*listaComida)[:i], (*listaComida)[i+1:]...)
				break
			}
		}
		// encontrar mais proxima
		if len(*listaComida) == 0 {
			//encerra o fluxo se lista vazia
			return
		}
		distancia := medirDistancia((*listaComida)[0].posX, mickey.posX, (*listaComida)[0].posY, mickey.posY)
		indiceProximo := 0
		for i := 0; i < len(*listaComida); i++ {
			distanciaAux := medirDistancia((*listaComida)[i].posX, mickey.posX, (*listaComida)[i].posY, mickey.posY)
			if distanciaAux < distancia {
				distancia = distanciaAux
				indiceProximo = i
			}
		}
		mickey.destY = (*listaComida)[indiceProximo].posY
		mickey.destX = (*listaComida)[indiceProximo].posX

		// se ainda está navegando
	} else {
		if mickey.posX < mickey.destX { // primeiro anda no eixo X depois Y
			mapa[mickey.posY][mickey.posX] = ' '
			mickey.posX += 1
			mapa[mickey.posY][mickey.posX] = 'C'
		} else if mickey.posX > mickey.destX {
			mapa[mickey.posY][mickey.posX] = ' '
			mickey.posX -= 1
			mapa[mickey.posY][mickey.posX] = 'C'
		} else if mickey.posY < mickey.destY {
			mapa[mickey.posY][mickey.posX] = ' '
			mickey.posY += 1
			mapa[mickey.posY][mickey.posX] = 'C'
		} else if mickey.posY > mickey.destY {
			mapa[mickey.posY][mickey.posX] = ' '
			mickey.posY -= 1
			mapa[mickey.posY][mickey.posX] = 'C'
		}
		mickey.passos += 1
	}
}

func preencheMapa(mapa *[ALT][LAR]int, listaComida *[]comida) {

	var c comida

	mapa[0][0] = '+'
	mapa[ALT-1][0] = '+'
	mapa[0][LAR-1] = '+'
	mapa[ALT-1][LAR-1] = '+'

	for i := 1; i < ALT-1; i++ {
		mapa[i][0] = '|'
		mapa[i][LAR-1] = '|'
	}

	for j := 1; j < LAR-1; j++ {
		mapa[0][j] = '-'
		mapa[ALT-1][j] = '-'
	}

	for i := 1; i < ALT-1; i++ {
		for j := 1; j < LAR-1; j++ {
			if rand.Intn(100) > 97 {
				mapa[i][j] = '*'

				c.posY = i
				c.posX = j
				*listaComida = append(*listaComida, c)
			} else {
				mapa[i][j] = ' '
			}
		}
	}

}

func exibeMapa(mapa [ALT][LAR]int) {
	//fmt.Println(mapa)

	for i := 0; i < ALT; i++ {
		for j := 0; j < LAR; j++ {
			fmt.Printf("%c", mapa[i][j])
		}
		fmt.Print("\n")
	}

}

func main() {

	var mapa [ALT][LAR]int
	var listaComida []comida
	var mickey rato

	mickey.inicializaRato()

	preencheMapa(&mapa, &listaComida)
	mapa[mickey.posY][mickey.posX] = 'C'
	exibeMapa(mapa)

	for len(listaComida) > 0 {

		//limpar tela
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()

		atualizaMapa(&mapa, &listaComida, &mickey)
		exibeMapa(mapa)
		time.Sleep(time.Second / 4)
	}
	fmt.Println("Numero de passos: ", mickey.passos)
}
