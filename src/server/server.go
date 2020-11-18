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

// for shortcuts
type obj map[string]interface{}
type arr []interface{}

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
	return jokes["anek"].([]interface{})[rand.Intn(len(jokes["anek"].([]interface{})))].(map[string]interface{})["text"].(string)
}

func getJokes(path string) obj{
	content, err := ioutil.ReadFile(path)
	if err != nil{
		fmt.Println("cant read file with jokes")
		panic(err.Error())
	}
	var jokes obj
	err = json.Unmarshal(content, &jokes)
	if err != nil{
		fmt.Println("cant unmarshal json")
		panic(err.Error())
	}
	return jokes
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
