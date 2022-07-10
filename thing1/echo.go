package thing1

import (
	"embed"
	"fmt"
)

//go:embed message.txt
var f embed.FS

func Speak() (string, error) {
	data, err := f.ReadFile("message.txt")
	if err != nil {
		fmt.Printf("could not read message: %v", err)
		return "", nil
	} else {
		fmt.Printf("%s\n", string(data))
		return string(data), nil
	}
}
