package models

type Element struct{
  Title string
  Class string
  Subject string
  Tag string
  Date string
  Url string
}

type User struct {
  Name string
  Username string
  Password string
  Class string
}

type ClientUser struct {

  Username string
  Password string

}
