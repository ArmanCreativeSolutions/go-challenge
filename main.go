package main

import "go-challenge/infrastructure/dbconfig"

func main() {
	dbconfig.Initialize("test.db")
}
