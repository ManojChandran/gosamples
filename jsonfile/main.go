// raeding json file
package main

import (
  "fmt"
  "os"
  "encoding/json"
)

type Query struct {
  TableName string `json:TableName`
  AttributeDefinition [] struct {
    AttributeName string `jason:AttributeDefinition`
    AttributeType string `json:AttributeType`
  }`json:AttributeDefinition`
  KeySchema [] struct {
    AttributeName string `json:AttributeName`
    KeyType string `json:KeyType`
  }`json:KeySchema`
  ProvisionedThroughput struct {
    ReadCapacityUnits int `json:ReadCapacityUnits`
    WriteCapacityUnits int `json:WriteCapacityUnits`
  }`json:ProvisionedThroughput`
} // Query

func loadQuery(filename string) (Query, error)  {
  var query Query
  queryFile,err := os.Open(filename)
  defer queryFile.Close()
  if err !=nil {
    return query, err
  }
  jsonParser := json.NewDecoder(queryFile)
  err = jsonParser.Decode(&query)
  return query, nil
} // loadQuery

func main() {
  query,_ := loadQuery("Query.json")

  for _,v := range query.AttributeDefinition{
    fmt.Println(v.AttributeName)
  }
  for _,v := range query.KeySchema{
    fmt.Println(v.AttributeName)
  }
}
