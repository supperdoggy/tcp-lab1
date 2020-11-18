package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type obj map[string]interface{}
type arr []interface{}
func main(){
	fmt.Println("opening aneks database")
	content, err := ioutil.ReadFile("/home/user/go/src/projects/tcp/tcp-lab1/src/files/500.json")
	if err != nil{
		panic(err.Error())
	}
	aneksMap := obj{}
	err = json.Unmarshal(content, &aneksMap)
	if err != nil{
		panic(err.Error())
	}
	var aneksArr arr = aneksMap["anek"].([]interface{})
	fmt.Println(aneksArr[0].(map[string]interface{})["text"])

}
