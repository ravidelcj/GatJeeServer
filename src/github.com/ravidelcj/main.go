package main

import (

  "time"
  "fmt"
  "net/http"
  "io"
  "os"
  "github.com/ravidelcj/models"
  "github.com/ravidelcj/database"
)

const (
  Class_No string = "class_no"
  Subject string = "subject"
  Tag string = "tag"
  Title string = "title"
  File string = "file"
)
//handles uploading of file
func uploadNotes(res http.ResponseWriter, req *http.Request){
  fmt.Println("Request Method : ", req.Method)
  if req.Method == "GET" {
    return
  }
  err := req.ParseMultipartForm(32 << 20)
  if err != nil {
    fmt.Println(err)
    return
  }
  var formElem models.Element
  currentTime := time.Now().Local()
  //Initialising Form Element
  file, _, err := req.FormFile(File)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()
  formElem.Title = req.FormValue(Title)
  formElem.Tag = req.FormValue(Tag)
  formElem.Subject = req.FormValue(Subject)
  formElem.Class = req.FormValue(Class_No)
  formElem.Date = currentTime.Format("2006-01-02")
  fmt.Println("class", formElem.Class)
  fmt.Println("tag", formElem.Tag)
  fmt.Println("subject", formElem.Subject)
  fmt.Println("title", formElem.Title)
  fmt.Println("date", formElem.Date)

  urlPath := "/files/" + formElem.Class + "/" + formElem.Title + ".pdf"
  path := "." + urlPath

  out, err := os.Create(path)
  if err != nil {
      fmt.Println(err)
      return
  }
  defer out.Close()

  _, err = io.Copy(out, file)
  if err != nil {
    fmt.Println("Error in uploading : ", err)
    return
  }
  var tableName string
  switch formElem.Class {
    case "8": tableName = "eight"
    case "9": tableName = "nine"
    case "10": tableName = "ten"
    case "11": tableName = "eleven"
    case "12": tableName = "twelve"
  }
  formElem.Url = urlPath
  success := database.InsertData(formElem, tableName)
  if !success {
    fmt.Println("Error while inserting to database")
  }
}

func main() {

  //init Database
  database.InitDatabase()
  defer database.Db.Close()

  //handling Form Data from client
  http.HandleFunc("/getNotes", uploadNotes)

  //Sends Form file to client
  http.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
  })

  http.ListenAndServe(":9000", nil)
}
