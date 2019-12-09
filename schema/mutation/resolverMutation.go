package mutation

import (
   "graphql1/config"
   "graphql1/schema/types"
   "github.com/graphql-go/graphql"
)



func CreateProductMutation(param graphql.ResolveParams) (interface{},error) {
  var products = []types.Product{}
  product := types.Product{
    Id : param.Args["id"].(int),
    Name : param.Args["name"].(string),
    Qty : param.Args["qty"].(int),
    Img : param.Args["img"].(string),
  }

   db, err := config.GetConnection()
   if err != nil {
      panic(err.Error())
   }
   _ , err = db.Query("INSERT INTO Products values (?,?,?,?)",product.Id,product.Name,product.Qty,product.Img)
   if err != nil {
      panic(err.Error())
   }

   products = append(products, product)
	 return products, nil
}

func UpdateProductMutation(param graphql.ResolveParams) (interface{},error) {
  var a types.Product
  var products []types.Product
   idpro := param.Args["id"].(int)
   nama := param.Args["name"].(string)
   db, err := config.GetConnection()
   if err != nil {
      panic(err.Error())
   }

   _ , err = db.Query("UPDATE Products Set nama_product = (?) Where id = (?)",nama,idpro)
   if err != nil {
      panic(err.Error())
   }

   products = products[:0]
   result,err := db.Query("select id,nama_product,qty_stock,img_product from Products where id = ?",idpro)
   if err != nil{
      panic(err.Error())
   }

   for result.Next(){
      err = result.Scan(&a.Id,&a.Name,&a.Qty,&a.Img)
      if err != nil{
         panic(err.Error())
      }
      products = append(products,a)

   }

   return products , nil
}

func DetailProductMutation(param graphql.ResolveParams) (interface{},error){
   var a types.Product
   var products []types.Product
   idpro := param.Args["id"].(int)
   db ,err:= config.GetConnection()
   if err != nil {
      panic(err.Error())
   }
   products = products[:0]
   result,err := db.Query("select id,nama_product,qty_stock,img_product from Products where id = ?",idpro)
   if err != nil{
      panic(err.Error())
   }

   for result.Next(){
      err = result.Scan(&a.Id,&a.Name,&a.Qty,&a.Img)
      if err != nil{
         panic(err.Error())
      }
      products = append(products,a)

   }

   return products , nil
}

func DeleteProductMutation(param graphql.ResolveParams) (interface{},error) {
  var a types.Product
  var products []types.Product
  idpro := param.Args["id"].(int)
  db ,err:= config.GetConnection()
  if err != nil {
     panic(err.Error())
  }
  products = products[:0]
  result,err := db.Query("select id,nama_product,qty_stock,img_product from Products where id = ?",idpro)
  if err != nil{
     panic(err.Error())
  }

  _ , err = db.Query("DELETE from Products where id = ?",idpro)
  if err != nil{
     panic(err.Error())
  }

  for result.Next(){
     err = result.Scan(&a.Id,&a.Name,&a.Qty,&a.Img)
     if err != nil{
        panic(err.Error())
     }
     products = append(products,a)

  }


  return products , nil
}
