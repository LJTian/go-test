package main

func main() {

	shell := "none"
	data := ""
	if shell == "none" {

		// This is a simple switch statement
		switch data {
		case "":
			println("one")
			//break // Useless break statement
		case "two":
			println("two")
		default:
			println("default")
		}

		println("if shell is none")
	}
}
