package database

import(

  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/ravidelcj/models"
  "database/sql"

)

//Global database variable
var Db *sql.DB

func InitDatabase() bool {

  var err error

  Db, err = sql.Open("mysql", "root:admin@/gat_jee")

  if err != nil {
    fmt.Println("Database Connection Error : ", err)
    return false
  }

  err = Db.Ping()
  if err != nil {
    fmt.Println("Database Ping Error : ", err)
    return false
  }
  return true
}

func InsertData(elem models.Element, tableName string ) bool {
   stmt, err := Db.Prepare("Insert " + tableName + " SET title = ? , tag = ? , date = ? , url = ? ")
   if err != nil {
     fmt.Println("Error Preparing Statement : ", err)
     return false
   }
   _, err1 := stmt.Exec(elem.Title, elem.Tag, elem.Date, elem.Url)
   if err1 != nil {
     fmt.Println("Error Inserting into table : ", err)
     return false
   }
   return true
}
