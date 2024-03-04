package main 

import (
	"os"
)


func main() {
	password := "1233456"
	ret, err := os.ReadFile(password)
	if err != nil {
		return 
	}
}

