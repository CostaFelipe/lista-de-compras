package v1

import (
	"errors"
	"fmt"
	"strings"
)

const message = "ok"

type Item struct {
	Nome       string
	Quantidade int
}

type ListaDeCompras struct {
	Itens []Item
}

func (l *ListaDeCompras) AdicionarItem(nome string, quantidade int) (string, error) {
	if nome == "" || quantidade <= 0 {
		return "", errors.New("error")
	}

	item := Item{Nome: nome, Quantidade: quantidade}
	l.Itens = append(l.Itens, item)

	return message, nil
}

func (l *ListaDeCompras) RemoverItem(nome string) {
	for i, item := range l.Itens {
		if strings.ToLower(item.Nome) == strings.ToLower(nome) {
			l.Itens = append(l.Itens[:i], l.Itens[i+1:]...)
			break
		}
	}
}

func (l *ListaDeCompras) Display() {
	for _, item := range l.Itens {
		fmt.Printf("Nome:%s Quantidade:%d\n", item.Nome, item.Quantidade)
	}

}
