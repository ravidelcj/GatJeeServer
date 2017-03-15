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
  Status string
  Name string
  Username string
  Password string
  Class string
  LastName string
}

type ClientUser struct {
  Username string `json:,string`
  Password string
}

type Value struct {
  Title string
  Tag string
  Date string
  Url string
}

type Res struct {
  Status int
  Values []Value
}

type RegisterUser struct{
  FirstName string
  LastName string
  Username string
  Password string
  ClassNo string
}
