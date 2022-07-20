package main

import (
	linkedlist "data-structures/linked_list"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type InputReader struct {
	reader *bufio.Reader
}

func (i *InputReader) getLine(label string) string {
	fmt.Print(label)
	text, _ := i.reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func parseNumber(input string) int {
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(input, " is not an integer.")
		os.Exit(1)
	}
	return number
}

func main() {
	linkedList := linkedlist.LinkedList{}
	reader := InputReader{reader: bufio.NewReader(os.Stdin)}
	input := reader.getLine(">>>> ")
	inputs := strings.Split(strings.TrimSpace(input), " ")
	for inputs[0] != "exit" {
		execute(&linkedList, inputs)
		input = reader.getLine(">>>> ")
		inputs = strings.Split(strings.TrimSpace(input), " ")
	}
	os.Exit(0)

}

func execute(linkedList *linkedlist.LinkedList, inputs []string) {
	switch inputs[0] {
	case "insert_first":
		if len(inputs) != 2 {
			fmt.Println("Invalid input")
			return
		}
		linkedList.InsertAt(0, parseNumber(inputs[1]))
	case "insert":
		if len(inputs) < 2 {
			fmt.Println("Invalid input")
			return
		}
		for _, number := range inputs[1:] {
			linkedList.InsertAt(linkedList.Length-1, parseNumber(number))
		}
	case "insert_at":
		if len(inputs) != 3 {
			fmt.Println("Invalid input")
			return
		}
		linkedList.InsertAt(parseNumber(inputs[2]), parseNumber(inputs[1]))
	case "insert_last":
		if len(inputs) != 2 {
			fmt.Println("Invalid input")
			return
		}
		linkedList.InsertAt(linkedList.Length, parseNumber(inputs[1]))
	case "delete_first":
		if len(inputs) != 1 {
			fmt.Println("Invalid input")
			return
		}
		linkedList.DeleteAt(0)
	case "delete":
		if len(inputs) != 1 {
			fmt.Println("Invalid input")
			return
		}
		linkedList.DeleteAt(linkedList.Length)
	case "delete_at":
		if len(inputs) != 2 {
			fmt.Println("Invalid input")
			return
		}
		linkedList.DeleteAt(parseNumber(inputs[1]))
	case "delete_last":
		if len(inputs) != 1 {
			fmt.Println("Invalid input")
			return
		}
		linkedList.DeleteAt(linkedList.Length - 1)
	case "print":
		linkedList.PrintAll()
	default:
		fmt.Print("invalid command")
	}
	linkedList.PrintAll()
}

