package mutation

import (
   "graphql1/schema/types"
   "github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
   Name:"Mutation",
   Fields:graphql.Fields{
     /* create product
         http://localhost:8089/graphql?Query=mutation+{CreateProducts(id:6,name:"Test Json",qty:25,img:"image3"){id,name,qty,img}}
      */
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
      /* update product
          http://localhost:8089/graphql?Query=mutation+{UpdateProducts(id:3,name:"Tester"){id,name,qty,img}}
       */
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
      /* detail product by id
          http://localhost:8089/graphql?Query=mutation+{DetailProducts(id:5){id,name,qty,img}}
       */
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
      /* delete product
          http://localhost:8089/graphql?Query=mutation+{DeleteProducts(id:6){id,name,qty,img}}
       */
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


   },

})
