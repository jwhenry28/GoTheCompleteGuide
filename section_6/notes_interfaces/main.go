package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes_interfaces/note"
	"example.com/notes_interfaces/todo"
)

type Saver interface {
	Save() error
}

type Outputtable interface {
	Saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {	
		return
	}

	err = outputData(userNote)
	if err != nil {	
		return
	}

	printSomething(1)
	printSomething(1.5)
	printSomething("hello")
	printSomething(todo)
}

func printSomething(value interface{}) {
	intVal, ok := value.(int)
	if ok {
		fmt.Println("Integer:", intVal)
	}

	floatVal, ok := value.(float64)
	if ok {
		fmt.Println("Float:", floatVal)
	}

	stringVal, ok := value.(string)
	if ok {
		fmt.Println("String:", stringVal)
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer:", value)
	// case float64:
	// 	fmt.Println("Float:", value)
	// case string:
	// 	fmt.Println("String:", value)
	// }
}

func outputData(data Outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		return err
	}

	fmt.Println("Saving the note succeeded.")
	return nil
}

func getNoteData() (string, string){
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}