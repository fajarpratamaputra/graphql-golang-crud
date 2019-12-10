package query

import (
   "graphql1/schema/types" // sesuaikan dengan folder project
   "github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
  /* list product all
		   http://localhost:8089/graphql?Query={Products{id,name,qty,img}}
		*/

	Name:"Query",
	Fields:graphql.Fields{
		"Products":&graphql.Field{
			Type:graphql.NewList(types.ProductTypes),

			Resolve: ProductResolve,
		},


	},

})
