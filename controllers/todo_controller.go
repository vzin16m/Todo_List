package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"todo/models"
	"todo/views"
)

func Menu(m models.SaveMethods) {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		views.MenuOptions()
		input := views.ReadInput(scanner)
		option, _ := strconv.Atoi(input)

		switch option {
		case 1:
			fmt.Println("Digite o titulo:")
			title := views.ReadInput(scanner)
			fmt.Println("Digite a descricao:")
			description := views.ReadInput(scanner)

			if _, err := m.Create(title, description); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Tarefa criada!")
			}
		case 2:
			list, err := m.List()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(list)
			}

		case 3:
			fmt.Println("Digite o id da tarefa que será alterada:")
			id, err := views.ReadUUID(scanner)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Informações atuais")
			edit, _ := m.Get(id)
			fmt.Println(edit)
			fmt.Println("Altere o titulo:")
			title := views.ReadInput(scanner)
			fmt.Println("Altere a descrição:")
			description := views.ReadInput(scanner)
			m.Edit(id, title, description)

		case 4:
			fmt.Println("Digite o id da tarefa que será finalizada:")
			id, err := views.ReadUUID(scanner)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Informações atuais")
			update, _ := m.Get(id)
			fmt.Println(update)
			fmt.Println("Certeza que gostaria de finalizar a tarefa (y/n):")
			confirm := views.ReadInput(scanner)
			if confirm == "y" {
				m.Update(id, true)
				fmt.Println("Tarefa finalizada!")
			} else {
				fmt.Println("Tarefa ainda aberta!")
			}

		case 5:
			fmt.Println("Digite o id da tarefa que será deletada:")
			id, err := views.ReadUUID(scanner)
			if err != nil {
				fmt.Println(err)
				return
			}
			m.Delete(id)

		case 6:
			os.Exit(0)

		default:
			fmt.Println("Opção inválida")
		}
	}
}
