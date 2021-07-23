package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {

	var l1 string
	fmt.Println("Введите данные из искомой строчки")
	fmt.Scanf("%s\n", &l1)

	var Username string
	fmt.Println("Введите Username")
	fmt.Scanf("%s\n", &Username)

	var IP string
	fmt.Println("Введите IP")
	fmt.Scanf("%s\n", &IP)

	input, err := ioutil.ReadFile("user-aliases")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	var texts = Username + " " + IP
	for i, line := range lines {
		if strings.Contains(line, l1) {
			lines[i] = texts
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile("user-aliases", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
