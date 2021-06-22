package metodos14

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

//simulacao de orientação a objetos
//metodo eh uma função associada a uma struct
func (u usuario) salvar() {
	fmt.Printf("Metodo de salvar o usuário %s no BD\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

// se for alterar um campo tem que ser um ponteiro
func (u *usuario) fazerAniversario() {
	u.idade++
}

func Funcao1() {

	user1 := usuario{"João", 22}
	user1.salvar()

	user2 := usuario{"Davi", 17}
	user2.salvar()
	fmt.Println(user2.maiorDeIdade())
	user2.fazerAniversario()
	fmt.Println(user2.maiorDeIdade())
}
