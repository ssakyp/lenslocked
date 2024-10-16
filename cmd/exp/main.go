package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
	Date []float64
	Address map[string]string
	Family []string
	Dad string
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
		Age:  125,
		Date: []float64{24.1996, 02.1962, 06.1963},
		Address: map[string]string{
			"City": "Almaty",
			"Street": "Abay",
			"House": "24",
		},
		Family: []string{"Son", "Dad", "Mom"},
		Dad: "Dad",
	}

	err = t.Execute(os.Stdout, user)

	if err != nil {
		panic(err)
	}
}
