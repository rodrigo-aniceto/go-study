package enderecos

import "strings"

//TipoDeEndereco função para verificar com um endereço começa dentro de um padrão
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	endereco = strings.ToLower(endereco)                        //mudar para letra minuscula
	primeraPalavraDoEndereco := strings.Split(endereco, " ")[0] //apenas o primeiro indice no slice de palavras

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
			break
		}
	}

	if enderecoTemUmTipoValido {
		return primeraPalavraDoEndereco
	}

	return "Tipo Inválido"
}
