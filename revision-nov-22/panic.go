package main

func crash() {
	panic("unexpected error occured.")
}

func main() {
	crash()

}
