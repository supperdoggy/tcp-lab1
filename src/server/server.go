package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"time"
)


var (
	jokes = getJokes("../files/500.json")
)

func sendAnswer(answer string, conn net.Conn) error{
	_, err := conn.Write([]byte(answer + "\n"))
	if err != nil{
		return err
	}
	return nil
}

// returns random joke
func pickRandomJoke() string{
	rand.Seed(time.Now().UnixNano())
	return jokes.A[rand.Intn(500)].Text
}

func getJokes(path string) (jokes Aneks){
	content, err := ioutil.ReadFile(path)
	if err != nil{
		fmt.Println("cant read file with jokes")
		panic(err.Error())
	}
	err = json.Unmarshal(content, &jokes)
	if err != nil{
		fmt.Println("cant unmarshal json")
		panic(err.Error())
	}
	return
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8181")
	defer ln.Close()
	conn, _ := ln.Accept()
	// listening
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		if string(message) == "get joke\n"{
			answer := pickRandomJoke()

			if err := sendAnswer(answer, conn);err!=nil{
				fmt.Println(err.Error())
			}
		}else{
			answer := "error"
			if err := sendAnswer(answer, conn);err!=nil{
				fmt.Println(err.Error())
			}
		}

	}
}
