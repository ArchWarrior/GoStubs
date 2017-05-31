package main

import (
    "net/http"
    "fmt"
    "github.com/gorilla/mux"
    // "reflect"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    
    for _, route := range routes {
        var handler http.HandlerFunc
        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
         fmt.Println("Endpoint created at path:"+route.Pattern)

    }
//    handler:= getNewhandler()
//    handler = Logger(handler,"polarisConfig" )`
//    fmt.Println(reflect.TypeOf(handler))
//    router.
//            Methods("POST").
//            Path("/PolarisService/createSurvey").
//            Name("polarisConfig").
//            Handler(handler)
//         fmt.Println("Endpoint created at path:"+"/PolarisService/createSurvey")
    return router
}