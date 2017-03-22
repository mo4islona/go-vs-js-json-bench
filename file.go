package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"time"
)

type Collection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Properties Properties `json:"properties"`
	Geometry   Geometry   `json:"geometry"`
}

type Properties struct {
	MAPBLKLOT string `json:"MAPBLKLOT"`
	BLKLOT    string `json:"BLKLOT"`
	BLOCK_NUM string `json:"BLOCK_NUM"`
	LOT_NUM   string `json:"LOT_NUM"`
	FROM_ST   string `json:"FROM_ST"`
	TO_ST     string `json:"TO_ST"`
	STREET    string `json:"STREET"`
	ST_TYPE   string `json:"ST_TYPE"`
	ODD_EVEN  string `json:"ODD_EVEN"`
}

type Geometry struct {
	Type        string      `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

func main() {
	bytes, err := ioutil.ReadFile("./citylots.json")
	if err != nil {
		fmt.Printf("Error read file: %v\n", err)
		return
	}

	data := new(Collection)
	{
		start := time.Now()
		err = json.Unmarshal(bytes, data)
		elapsed := time.Since(start)
		fmt.Printf(" [x] Golang unmarshal: %s\n", elapsed)
		if err != nil {
			fmt.Printf("Error unmarshal: %v\n", err)
			return
		}
	}
	{
		start := time.Now()
		_, err = json.Marshal(data)
		elapsed := time.Since(start)
		fmt.Printf(" [x] Golang marshal: %s\n", elapsed)
		if err != nil {
			fmt.Printf("Error unmarshal: %v\n", err)
			return
		}
	}
}
