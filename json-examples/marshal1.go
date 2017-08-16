package main

import (
	"encoding/json"
	"fmt"
)

type country struct {
	Name string
}

func main() {
	country_code := "US"

	the_country := country{}

	switch country_code {
	case "US":
		the_country.Name = "United States"
	}

	c, err := json.Marshal(the_country)
	if err != nil {
		panic(err)
	}

	// c is now a []byte containing the encoded JSON
	fmt.Print(string(c))
}
