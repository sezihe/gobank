package main

func main() {
	server := NewAPIServer(":3050")

	server.Run()
}
