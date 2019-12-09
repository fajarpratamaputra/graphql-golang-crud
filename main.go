package main

import (
   "graphql1/schema/query"
   "encoding/json"
   "fmt"
   //"github.com/go-chi/render"
   "github.com/graphql-go/graphql"
   "graphql1/schema/mutation"
   "net/http"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
   result := graphql.Do(graphql.Params{
      Schema:   schema,
      RequestString: query,
   })
   if len(result.Errors) > 0 {
      fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
   }
   return result
}

func main(){

   var Schame,err = graphql.NewSchema(graphql.SchemaConfig{
      Query:query.RootQuery,
    //  Fuery:query.RootFuery,
      Mutation: mutation.Mutation,
   })
   if err != nil{
      panic(err.Error())
   }
   http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
      w.Header().Set("Content-Type", "application/json")
      result := executeQuery(r.URL.Query().Get("Query"), Schame)
      json.NewEncoder(w).Encode(result)

   })
   fmt.Println("Server is running on port 8089")
	  http.ListenAndServe(":8089", nil)

}
