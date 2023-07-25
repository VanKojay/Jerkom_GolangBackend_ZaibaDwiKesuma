// Main.go
package main

import (
	"section11.ORM-CodeStructure/config"
	"section11.ORM-CodeStructure/routes"
)

// "project/config"
// "project/routes"

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
