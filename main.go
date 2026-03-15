package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

// Interface

type ICreateMethods interface {
	Create(title string, description string) (*TodoStruct, error)
	Edit(id uuid.UUID, title string, description string) error
	Update(id uuid.UUID, status bool) error
	Delete(id uuid.UUID) error
}

type IReadMethods interface {
	Get(id uuid.UUID) (*TodoStruct, error)
	List() ([]TodoStruct, error)
}

type SaveMethods interface {
	ICreateMethods
	IReadMethods
}

// Struct

type TodoStruct struct {
	ID          uuid.UUID
	Title       string
	Description string
	Status      bool
}

// Slice

type TodoRepository struct {
	todos map[uuid.UUID]TodoStruct
}

// Function implementation

func (t *TodoRepository) Create(title string, description string) (*TodoStruct, error) {
	todo := TodoStruct{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		Status:      false,
	}
	t.todos[todo.ID] = todo
	return &todo, nil
}

func (t *TodoRepository) Edit(id uuid.UUID, title string, description string) error {
	//if _, ok := t.todos[id]; !ok {
	//	return fmt.Errorf("todo not found")
	//}
	//t.todos[id] = TodoStruct{
	//	ID:          id,
	//	Title:       title,
	//	Description: description,
	//	Status:      false,
	//}
	//return nil

	todo, ok := t.todos[id]
	if !ok {
		return fmt.Errorf("todo not found")
	}

	todo.Title = title
	todo.Description = description
	t.todos[id] = todo
	return nil
}

func (t *TodoRepository) Update(id uuid.UUID, status bool) error {
	//if _, ok := t.todos[id]; !ok {
	//	return fmt.Errorf("todo not found")
	//}
	//t.todos[id] = TodoStruct{
	//	ID:     id,
	//	Status: status,
	//}
	//return nil

	todo, ok := t.todos[id]
	if !ok {
		return fmt.Errorf("todo not found")
	}

	todo.Status = status
	t.todos[id] = todo
	return nil
}

func (t *TodoRepository) Delete(id uuid.UUID) error {
	if _, ok := t.todos[id]; !ok {
		return fmt.Errorf("todo not found")
	}
	delete(t.todos, id)
	return nil
}

func (t *TodoRepository) Get(id uuid.UUID) (*TodoStruct, error) {
	todo, ok := t.todos[id]
	if !ok {
		return nil, fmt.Errorf("todo not found")
	}

	return &todo, nil
}

func (t *TodoRepository) List() ([]TodoStruct, error) {
	var todos []TodoStruct
	for _, todo := range t.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

// Function
func readInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func readUUID(scanner *bufio.Scanner) (uuid.UUID, error) {
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())
	return uuid.Parse(text)
}

func MenuOptions() {
	fmt.Println("Escolha uma opção:")
	fmt.Println("1 - Criar")
	fmt.Println("2 - Listar")
	fmt.Println("3 - Atualizar")
	fmt.Println("4 - Finalizar")
	fmt.Println("5 - Deletar")
	fmt.Println("6 - Sair")
}

func Menu(m SaveMethods) {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		MenuOptions()
		input := readInput(scanner)
		option, _ := strconv.Atoi(input)

		switch option {
		case 1:
			fmt.Println("Digite o titulo:")
			title := readInput(scanner)
			fmt.Println("Digite a descricao:")
			description := readInput(scanner)

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
			id, err := readUUID(scanner)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Informações atuais")
			edit, _ := m.Get(id)
			fmt.Println(edit)
			fmt.Println("Altere o titulo:")
			title := readInput(scanner)
			fmt.Println("Altere a descrição:")
			description := readInput(scanner)
			m.Edit(id, title, description)

		case 4:
			fmt.Println("Digite o id da tarefa que será finalizada:")
			id, err := readUUID(scanner)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Informações atuais")
			update, _ := m.Get(id)
			fmt.Println(update)
			fmt.Println("Certeza que gostaria de finalizar a tarefa (y/n):")
			confirm := readInput(scanner)
			if confirm == "y" {
				m.Update(id, true)
				fmt.Println("Tarefa finalizada!")
			} else {
				fmt.Println("Tarefa ainda aberta!")
			}

		case 5:
			fmt.Println("Digite o id da tarefa que será deletada:")
			id, err := readUUID(scanner)
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

// Main

func main() {
	// Implementation of data structures

	repo := &TodoRepository{
		todos: make(map[uuid.UUID]TodoStruct),
	}

	methods := SaveMethods(repo)

	// Call functions

	Menu(methods)

}
