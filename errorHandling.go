package main

import "fmt"

type NotFoundError struct {
	Message string
}

func (catchedError *NotFoundError) Error() string {
	return fmt.Sprintf("Error :  %v", catchedError.Message)
}

func throwError() error {
	return &NotFoundError{"Not found error was thrown"}
}

func main()  {
	currentError := throwError()

	if currentError != nil {
		fmt.Println(currentError)
	}
}
