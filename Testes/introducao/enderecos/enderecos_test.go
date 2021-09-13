// package enderecos -> podemos usar assim ou como está abaixo
package enderecos_test

import "testing"

type cenariosDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

//Precisa começar com Test
func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	cenariosDeTeste := []cenariosDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia ash", "Rodovia"},
		{"Praça ahuas", "Tipo Inválido"},
		{"Estrada qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"", "Tipo Inválido"},
	}
	// enderecoParaTeste := "Avenida Paulista"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s", cenario.retornoEsperado, tipoDeEnderecoRecebido)
		}
	}
}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}
