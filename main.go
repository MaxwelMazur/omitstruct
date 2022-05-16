package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string   `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	Number      int    `json:"number,omitempty"`
	Description string `json:"description,omitempty"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string   `json:"name,omitempty"`
		Address *Address `json:"address,omitempty"`
	}{
		Name:    u.Name,
		Address: u.Address,
	})
}

func (r *User) UnmarshalJSON(b []byte) error {
	decoded := new(struct {
		Name    string   `json:"name,omitempty"`
		Address *Address `json:"address,omitempty"`
	})
	err := json.Unmarshal(b, decoded)
	if err == nil {
		r.Name = decoded.Name
		r.Address = decoded.Address
	}
	return err
}

func main() {
	// str := `{"name": "pet", "address": {"number": 123, "description": "home"}}`
	str := `{"name": "pet"}`
	user := &User{}
	err := user.UnmarshalJSON([]byte(str))
	if err != nil {
		panic(err)
	}

	b, err := user.MarshalJSON()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
