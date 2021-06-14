package tiposdados3

import "fmt"

func Funcao() {
	//tipos de inteiros: int8, int16, int32, int64, se omitir ou escrever só int vai criar baseado na arq do seu PC
	//uint8, uint16, uint32, uint64 eh int sem sinal
	//outros nomes: rune = int32 e byte = uint8
	var num int = 100
	var num2 uint32 = 23290
	var num3 byte = 128

	fmt.Println(num, num2, num3)

	//exitem dois tipos de float
	var numReal float32 = 122.15
	var numReal2 float64 = 1622.15
	//var numReal3 float = 5622.15 //isso daria erro, não pode chamar de float

	fmt.Println(numReal, numReal2)

	texto := 'b' //variaveis de character são convertidas pra inteiro
	fmt.Println(texto)

	//booleano
	var bolean1 bool = true
	fmt.Println(bolean1)

	//tipo erro
	var erro error = nil
	fmt.Println(erro)

	//se não inicializar um numero vale 0, uma string eh vazia, o booelanico vale false e o erro vale nil
}
