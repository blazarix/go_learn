package main

type Person struct {
	Name string
	Age  int
}

type User Person

func main() {
	p1 := Person{
		Name: "kk",
		Age:  12,
	}
	var user User
	user = User(p1)
	print(user)
}
