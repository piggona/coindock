package main

import "vct/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
