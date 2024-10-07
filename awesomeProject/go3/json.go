package go3

import (
	"encoding/json"
	"fmt"
)

type MyStr struct {
	Number int    `json:"my,omitempty"`
	Str    string `json:"str,omitempty"`
}

func Json() {
	my := MyStr{44336, "poo-poo-poop"}
	by, err := json.Marshal(my)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(by)
}
