package estcontrole11

import "fmt"

func Funcao() {
	fmt.Println("Estruturas de Controle")

	numero := 15

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else if numero == 15 {
		fmt.Println("Igual a 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	// inicializar dentro do if

	if outroNumero := numero; outroNumero > 0 { //criou e verificou dentro do if
		fmt.Println("Numero eh maior que zero") //essa variavel só existe dentro do if e else que vir depois
	}

	//switch case

	dia := diaDaSemana(4)
	fmt.Println(dia)

	dia = diaDaSemana2(5)
	fmt.Println(dia)

}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default: //precisa dele aqui pra não ocorrer situação sem retorno
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string { //eh possivel colocar as ondições dentro de cada case

	var diaDaSemana string
	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		// fallthrough // faz ele entrar no próximo condição, sem nem avaliar
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}

	return diaDaSemana

}
