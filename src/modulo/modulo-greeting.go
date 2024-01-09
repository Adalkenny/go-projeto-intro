package modulo

import (
	"errors"
	"fmt"
)

func Comprimentar(nome string) (string, error) {

	if nome == "" {
		return "", errors.New("Nome vazia")
	}
	mensagem := fmt.Sprintf("Hello, %v. benvindo!", nome)
	return mensagem, nil
}
