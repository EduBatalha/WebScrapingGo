// Dentro de views/console_view.go

package views

import (
	"fmt"
)

// ConsoleView define métodos para exibir mensagens no console.
type ConsoleView struct{}

// DisplayError exibe um erro no console.
func (cv ConsoleView) DisplayError(err error) {
	fmt.Println("Erro:", err)
}

// ShowMessage exibe uma mensagem no console.
func (cv ConsoleView) ShowMessage(message string) {
	fmt.Println(message)
}
