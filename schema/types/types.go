package types
import (

"github.com/graphql-go/graphql"

)

var ProductTypes = graphql.NewObject(graphql.ObjectConfig{
   Name:"Products",
   Fields:graphql.Fields{
      "id":&graphql.Field{
         Type:graphql.Int,
      },
      "name":&graphql.Field{
         Type:graphql.String,
      },
      "qty":&graphql.Field{
         Type:graphql.Int,
      },
      "img":&graphql.Field{
         Type:graphql.String,
      },
   },
})
