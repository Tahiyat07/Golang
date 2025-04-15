package main

import "fmt"

///const age = 30

// shorthand := not possible
func main() {
	//const name = "golang"
	//*name cannot assign
	// name =" js"
	///fmt.Println(name)

	//*outter const
	///fmt.Println(age)

	//*multiple
	const (
		port = 9000
		host = "yu"
	)
	fmt.Println(port, host)

}
