package main

import (
	"fmt"
	"go-crud/route"
	"go-crud/db/migration"
	
)




func main() {
	fmt.Println("Go run on port 5000")
	migration.Migrate()
	route.HandleRequest()
}
