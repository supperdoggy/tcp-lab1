package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type obj map[string]interface{}

var(
	path = "/home/user/go/src/projects/tcp/tcp-lab1/src/files"
)

func decodeIBM866(m obj, data []byte) []string{
	result := []string{}
	i := 1
	for _, v := range data{
		fmt.Println(len(data)-i)
		if d, ok := m[formatTo16Bit(v)];ok{
			result = append(result, d.(string))
		}
		i++
	}
	return result
}

func sliceStringIntoString(s []string) string{
	result := ""
	i:=1
	for _, v := range s{
		fmt.Println(len(s)-i)
		result += v
		i++
	}
	return result
}

func formatTo16Bit(b byte) string{
	return strconv.FormatInt(int64(b), 16)
}

func formatToInt16(b byte) string{
	return strconv.FormatInt(int64(b), 16)
}
func main(){
	content, err := ioutil.ReadFile(path+"/parsedData.json")
	if err != nil{
		panic(err.Error())
	}
	decodeMap := obj{}
	if err := json.Unmarshal(content, &decodeMap);err!=nil{
		panic(err.Error())
	}
	var decoded []string
	for i:=1;i!=9;i++{
		fmt.Println("Making", i, "file...")
		content, err = ioutil.ReadFile(path+"/"+strconv.Itoa(i)+".txt")
		if err != nil{
			panic(err.Error())
		}
		decoded = append(decoded, decodeIBM866(decodeMap, content)...)
	}
	fmt.Println(len(decoded))
	anecdotesString := sliceStringIntoString(decoded)
	fmt.Println("Opening file...")
	f, err := os.OpenFile("/home/user/go/src/projects/tcp/tcp-lab1/src/files/decoded.txt", os.O_WRONLY, 0755)
	if err != nil{
		panic(err.Error())
	}
	defer f.Close()
	fmt.Println("Writing to file....")
	if _, err = f.Write([]byte(anecdotesString));err!=nil{
		panic(err.Error)
	}
	fmt.Println("Done!")

}
