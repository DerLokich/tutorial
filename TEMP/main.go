// pass_fail сообщает, сдал ли пользователь экзамен.
package main

import (
	"fmt"
	"github.com/Velzevol/keyboard/keyboard"
	"log"
)

func main() {
	var status string

	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()

	if err != nil {
		log.Fatal(err)
	}

	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println(status)
}
