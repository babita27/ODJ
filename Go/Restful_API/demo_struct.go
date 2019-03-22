package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

func (u *User) Greeting() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
func main() {
	u := &User{"Babita", "Gurjar"}
	fmt.Println(u.Greeting())
}
