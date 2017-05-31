package main

import (
    "log"
    "net/http"
   "fmt"
     "github.com/BurntSushi/toml"//For TOML
      "github.com/Sirupsen/logrus"
      "os/exec"  
)

func main() {
	
	//To Launch the User Guide
command:=exec.Command("cmd","/C","UserGuide\\VirtualizeQD_UserGuide.html")
	if err :=  command.Run(); err != nil { 
	    fmt.Println("Error: ", err)
    }
	
	logrus.SetLevel(logrus.ErrorLevel)
    var config Config
    if _, err := toml.DecodeFile("config.toml", &config); err != nil {
   panic(err)
  }//For TOML
   
   
//    Services:=getServicesFromDB()//For DB
//    saveServiceinDB(&config.Services[0])
  //   fmt.Println(Services[1].Operations[0].Outputs[0].Variables)
    initializeRoutes(config.Services)//For TOML
//     initializeRoutes(Services) //for DB
    router := NewRouter()
    port :="8383"
    
  
     fmt.Println("Server is about to listen at PORT :"+port)
    log.Fatal(http.ListenAndServe(":"+port, router))


      
}