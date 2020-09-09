package main

import (
"bufio"
"fmt"
"motd/message"
"strings"
"os"
"io/ioutil"
)


func main() {
	f, err := os.OpenFile("motd.txt", os.O_WRONLY, 0644)

	if err != nil {
	fmt.Println("error occured : unable to open file")
	os.Exit(1)
}
	defer f.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("your Greeting: " )
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

        fmt.Print("your Name: " )
        name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message:= message.Greeting(name, phrase)
	//fmt.Println(message)
	_ , err = f.Write([]byte(message)) 
	err = ioutil.WriteFile("motd.txt", []byte(message), 0644)


	if err != nil {
	fmt.Println("Error Occured : Unable to write /etc/motd")
	os.Exit(1)
	}


}

