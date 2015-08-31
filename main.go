package main

import (
	"flag"
	"fmt"
)

const version = "0.0.1"

func main() {
	port := flag.Int("port", 3000, "set port (defaults to 3000)")
	dbname := flag.String("dbname", "martini", "set db name")
	flag.Parse()

	db := OpenDatabase(*dbname)
	Server(db).RunOnAddr(fmt.Sprintf(":%d", *port))
}
