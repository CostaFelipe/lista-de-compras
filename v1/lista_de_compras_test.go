package v1

import "testing"

func TestAdicionarItens(t *testing.T) {
	lista := ListaDeCompras{}

	lista.AdicionarItem("Banana", 5)
	resultado := lista.Itens[0].Nome
	esperado := "Banana"

	if len(lista.Itens) != 1 {
		t.Errorf("Esperado 1 item, mas obteve '%d'", len(lista.Itens))
	}

	if resultado != esperado {
		t.Errorf("resultado: '%s', esperado: %s", resultado, esperado)
	}

	if lista.Itens[0].Quantidade != 5 {
		t.Errorf("resultado: '%d', esperado 5", lista.Itens[0].Quantidade)
	}
}

func TestRemoverItem(t *testing.T) {
	lista := ListaDeCompras{}

	lista.AdicionarItem("Laranja", 3)
	lista.AdicionarItem("Arroz blue ville 1kg", 2)

	lista.RemoverItem("Laranja")

	if len(lista.Itens) != 1 {
		t.Errorf("Esperado 1 item, mas obteve '%d'", len(lista.Itens))
	}

	if lista.Itens[0].Nome != "Arroz blue ville 1kg" {
		t.Errorf("esperado: 'Arroz blue ville 1kg' resultado: '%v'", lista.Itens[0].Nome)
	}
}
