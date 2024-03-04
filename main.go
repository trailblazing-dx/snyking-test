package main 

import (
	"os"
)


func main() {
	password := "123345"
	ret, err := os.ReadFile(password)
	if err != nil {
		return 
	}
}

