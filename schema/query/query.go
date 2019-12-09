package query

import (
   "graphql1/schema/types" // sesuaikan dengan folder project
   "github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:"Query",
	Fields:graphql.Fields{
		"Products":&graphql.Field{
			Type:graphql.NewList(types.ProductTypes),
			//config param argument

			Resolve: ProductResolve,
		},
		// untuk membuat object lainya tinggal di ulang

	},

})
