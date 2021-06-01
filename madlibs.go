package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type Inventory struct {
	Material string `json: "Material"`
	Count    uint   `json: "Count"`
}

// Read a json file containing an inventory struct and parse it
func ReadFile(filename string) (Inventory, error) {
	file, _ := ioutil.ReadFile(filename)
	sweaters := Inventory{}

	_ = json.Unmarshal([]byte(file), &sweaters)

	if err := json.Unmarshal([]byte(file), &sweaters); err != nil {
		log.Println(err)
		return sweaters, err
	}

	return sweaters, nil
}

func main() {

	sweaters, err := ReadFile("inventory.json")
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		log.Fatal(err)
	}
}
