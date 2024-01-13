package main

import (
	"log"

	undine "github.com/igmagollo/undine/pkg/v1"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{}, entc.Extensions(
		undine.Extension{},
	))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
