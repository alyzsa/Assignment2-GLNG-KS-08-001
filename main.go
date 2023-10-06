package main

import(
	"github.com/alyzsa/Assignment2-GLNG-KS-08-001/database"
	"github.com/alyzsa/Assignment2-GLNG-KS-08-001/routers"
)

func main() {
    database.StartDB()
    var PORT = ":8083"
    routers.StartServer().Run(PORT)
}