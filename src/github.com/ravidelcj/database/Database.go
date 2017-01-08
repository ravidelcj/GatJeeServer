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

func CheckUser(user models.ClientUser) (models.User, bool) {
  var info models.User
  exist := rowExist(user)
   if exist {
     query := "select * from user where username = '" + user.Username + "' AND password = '" + user.Password + "';"
     err := Db.QueryRow(query).Scan(&info.Name, &info.Username, &info.Password, &info.Class)
     if err != nil {
       fmt.Println("CheckUser : ", err)
       return info, false
     }
     return info, true
   }else {
     return info, false
   }
}

//check whether user exist in database
func rowExist(user models.ClientUser) bool {
  query := "Select exists(select 1 from user where username = '" + user.Username + "' AND password = '" + user.Password + "');"
  var exist bool
  err := Db.QueryRow(query).Scan(&exist)
  if err != nil {
    fmt.Println("rowExist : ", err)
    return false
  }
  if exist {
    return true
  }else{
    return false
  }
}

func TotalRows(classno string) int {
  countQuery := "Select Count(*) from " + classno + ";"
  var total int
  err := Db.Query(countQuery).Scan(&total)
  if err != nil {
    fmt.Println("Error in retreiving total rows : ", err)
    return -1
  }
  return total
}
