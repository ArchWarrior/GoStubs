package main

import (
	
	"strconv"
	"strings"
	"time"
	//"fmt"
	"github.com/clbanning/mxj"
	"bytes"
	

)
func evaluateInputVariables(rawmap map[string]string, data []byte) map[string]string{
 //fmt.Println("inside IP evaluation")
 varmap := make(map[string]string)
	for k,ip:= range rawmap{
		
	
					//Common function :1
			    	varmap[k]=ip
			    	if strings.Contains(ip, "getRandomNumber") {
				    	z := strings.Split(ip, "(")
				    	z[1]=strings.TrimRight(z[1], ")")
				    	limit:=strings.Split(z[1],",")
				    	min,_:=strconv.Atoi(limit[0])
				    	max,_:=strconv.Atoi(limit[1])
					    varmap[k]=strconv.Itoa(getRandomNumber(min,max))
			    		
			    	}else//Common function :2
			    	if strings.Contains(ip,"getFormattedTimeStampWithOffset"){
			    		
			    		z := strings.Split(ip, "(")
			    		z[1]=strings.TrimRight(z[1], ")")
			    		y:=strings.Split(z[1],",")
			    	
			    		duration,_:=strconv.Atoi(y[1])
			     		varmap[k]=getFormattedTimeStampWithOffset(y[0], time.Duration(duration), y[2])

			    	}else//Common function :3
			    	if strings.Contains(ip, "getFormattedTimeStamp"){
			    		z := strings.Split(ip, "(")
				    	z[1]=strings.TrimRight(z[1], ")")
				    	varmap[k]=getFormattedTimeStamp(z[1])
				    	
					}else//Common function :4
			    	if strings.Contains(ip, "shuffle"){
			    		z := strings.Split(ip, "(")
				    	z[1]=strings.TrimRight(z[1], ")")
			    		varmap[k]=shuffle(z[1])
			    	}else//Common function :5
			    	if strings.Contains(ip,"getGUID"){
			    		varmap[k]=getGUID()
			    	}else//Common function :6
			    	if strings.Contains(ip,"DBInsertValue"){
			    		z := strings.Split(ip, "(")
				    	z[1]=strings.TrimRight(z[1], ")")
				    	y:=strings.Split(z[1], ",")
			    		value :=tagextractor(data, y[0])
			    		varmap[k]=DBInsertValue(k,value,y[1],y[2])
			    	}else//Common function :7
			    	if strings.Contains(ip,"DBFetch"){
			    		z := strings.Split(ip, "(")
				    	y := strings.Split(z[1], ")") //y[0]
				    	x := strings.Split(ip, "when(")
				    	w := strings.Split(x[1], ")") //w[0]
				    	value :=tagextractor(data, w[0])
				    	v := strings.Split(ip, "matches(")
				    	u := strings.Split(v[1], ")") //u[0]
			    		
			    		varmap[k]=DBFetch(y[0],value,u[0])
			    	}else //Common function :8
			    	if strings.Contains(ip,"Repeat"){
			    		
			    		var returnString bytes.Buffer
			    		z := strings.Split(ip, "wrtTag(")
			    		repeatStringFromReq:=strings.TrimRight(z[1], ")")//WRT string
			    		//fmt.Println("WRT String:",repeatStringFromReq)
			    		
			    		dynamicCorelation:=strings.Split(ip, "@@")
						dynamicValueToReplace:=strings.Split(dynamicCorelation[1], "@@")
						lastTagNames:=strings.Split(dynamicValueToReplace[0], ".")
						length:=len(lastTagNames) 
						lastTagName:=lastTagNames[length-1] //To get the last tag name
							
						
						m, err := mxj.NewMapXml(data)
						if err != nil {
					//fmt.Println("NewMapXml err:", err)	
							}
						values, _ := m.ValuesForPath(repeatStringFromReq+".*")//  Envelope.Body.AddBVEUids.addBVEUidInfo.bveUIDInfo.bveUid.*
						//fmt.Println("values:",values)
						//fmt.Println("Occurence, len:", len(values))
						occurence:=len(values)
						//fmt.Println("Occurence",occurence)
			    		
				   
				     	for i:=1;i<occurence;i++{	  
						returnString.WriteString(tagextractorForArray(data ,ip,repeatStringFromReq ,i, dynamicValueToReplace[0], lastTagName))
						}
			    		
			    		varmap[k]=returnString.String()
			    	}else{
//			    		fmt.Println("before tag extraction")
			    		varmap[k]=tagextractor(data, ip)
//			    		fmt.Println("value for tag")
//			    		fmt.Println(varmap[k])
			    	}
   	}
  //	fmt.Println(varmap)
  	return varmap
} 

