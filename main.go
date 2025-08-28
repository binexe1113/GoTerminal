package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Mostra prompt
		fmt.Print("go-shell> ")

		// Lê entrada do usuário
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Se digitar "exit" sai do terminal
		if input == "exit" {
			fmt.Println("Saindo...")
			break
		}

		// Divide comando e argumentos
		args := strings.Split(input, " ")
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// Executa o comando
		err := cmd.Run()
		if err != nil {
			fmt.Println("Erro:", err)
		}
	}
}
