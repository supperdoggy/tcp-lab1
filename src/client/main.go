package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type obj map[string]interface{}

var(
	path = "/Users/maks/go/src/github.com/supperdoggy/tcp/src/files"
)

func decodeIBM866(m obj, data []byte) []string{
	result := []string{}
	for _, v := range data{
		if d, ok := m[formatTo16Bit(v)];ok{
			result = append(result, d.(string))
		}
	}
	return result
}

func sliceStringIntoString(s []string) string{
	result := ""
	for _, v := range s{
		fmt.Println(v)
		result += v
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
	result := obj{}
	if err := json.Unmarshal(content, &result);err!=nil{
		panic(err.Error())
	}

	content, err = ioutil.ReadFile(path+"/1.txt")
	if err != nil{
		panic(err.Error())
	}
	decoded := decodeIBM866(result, content)
	fmt.Println(decoded)

}
