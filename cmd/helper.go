package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func LoadData() ConList {

	if _, err := os.Stat(DATA_PATH); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		// fmt.Println("file not exist")
		return ConList{}

	}
	plan, err := ioutil.ReadFile(DATA_PATH)
	// plan, err := ioutil.ReadFile("output.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return ConList{}
	}

	var data ConList
	err = json.Unmarshal(plan, &data)
	if err != nil {
		fmt.Println("error occurred")
		fmt.Println(err)
	}
	// fmt.Println(data)
	for i, v := range data.List {
		SshIndexMap[i] = v
		SshShortMap[v.Shortcut] = v
	}
	// fmt.Println(SshIndexMap)
	// fmt.Println(SshShortMap)
	return data
}

func SaveData(data ConList) {
	// fmt.Println(data)
	// reqBodyBytes := new(bytes.Buffer)

	// json.NewEncoder(reqBodyBytes).Encode(data)
	// fmt.Println(reqBodyBytes.Bytes())

	if _, err := os.Stat(DATA_PATH); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		// fmt.Println("file not exist")
		f, e := os.Create(DATA_PATH)

		if e != nil {
			fmt.Println(e)
		}
		f.Close()

	}

	jsonData, _ := json.Marshal(data)
	// fmt.Println(jsonData)
	err := ioutil.WriteFile(DATA_PATH, jsonData, 0666)
	if err != nil {
		fmt.Println(err)
	}

}

func getConnection(data string) []string {
	// fmt.Println("in getConnection")
	keys := []string{}
	for k, _ := range SshShortMap {

		m, err := regexp.MatchString(`.*`+data+`.*`, k)
		if err != nil {
			fmt.Println(err)
		}
		if m {
			// fmt.Println("found: ", m, k)

			keys = append(keys, k)
		}
	}
	fmt.Println(keys)

	return keys
}
