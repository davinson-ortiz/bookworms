package main

import "os"

func Example_main() {
	// Simulamos pasar argumentos desde la CLI
	os.Args = []string{"cmd", "-path=testdata/testBookworms.json"}
	main()
	// Output:
	// Here are books in common:
	// - The Handmaid's Tale by Margaret Atwood
}
