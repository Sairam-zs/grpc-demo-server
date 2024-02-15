package main

import (
	"github.com/zopping/grpc-demo/grpc/employee"
	"gofr.dev/pkg/gofr"
)

func main() {
	g := gofr.New()

	employeeServer := employee.New(g)

	employee.RegisterEmployeeServiceServer(g.Server.GRPC.Server(), employeeServer)

	g.Start()
}
