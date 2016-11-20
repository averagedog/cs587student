package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	"log"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error opening %v: %v\n", path, status)
	}


	data    := MyType { "foo", "bar" }

	data.Write(fd)
	data.WriteVar(path +"foobar")

}
