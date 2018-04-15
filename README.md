# Go MLB
[![Go Report Card](https://goreportcard.com/badge/github.com/stevepartridge/mlb)](https://goreportcard.com/report/github.com/stevepartridge/mlb)
[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/stevepartridge/mlb) 
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/stevepartridge/mlb/master/LICENSE) 
[![Coverage](http://gocover.io/_badge/github.com/stevepartridge/mlb)](http://gocover.io/github.com/stevepartridge/mlb)
[![Build Status](https://travis-ci.org/stevepartridge/mlb.svg?branch=master)](https://travis-ci.org/stevepartridge/mlb) 

API Client for MLB statsapi.mlb.com

### Example Usage

```golang

import(
  "fmt"
  "time"

  "github.com/stevepartridge/mlb"
)

mlbApi, err := mlb.New()
if err != nil {
  fmt.Println("aww, bad things happened", err.Error())
}

teams, err := mlbApi.GetTeams()
if err != nil {
  fmt.Println("well, dang", err.Error())
}

for i := range teams {
  fmt.Println(teams[i].Name)
}

start, _ := time.Parse("2006/01/02", "2017/05/17")
end := start.AddDate(0, 1, 0) // one month

games, err := mlbApi.GetGamesForRange(start, end)
if err != nil {
  fmt.Println("well, dang", err.Error())
}

for i := range games {
  fmt.Printf("----- \n %+v \n-----\n\n", games[i])
}


```

##### Work in progress
This is a continued work in progress.  For any of the unimplemented endpoints using the ```mlbApi.Call(...)``` method should work with any of them.

The full documention of the [MLB API](http://statsapi.mlb.com/docs/).
