package main

import (
//    "github.com/BurntSushi/toml"
  //"fmt"
 // "strings"
 //"reflect"
 //"time"
 "gopkg.in/mgo.v2/bson"
)
//type tomlConfig struct {
//	Title   string
//}


type Config struct {
 Services []Service 
 
}

type Service struct{
	ID    bson.ObjectId `bson:"_id,omitempty"`
	Sname string //`bson:"Sname"`
	Path string //`bson:"EndPoint"`
	Type string //`bson:"Type"`	
	Operations []Operation //`bson:"Operations"`
	

}

type Operation struct{
	Opname string //`bson:"Opname"`
	Delay int// `bson:"Delay"`
	Outputs []Output // `bson:"Output"`
	Monitoring bool 
	//Path string
	
	
}
type Output struct{
	Variables map[string]string
	Tagvalue string //`bson:"TagName"`
	Response string //`bson:"Response"`
}

//type Variable struct{
//	
// Var map[string] input
//}
//
type inputs struct {
	Input string
}