package mutation

import (
   "graphql1/schema/types"
   "github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
   Name:"Mutation",
   Fields:graphql.Fields{
      "CreateProducts":&graphql.Field{
         Type:graphql.NewList(types.ProductTypes),
         //config param argument
         Args:graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
               Type:graphql.NewNonNull(graphql.Int),
            },
            "name": &graphql.ArgumentConfig{
               Type:graphql.NewNonNull(graphql.String),
            },
            "qty": &graphql.ArgumentConfig{
               Type:graphql.NewNonNull(graphql.Int),
            },
            "img": &graphql.ArgumentConfig{
               Type:graphql.NewNonNull(graphql.String),
            },
         },
         Resolve: CreateProductMutation ,
      },
      "UpdateProducts":&graphql.Field{
         Type:graphql.NewList(types.ProductTypes),
         //config param argument
         Args:graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
               Type:graphql.NewNonNull(graphql.Int),
            },
            "name": &graphql.ArgumentConfig{
               Type:graphql.NewNonNull(graphql.String),
            },

         },
         Resolve: UpdateProductMutation ,
      },
      "DetailProducts":&graphql.Field{
         Type:graphql.NewList(types.ProductTypes),
         //config param argument
         Args:graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
               Type:graphql.NewNonNull(graphql.Int),
            },
         },
         Resolve: DetailProductMutation ,
      },
      "DeleteProducts":&graphql.Field{
         Type:graphql.NewList(types.ProductTypes),
         //config param argument
         Args:graphql.FieldConfigArgument{
            "id": &graphql.ArgumentConfig{
               Type:graphql.NewNonNull(graphql.Int),
            },
         },
         Resolve: DeleteProductMutation,
      },
      // untuk membuat object lainya tinggal di ulang

   },

})
