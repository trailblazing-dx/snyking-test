package main 

import (
	"os"
)


func main() {
	password := "123345678"
	ret, err := os.ReadFile(password)
	if err != nil {
		return 
	}
}

