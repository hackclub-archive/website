package main

import (
	"github.com/go-martini/martini"
	"github.com/hackedu/backend/v1"
)

func main() {
	m := martini.Classic()

	v1.Setup(m)
	defer v1.Dbm.Db.Close()

	m.Run()
}
