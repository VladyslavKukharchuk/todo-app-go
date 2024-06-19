package main

import (
	"ariga.io/atlas-provider-gorm/gormschema"
	"fmt"
	"io"
	"os"
	"todo-app-go/pkg/model"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(&model.User{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
