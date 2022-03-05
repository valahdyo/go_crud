package main

import (
	"fmt"
	"go-crud/db/migration"
	"go-crud/route"
)

func main() {
	fmt.Println("Go run on port 5000")
	migration.Migrate()
	route.HandleRequest()
}
