package main

import (
	"net/http"
//    "fmt"
)
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

//func initializeRoutes(config Config) {
//	for _,service := range config.Services {
//	//fmt.Println(reflect.TypeOf(service))
//		for _,operation:= range service.Operations {
//			r:=Route{operation.Opname,
//                     "POST",
//                     operation.Path,
//                     getNewhandler(operation)}
//			routes=append(routes, r)
//		//fmt.Println(operation.Opname+operation.Path)
//	}
//
//}
	func initializeRoutes(Services []Service) {
	for _,service := range Services {
	//fmt.Println(reflect.TypeOf(service))
			r:=Route{service.Sname,
                     "POST",
                     service.Path,
                     getNewhandler(service)}
			routes=append(routes, r)
			//fmt.Println(service.Operations[0].Outputs[0].Formats)
//			fmt.Println(service.Operations[0].Outputs[0].Response)
		//fmt.Println(operation.Opname+operation.Path)
	}

}

	
//routes=append(routes, Route{"appended",
//                             "GET",
//                             "/appended",
//                              TodoIndex})

type Routes []Route

var routes Routes
    
//    Route{
//    "PolarisService",
//    "POST",
//    "/PolarisService/createSurvey",
//    PolarisService,
//    },
//}
