package query
import (
   "github.com/graphql-go/graphql"
   "graphql1/schema/types"
   "graphql1/config"
)

func ProductResolve(param graphql.ResolveParams) (interface{},error){
   var a types.Product
   var products []types.Product
   db ,err:= config.GetConnection()
   if err != nil {
      panic(err.Error())
   }
   products = products[:0]
   result,err := db.Query("select id,nama_product,qty_stock,img_product from Products")
   if err != nil{
      panic(err.Error())
   }

   for result.Next(){
      err = result.Scan(&a.Id,&a.Name,&a.Qty,&a.Img)
      if err != nil{
         panic(err.Error())
      }
      products = append(products, a)

   }


   return products , nil
}
