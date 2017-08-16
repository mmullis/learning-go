package main

import (
	"encoding/json"
	"fmt"
)

type squad struct {
	SquadName  string
	HomeTown   string
	Formed     int
	SecretBase string
	Active     bool
}

func main() {
	input := []byte(`
        {
            "squadName": "Super hero squad",
            "homeTown": "Metro City",
            "formed": 2016,
            "secretBase": "Super tower",
            "active": true,
            "members": [
                {
                    "name": "Molecule Man",
                    "age": 29,
                    "secretIdentity": "Dan Jukes",
                    "powers": [
                        "Radiation resistance",
                        "Turning tiny",
                        "Radiation blast"
                    ]
                },
                {
                    "name": "Madame Uppercut",
                    "age": 39,
                    "secretIdentity": "Jane Wilson",
                    "powers": [
                        "Million tonne punch",
                        "Damage resistance",
                        "Superhuman reflexes"
                    ]
                },
                {
                    "name": "Eternal Flame",
                    "age": 1000000,
                    "secretIdentity": "Unknown",
                    "powers": [
                        "Immortality",
                        "Heat Immunity",
                        "Inferno",
                        "Teleportation",
                        "Interdimensional travel"
                    ]
                }
            ]
        }

	`)

	s := squad{}
	err := json.Unmarshal(input, &s)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", s)
}
