package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fastjson"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var p fastjson.Parser

	jsonData := `{"foo":"bar", "num":123, "bool":true, "arr":[1, 2, 3]}`

	v, err := p.Parse(jsonData)

	if err != nil {
		panic(err)
	}

	fmt.Printf("foo=%s\n", v.GetStringBytes("foo"))
	fmt.Printf("num=%d\n", v.GetInt("num"))
	fmt.Printf("bool=%v\n", v.GetBool("bool"))

	a := v.GetArray("arr")

	for i, value := range a {
		fmt.Printf("Index: %d: %s\n", i, value)
	}

	jsonDataUser := `{"user": {"name":"John Doe", "age":28}}`
	v, err = p.Parse(jsonDataUser)

	if err != nil {
		panic(err)
	}

	userJSON := v.Get("user").String()

	var user User

	if err := json.Unmarshal([]byte(userJSON), &user); err != nil {
		panic(err)
	}

	fmt.Println(user.Name, user.Age)
}
