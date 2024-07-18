package todolistapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


type Todo struct {
	ID int
	TodoTitle string
	TodoDescription string
	TodoStartDate string
}

var todos []Todo
var count uint = 0
var todoTitle string
var todoDescription string
var todoStartDate string
var todoID int



func Todolistapp() {
	for {
		var userInput int
		
		commands()
		
		fmt.Println("Enter your choice")
		fmt.Scan(&userInput)

		if userInput == 1{
			create_todo()	
		}else if userInput == 2{
			readTodos()
		}else if userInput == 3{
			updateTodos()
		}else if userInput == 4{
			deleteTodos()
		}
		
		fmt.Printf("Todos %v", todos)
		count += 1
	}
	
}

func commands() {
	fmt.Print("\n\n\n")
	fmt.Println("Press 1 to create todo")
	fmt.Println("Press 2 to read todo")
	fmt.Println("Press 3 to update todo")
	fmt.Println("Press 4 to delete todo")

	fmt.Print("\n\n\n")
	
}

func create_todo() {
	


	fmt.Println("Enter the todo title")
	fmt.Scan(&todoTitle)

	fmt.Println("Enter todo description")
	fmt.Scan(&todoDescription)

	fmt.Println("Enter todo start date")
	fmt.Scan(&todoStartDate)
	
	var todo  = Todo {
		TodoTitle: todoTitle,
		TodoDescription: todoDescription,
		TodoStartDate: todoStartDate,
		ID: int(count),
	}
	todos = append(todos, todo)
}

// func updateTodos() {	
	
// 	fmt.Println("Enter Todo ID to update: ")
// 	fmt.Scan(&todoID)
// 	fmt.Println("Leave blank where you are not updating")
// 	for index, todo := range(todos) {
// 		if todo.ID == todoID{
// 			fmt.Println("Enter todo title")	
// 			fmt.Scanln(&todoTitle)
// 			if todoTitle != "" {
// 				todos[index].TodoTitle = todoTitle				
// 			}
// 			fmt.Println("Enter todo description")
// 			fmt.Scanln(&todoDescription)
// 			if todoDescription != "" {
// 				todos[index].TodoDescription = todoDescription				
// 			}
// 			fmt.Println("Enter todo start date")
// 			fmt.Scanln(&todoStartDate)

// 			if todoStartDate != "" {
// 				todos[index].TodoStartDate = todoStartDate
// 			}
// 			fmt.Printf("Todo with ID %d updated successfully!\n", todoID)
// 		}
// 	}	
// }
func updateTodos() {
	var todoID int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Todo ID to update: ")
	fmt.Scan(&todoID)
	fmt.Println("Leave blank where you are not updating")

	for index, todo := range todos {
		if todo.ID == todoID {
			fmt.Println("Enter todo title")
			todoTitle, _ := reader.ReadString('\n')
			todoTitle = strings.TrimSpace(todoTitle)
			if todoTitle != "" {
				todos[index].TodoTitle = todoTitle
			}

			fmt.Println("Enter todo description")
			todoDescription, _ := reader.ReadString('\n')
			todoDescription = strings.TrimSpace(todoDescription)
			if todoDescription != "" {
				todos[index].TodoDescription = todoDescription
			}

			fmt.Println("Enter todo start date")
			todoStartDate, _ := reader.ReadString('\n')
			todoStartDate = strings.TrimSpace(todoStartDate)
			if todoStartDate != "" {
				todos[index].TodoStartDate = todoStartDate
			}

			fmt.Printf("Todo with ID %d updated successfully!\n", todoID)
			return
		}
	}
	fmt.Printf("Todo with ID %d not found.\n", todoID)
}

func readTodos() {
	if len(todos) == 0{
		fmt.Println("Todos not found")
	}else{
		fmt.Println("Todos")
		for _, todo := range(todos) {		

			fmt.Printf("ID: %d, Title: %s, Description: %s, Start Date: %s\n",
			todo.ID, todo.TodoTitle, todo.TodoDescription, todo.TodoStartDate)
			fmt.Println("################################")
		}
	}
}

func deleteTodos() {
	var todoIndex int
	fmt.Println("Enter Todo ID to delete:")
	fmt.Scan(&todoID)
	
	for index, todo := range(todos) {
		if todo.ID == int(todoID){
			todoIndex = index
		}
	}

	todos = append(todos[:todoIndex], todos[todoIndex + 1:]...)
}