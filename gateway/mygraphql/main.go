package mygraphql

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

var queryConfig graphql.ObjectConfig = graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type:    graphql.String,
			Resolve: resolveWorld,
		},
	},
}

type HelloWorld struct {
	Message string `json:"string"`
}

var mutationConfig graphql.ObjectConfig = graphql.ObjectConfig{
	Name: "Hello",
	Fields: graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name: "Params",
				Fields: graphql.Fields{
					"message": &graphql.Field{
						Type: graphql.String,
					},
				},
			}),
			Args: graphql.FieldConfigArgument{
				"message": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				message, _ := params.Args["message"].(string)
				return &HelloWorld{
					Message: message,
				}, nil
			},
		},
	},
}

var schemaConfig graphql.SchemaConfig = graphql.SchemaConfig{
	Query:    graphql.NewObject(queryConfig),
	Mutation: graphql.NewObject(mutationConfig),
}
var Schema, _ = graphql.NewSchema(schemaConfig)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	r := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(r.Errors) > 0 {
		fmt.Printf("エラー: %v", r.Errors)
	}

	j, _ := json.Marshal(r)
	fmt.Printf("%s \n", j)

	return r
}

func resolveWorld(p graphql.ResolveParams) (interface{}, error) {
	return "world", nil
}
