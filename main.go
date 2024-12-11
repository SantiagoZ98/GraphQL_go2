package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	// Definir el tipo "Query"
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "¡Hello World, my name is Santiago Zurita, i'm Student of UCE!", nil
			},
		},
	}

	// Definir el esquema GraphQL
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: fields,
		}),
	})

	if err != nil {
		fmt.Println("Error creando el esquema:", err)
		return
	}

	// Crear un manejador HTTP para GraphQL
	graphQLHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	// Servir el endpoint de GraphQL en el puerto 8080
	http.Handle("/graphql", graphQLHandler)

	fmt.Println("Servidor corriendo en http://localhost:8080/graphql")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
