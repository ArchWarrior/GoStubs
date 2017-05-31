package main

import (
	"fmt"
	"strings"
	"github.com/clbanning/mxj"
	"math/rand"
	crand "crypto/rand" 
    "time"
    //"os"
    "errors"
      "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
   // "reflect"
)
      
func tagextractor(data []byte, tag string) string{
	m, err := mxj.NewMapXml(data)
	
	if err != nil {
		//fmt.Println("NewMapXml err:", err)
	}
	xval,xerr:=m.ValueForPath(tag)
	if xerr!=nil{
	fmt.Println("value not found")
	}
//	fmt.Println(xval)
	var tagval string
	switch xval.(type){
	 	case string: tagval=xval.(string)
	 	case map[string]interface{}: tagval=xval.(map[string]interface{})["#text"].(string)
	 	default: tagval=""
	}
	if tagval==""{
		return tag
	}else{
		return tagval
	}
}

func addDelay(delay time.Duration,ch chan<- bool){
	time.Sleep(delay * time.Second)
//	fmt.Println("woke up")
	ch <- true
}

func getRandomNumber(min int, max int) int{
	rand.Seed(time.Now().UnixNano())
	rnum:=rand.Intn(max-min)+min
	//fmt.Println(rnum)
    return rnum
}

func getFormattedTimeStamp(format string) string{
	ts:=time.Now().Format(format)//20060102150405.00
	return ts
}

func getFormattedTimeStampWithOffset(format string, d time.Duration, unit string ) string{
	
		var ts string		
	switch unit{
		case "s": ts=time.Now().Add(d*time.Second).Format(format)
		case "m": ts=time.Now().Add(d*time.Minute).Format(format)
		case "h": ts=time.Now().Add(d*time.Hour).Format(format)
		default : panic(errors.New("unit not recognized pass s for seconds m for minutes or h for hour"))
	}
	
		return ts	
	
}

func shuffle(source string) string{
	 rand.Seed(time.Now().UnixNano())
     src:=[]byte(source)
     dest := make([]byte, len(src))
     perm := rand.Perm(len(src))
     for i, v := range perm {
	     dest[v] = src[i]
     }
	return string(dest)
}

func getGUID() string{
	// generate 32 bits timestamp
 	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 12)
	numRead, err := crand.Read(buff)
	
 	if numRead != len(buff) || err != nil {
 		panic(err)
 	}

 	uuid:=fmt.Sprintf("%x-%x-%x-%x-%x-%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
 	return uuid
}



func DBInsertValue(column string,value string,dbName string, collectionName string) string{
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	//c := session.DB("test").C("CafeDB")
	c := session.DB(dbName).C(collectionName)
	
	err = c.Insert(bson.M {column: value})
	if err != nil {
		panic(err)
	}
	return "Success"
}

		
//Function to fetch a value from DB decided by the a value in the request
	
	func DBFetch(fetchDetails string,requestValue string,matchColumnName string) string{
	
	z := strings.Split(fetchDetails, ",")
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(z[1]).C(z[2])

	//fmt.Println("Macthing value",requestValue)
	var m bson.M
	err = c.Find(bson.M{matchColumnName: requestValue}).One(&m)
	//fmt.Println("Search Names:",matchColumnName)
	//panic(err)
	
	for key, value := range m{
		if key ==z[0]{//comparing the DB key whose value will match the value sent in z[0]

    return value.(string)
		}
	}
	return "Value Not Found"
}   
	
// Function to do dynamic correlation,i.e. to construct the response dynamically decided by the number of occurence of a 
//tag in the incoming request
 	
	func tagextractorForArray(data []byte,tag string, repeatStringFromReq string, occurence int, dynamicValueToReplace string, tagtocapture string )string {
		m, err := mxj.NewMapXml(data)
	   if err != nil {
			
	}
	   
	   	
	/*childTagNames:=strings.Split(repeatStringFromReq, ".")
			    		childTagName:=childTagNames[len(childTagNames)-1]//WRT string last child tag name
			    		//fmt.Println("WRT String:",repeatStringFromReq)
						//fmt.Println("child tag name:","<"+childTagName+">")  
						parentXpath:=strings.TrimRight(repeatStringFromReq,"<"+childTagName+">")
						//fmt.Println("Parent tag name:",parentXpath)   
						*/
	
    values, _ := m.ValuesForPath(repeatStringFromReq+".*")
    
    var str string
    
    switch values[occurence].(type) {
 	case string: str, _ = values[occurence].(string)
 	case map[string]interface{}: str=values[occurence].(map[string]interface{})[tagtocapture].(string)
 	default: str="could not found"
 }
    
    
   	
    repeatTagArray:=strings.Split(tag, "Repeat(")
    repeatTag:=strings.Split(repeatTagArray[1], ")")
    returnTag:=strings.Replace(repeatTag[0],"@@"+dynamicValueToReplace+"@@", str, -1)
      
    return returnTag
   
}
