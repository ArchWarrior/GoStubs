package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	
   //"github.com/Sirupsen/logrus"
   "errors"
   "os"
)

func connect() (session *mgo.Session) {
		
	connectURL := "localhost"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	return session
}

func getServicesFromDB() (Services []Service) {
	session := connect()
	defer session.Close()

	collection := session.DB("MockyServices").C("StubsNew")
	fmt.Println("Query Executing\n")

	// Retrieve the list of service.
	//var Services []Service
	err := collection.Find(nil).All(&Services)
	if err != nil {
		fmt.Println("RunQuery : ERROR : %s\n", err)
		return
	}

	fmt.Println("RunQuery:no of services :\n", len(Services))
	
	return Services
}

func saveServiceinDB( s *Service) error{
	session := connect()
	defer session.Close()

	c := session.DB("MockyServices").C("StubsNew")
	
	err := c.Insert(s)

	if err == nil {
		return nil
	}else{
		panic(err)
	return errors.New("Service not added")
	}
}
