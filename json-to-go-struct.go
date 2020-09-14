package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/golang/glog"
)

func init() {
	flag.Set("stderrthreshold", "0")
	flag.Set("v", "2")
	flag.Parse()
}

func main() {
	dat, err := ioutil.ReadFile("input.json")
	if err != nil {
		glog.Error("Error Reading FIle:", err)
		os.Exit(1)
	}
	var JSONInput interface{}
	err = json.Unmarshal(dat, &JSONInput)
	if err != nil {
		glog.Error("Error JSON Unmarshal Failed:", err)
		os.Exit(1)
	}

	KeysPresent = make(map[int][]string)
	FinalStruct = "type AutoGenStruct struct {\n"
	IterateOverMap(JSONInput)
	FinalStruct += "}"
	glog.Info("OutputStruct\n\n", FinalStruct)

}

//FinalStruct Final Structure
var FinalStruct string
var tabs string

//KeysPresent inside map
var KeysPresent map[int][]string
var level int

//IterateOverMap Function to Iterate over a map and find out it's DS
func IterateOverMap(input interface{}) {

	switch input.(type) {
	case map[string]interface{}:
		tabs += "\t"
		level++
		close := false
		MapElement := input.(map[string]interface{})
		for key, element := range MapElement {
			if element != nil {
				if !IsKeyPresent(key, KeysPresent[level]) {
					KeysPresent[level] = append(KeysPresent[level], key)

					FinalStruct += tabs + strings.Replace(strings.Title(key), "_", "", -1) + " "
					if element != nil {
						TypeObObj := reflect.TypeOf(element).String()
						if TypeObObj == "[]interface {}" {
							TypeObObj = "[]struct {"
							close = true
						} else if TypeObObj == "map[string]interface {}" {
							TypeObObj = "struct {"
							close = true
						} else {
							close = false
						}

						if close {
							FinalStruct += TypeObObj + "\n"
						} else {
							FinalStruct += TypeObObj + "\t`json:\"" + key + "\"`" + "\n"
						}

						IterateOverMap(element)

						if close {
							FinalStruct += tabs + "}\t`json:\"" + key + "\"`" + "\n"
						}
					}
				}
			}
		}

		tabs = tabs[:len(tabs)-1]
		level--

	case []interface{}:
		inputObj := input.([]interface{})
		for _, elementObj := range inputObj {
			//fmt.Print("\nk:", key)
			IterateOverMap(elementObj)
		}

	}
}

//IsKeyPresent if ket present
func IsKeyPresent(i string, list []string) bool {
	for _, element := range list {
		if i == element {
			return true
		}
	}
	return false
}
