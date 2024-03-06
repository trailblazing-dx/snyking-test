package main 

import (
	"os"
)


func main() {
	password := "123456" 
	ret, err := os.ReadFile(password)
	if err != nil {
		return 
	}
}

