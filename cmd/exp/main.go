package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// user := struct {
	// 	Name string
	// }{
	// 	Name: "Sultan Sakyp",
	// }

	user := User{
		Name: "Sultan Sakyp",
		Bio:  `<script>alert("Haha, you have been h4x0rt3d!");</script>`,
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
}
