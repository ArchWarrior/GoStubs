package main

import (
    "fmt"
	"net/http"
	"io/ioutil"
	"strings"
	 "math/rand"
   "time"
    "github.com/Sirupsen/logrus"
)

func getNewhandler(service Service) http.HandlerFunc {
	//fmt.Println("inside getHandler")
	return func (w http.ResponseWriter, r *http.Request) {
			//fmt.Println("inside handeler")
			start:=time.Now()
			var operation Operation
			var response string
			var evaluatedIPVariables map[string]string
			 w.Header().Set("content-type", service.Type)
			 hah, err := ioutil.ReadAll(r.Body)
	         if err != nil {
		       fmt.Fprintf(w, "%s", err)
	         }
	         reqbody := string(hah)
//	         fmt.Println("-----reqbody:")
//	         	fmt.Println(reqbody)
	         for _,soperation := range service.Operations {
	         	//fmt.Println(soperation.Opname)
			if(strings.Contains(reqbody, soperation.Opname)||strings.Contains(r.RequestURI, soperation.Opname)){
				operation=soperation
			}
			}
	         if operation.Monitoring{
		monitoringRequestLogger(start,r.Header,reqbody,operation.Opname,r.RequestURI)
	}
//			  opname := ""
			  if (operation.Opname!=""){
			  	// Request validation(to check whether the request has correct operation name or any essential data)
//			  	ch := make(chan bool)
//			    go addDelay(operation.Delay,ch)
			rseed:=time.Now().UnixNano()
			rand.Seed(rseed)
			
	         //output:=operation.Output
	         logrus.WithFields(logrus.Fields{"configured outputs":operation.Outputs}).Debug()
	          for _,output := range operation.Outputs {
	          	logrus.WithFields(logrus.Fields{"Output":output.Response}).Debug()
	          	if(strings.Contains(reqbody, output.Tagvalue)){
		          	 response=output.Response
			    }
	          	evaluatedIPVariables=evaluateInputVariables(output.Variables,hah)
	          	
	          	 //fmt.Println("rawmap:",output.Variables,"evaluated:",evaluatedIPVariables)
			}
	         
	       
	       
	

		
//			
//
			  
			startdelimiter:= "${"
			enddelimeter:="}"
           // fmt.Println(strings.Contains(output, delimiter))

             for strings.Contains(response, startdelimiter){
             
	         z := strings.SplitN(response, startdelimiter,2)
	         y := strings.SplitN(z[1], enddelimeter,2)
//	         fmt.Println(y[0],"response splits...")
	         response=z[0]+evaluatedIPVariables[y[0]]+y[1]
	              
   
			}  
//               
//			 <-ch
			 time.Sleep(time.Duration(operation.Delay)* time.Second)
			  //fmt.Println(response)
			//fmt.Fprintf(w, response)
	} else {
		response=`operation not found`
		

	}
	fmt.Fprintf(w, response)
	
	if operation.Monitoring{
		monitoringResponseLogger(time.Now(),w.Header(),response)
	}

		
			        
    }
}