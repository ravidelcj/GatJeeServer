package database

import(

  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/ravidelcj/models"
  "database/sql"
  "strconv"
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
  info.Status = "1"
  exist := rowExist(user)
   if exist {
     query := "select * from user where username = '" + user.Username + "' AND password = '" + user.Password + "';"
     err := Db.QueryRow(query).Scan(&info.Name, &info.Username, &info.Password, &info.Class, &info.LastName)
     if err != nil {
       fmt.Println("CheckUser : ", err)
       return info, false
     }
     return info, true
   }else {
     info.Status = "0"
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
  err := Db.QueryRow(countQuery).Scan(&total)
  if err != nil {
    fmt.Println("Error in retreiving total rows : ", err)
    return -1
  }
  return total
}

func GetRows(page int, classno string) (models.Res, bool) {
  page *= 10
  query := "Select tag, date, title, url from " + classno + " order by id desc limit " + strconv.Itoa(page) + ", 10 ;"
  var res models.Res
  rows, err := Db.Query(query)
  if err != nil {
    fmt.Println("Error in GetRows Query : ",err )
    return res, false
  }
  defer rows.Close()

  var values []models.Value
  for rows.Next() {
    var singleRow models.Value
    err := rows.Scan(&singleRow.Tag, &singleRow.Date, &singleRow.Title, &singleRow.Url)
    if err != nil {
      fmt.Println("Error in retrieving Row : ", err)
      return res, false
    }
    values = append(values, singleRow)
  }
  fmt.Println(values)
  res.Values = values
  return res, true
}

func AddUser(elem models.RegisterUser) bool {
  stmt, err := Db.Prepare("Insert user SET name = ? , username = ? , password = ? , class = ? , last_name = ? ")
  if err != nil {
    fmt.Println("Error Preparing Statement : ", err)
    return false
  }
  _, err1 := stmt.Exec(elem.FirstName, elem.Username, elem.Password, elem.ClassNo, elem.LastName)
  if err1 != nil {
    fmt.Println("Error Inserting into table : ", err)
    return false
  }
  return true
}
