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
