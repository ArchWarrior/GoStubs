package main

import (
    "log"// "github.com/Sirupsen/logrus"
    "net/http"
    "time"
    "os"
//    "net/http/httptest"
   // "github.com/mattn/go-colorable"
//   "net/http/httputil"
   "fmt"
//   "io/ioutil"
   "strings"
)

func monitoringResponseLogger(time time.Time,header http.Header,body string){

f, err := os.OpenFile("Monitor.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
                log.Fatal(err)
        }   

        //defer to close when you're done with it, not because you think it's idiomatic!
        defer f.Close()
        
        log.SetOutput(f)
        headers:=fmt.Sprintf("%v", header)
     
        headers=strings.Trim(headers, "map")
        
       
        log.Printf(
            "Response sent at %s\nHeaders: %s\nBody: %s\n\n",
           time,
            headers,
            body,
        )
	//	        resp := rec.Result()
//	   body, _ := ioutil.ReadAll(resp.Body)
//
//	   fmt.Println(resp.StatusCode)
//	fmt.Println(resp.Header.Get("Content-Type"))
//	fmt.Println(string(body))
}
func monitoringRequestLogger(time time.Time, header http.Header, body string,operationname string, url string){
f, err := os.OpenFile("Monitor.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
                log.Fatal(err)
        }   

        //defer to close when you're done with it, not because you think it's idiomatic!
        defer f.Close()
        
        log.SetOutput(f)
        headers:=fmt.Sprintf("%v", header)
     
        headers=strings.Trim(headers, "map")
        
       
        log.Printf(
            "Requested recieved at %s\nURL: %s \nHeaders: %s\nBody: %s\nServed By Operation: %s\n\n",
           time,
           url,
            headers,
            body,
            operationname,
        )
}

func Logger(inner http.HandlerFunc, name string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
//         x, err := httputil.DumpRequest(r, true)
       
        
        inner(w, r)
        

       f, err := os.OpenFile("Access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
                log.Fatal(err)
        }   

        //defer to close when you're done with it, not because you think it's idiomatic!
        defer f.Close()
//        log.SetFormatter(&log.TextFormatter{ForceColors: true})
//		log.SetOutput(colorable.NewColorableStdout())

//		log.Info("succeeded")
//		log.Warn("not correct")
//		log.Error("something error")
		//log.Fatal("panic")
        //set output of logs to f
       log.SetOutput(f)
//          log.SetOutput(os.Stdout)
        //test case
        //log.Println("check to make sure it works")
//       fmt.Println()
         

        log.Printf(
            "%s %s %f",
            r.Method,
            r.RequestURI,
            time.Since(start).Seconds(),
        )
//        log.WithFields(log.Fields{
//                                 "method": r.Method ,
//                                 "path": r.RequestURI,
//                                 "responseTime": time.Since(start).Seconds(),
//                                  }).Info()
//        log.Info(r.Method,
//            r.RequestURI,
//            time.Since(start).Seconds())
}
}