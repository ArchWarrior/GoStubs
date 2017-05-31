package main

import (
	"fmt"
	"bytes"
	"github.com/clbanning/mxj"
	 "math/rand"
   "time"
   "strconv"
   "strings"
    "github.com/Sirupsen/logrus"
      "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
   // "github.com/mattn/go-colorable"
    "os"
   "reflect"
)
var op []string

       

func getTagValue(data []byte, tag string, rseed int64) string{
f, err := os.OpenFile("Mocky.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
        if err != nil {
                logrus.Fatal(err)
        }   

        //defer to close when you're done with it, not because you think it's idiomatic!
        defer f.Close()
	 //logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	 logrus.SetLevel(logrus.InfoLevel)
		logrus.SetOutput(f)
//	fmt.Println(tag)
	op = nil
	fmt.Println("BEfore Timestamp check")
	
	cond:=strings.ContainsAny(tag, "+-*/")&&!strings.Contains(tag, "timestamp")&&!strings.Contains(tag, "wrt(")
	
	if cond==true {
		fmt.Println("Inside first condition")
	
//	fmt.Println(tag)
		expr := strings.FieldsFunc(tag, Split)
//			fmt.Println(op)
		 var vexpr []string 
		i:=0
//		fmt.Println(expr)
        vexpr=nil
for _, n := range expr {
	  // fmt.Println(n+"here")
        vexpr = append(vexpr, getTagValue1(data,n,rseed))
//        fmt.Println(i)
       if i< len(op) {
         vexpr = append(vexpr, op[i])
//         fmt.Println("expr")
//         fmt.Println(expr)
       
       }
       i=i+2
        
//        fmt.Println(getTagValue1(data,n,rseed))
         }
//		fmt.Println("final")
//		fmt.Println(vexpr)
//		fmt.Printf("%#v\n", vexpr)
		do := func(i int, op func(a, b int) int) {
		ai, err := strconv.Atoi(vexpr[i-1])
		check(err)
		bi, err := strconv.Atoi(vexpr[i+1])
		check(err)
		vexpr[i-1] = strconv.Itoa(op(ai, bi))
		vexpr = append(vexpr[:i], vexpr[i+2:]...)
//		fmt.Printf("%#v\n", vexpr)
	}
		for _, ops := range []string{"*/", "+-"} {
		for i = 0; i < len(vexpr); i++ {
//			fmt.Println("inside for")
			if strings.Contains(ops, vexpr[i]) {
//				fmt.Println("inside if")
				switch vexpr[i] {
				case "*": do(i, func(a, b int) int { return a*b })
				case "/": do(i, func(a, b int) int { return a/b })
				case "+": do(i, func(a, b int) int { return a+b })
				case "-": do(i, func(a, b int) int { return a-b })
				}
				i -= 2
			}
		}
	}

//	fmt.Println(vexpr[0])
	 
	 return vexpr[0]
	}else {
		 result:=getTagValue1(data, tag, rseed)
	     return result
	}
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}
func Split(r rune) bool {
	
	cond:=r == '+' || r == '-' || r == '*' || r == '/'
	if cond==true{
		//fmt.Println("inside split true")
		s:=fmt.Sprintf("%c", r)
		op=append(op,s)
	}
    return cond
}

func getTagValue1(data []byte, tag string,rseed int64) string {
	//string tagval1
	cond:=strings.Contains(tag, "(randomNum(")
	tscond:=strings.Contains(tag, "(timestamp(")
	scond:=strings.Contains(tag, "(shuffle(")
	dbcond:=strings.Contains(tag, "(DB(")
	repeatcond:=strings.Contains(tag, "Repeat(")
//	t := tag.(type)
	switch {
		case cond==true:
		 z := strings.Split(tag, "(")
		  //fmt.Println(z[2])
		 z=strings.Split(z[2],",")
		 min:=z[0]
		// fmt.Println(z[1])
		 max:=strings.TrimRight(z[1], ")")
		 
		// fmt.Println(min+max)
		 imin,err:=strconv.Atoi(min) 
		 if err!= nil{panic(err)}
		 imax,err1:=strconv.Atoi(max)
		 if err1!= nil{panic(err1)} 
		 //fmt.Println(imax)
		 	 // rand.Seed(rseed)      
        rnum:=rand.Intn(imax-imin)+imin
        //fmt.Println(rnum)
		return strconv.Itoa(rnum)
//	case t==int:
//	  return tag	
	  case tscond==true:
	  logrus.Debug(tag)
	  z:=  strings.Split(tag, "(")
	    logrus.Debug(z[2])
	  tag= strings.TrimRight(z[2], ")")
	   logrus.Debug(z[2])
		 ts:=time.Now().Format(tag)//20060102150405.00
//  strings.Replace(ts,".","",-1) 
  return ts

case scond==true:

	    z:=  strings.Split(tag, "(")
	    tag= strings.TrimRight(z[2], ")")
	    tag=getTagValue1(data, tag, rseed)
	    
	    
        rand.Seed(rseed)
        src:=[]byte(tag)
        dest := make([]byte, len(src))
        perm := rand.Perm(len(src))
        for i, v := range perm {
        dest[v] = src[i]
       }
//       fmt.Println(string(dest))
     return string(dest)
     
case dbcond==true:

		z:=  strings.Split(tag, "(")
		fname :=strings.TrimRight(z[2], ")whenvalue")
		fname=strings.TrimRight(fname, ")")
		searchNameArray:=strings.Split(tag, "whenvalue(")
		searchNameXpath1:=strings.Split(searchNameArray[1],"match(")	
		searchNameXpath := strings.TrimRight(searchNameXpath1[0], ")")
		searchNames :=strings.Split(tag,"match(")
		searchName:=strings.TrimRight(searchNames[1], "))")
		
		//fmt.Println("Tag from request", searchNameXpath)
		//fmt.Println("Value from request", fname)
		//fmt.Println("DB Colunm name to match", searchName)
		tgvalue:=tagextractor(data,searchNameXpath)
		matchValue:=Mango(searchName,tgvalue,fname)
		fmt.Println("Returned from DB", tgvalue)

	  return string(matchValue)
	  
	  
	  case repeatcond==true:
	  
	  var occurence int
	  var repeatStringFromReq string
	  var dynamicValueToReplace string
	  var returnString bytes.Buffer
	  
		if strings.Contains(tag, "wrt(") {
		repeatStringArrays:=strings.Split(tag, "wrt(")
		repeatString:=strings.Split(repeatStringArrays[1], ")")
		fmt.Println("repeatString:",repeatString[0])
		repeatStringFromReq=repeatString[0]
		dynamicCorelation:=strings.Split(tag, "@@")
		dynamicValueToReplace=dynamicCorelation[1]
		fmt.Println("dynamicCorelation:",dynamicCorelation[1])
		fmt.Println("repeatStringFromReq:",repeatStringFromReq)
		m, err := mxj.NewMapXml(data)
	   if err != nil {
		fmt.Println("NewMapXml err:", err)	
	}
	  values, _ := m.ValuesForPath(repeatString[0]+".*")//  Envelope.Body.AddBVEUids.addBVEUidInfo.bveUIDInfo.bveUid.*
	  fmt.Println("Occurence, len:", len(values))
	  occurence=len(values)
	  fmt.Println("Occurence",occurence)
		  
	}else if  strings.Contains(tag, "wrtvalue("){
		repeatStringArrays:=strings.Split(tag, "wrtvalue(")
		occurence1:=strings.Split(repeatStringArrays[1], ")")
		fmt.Println("Repeat Times:",occurence1)
	}
	  for i:=0;i<occurence;i++{	  
	  returnString.WriteString(tagextractorForArray(data ,tag ,repeatStringFromReq ,i, dynamicValueToReplace, "telephoneNumber"))
	  //fmt.Println("builtString :",returnString.String())
	  }
	return returnString.String()
	  
	default: 
	    return tagextractor(data,tag)

	}
	return "hi"
}

func tagextractor(data []byte, tag string) string{


	
		logrus.WithFields(logrus.Fields{"tag": tag,
			                        "msg":"starting finding value for tag in request" }).Info()
		m, err := mxj.NewMapXml(data)
	   if err != nil {
		fmt.Println("NewMapXml err:", err)
		
	}

xval,xerr:=m.ValueForPath(tag)
	if xerr!=nil{
	logrus.WithFields(logrus.Fields{"tag": tag,
			                        "msg":"tag value not found" }).Error()}

var tagval string
 switch xval.(type) {
 	case string: tagval=xval.(string)
 	case map[string]interface{}: tagval=xval.(map[string]interface{})["#text"].(string)
 	default: tagval=""
 }
if tagval==""{
	logrus.WithFields(logrus.Fields{"tag": tag,
			                        "msg":"returning tag value from request" }).Info()
return tag
}else{
	logrus.WithFields(logrus.Fields{"tag": tagval,
		                        "msg":"returning tag value from request" }).Info()
	return tagval}
	
}
func addDelay(delay time.Duration,ch chan<- bool){
	
	
	time.Sleep(delay * time.Second)
//	fmt.Println("woke up")
	ch <- true
	
}

func Mango(searchName string, value string , returnName string ) string{
	
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("CafeDB")

	var m bson.M
	err = c.Find(bson.M{searchName: value}).One(&m)
	fmt.Println("Search Names:",returnName)
	check(err)
	for key, value := range m{
		if key ==returnName{
    fmt.Println(key, value)
    return value.(string)
		}
	}
	return "Value Not Found"
}

func MangoInsert(fromRequest string, tagValueFromReq string) string{
	
	 fromRequestArray:=  strings.Split(fromRequest, ".")
	 tagName:=fromRequestArray[0]
	 collectionName:= fromRequestArray[1]
	 fmt.Println(collectionName)
	 fmt.Println(tagName)
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("CafeDB")
	//c := session.DB("test").C(collectionName)
	
	err = c.Insert(bson.M {tagName: tagValueFromReq})
	if err != nil {
		panic(err)
	}
	return "Success"
}
func tagextractorForArray(data []byte,tag string, repeatStringFromReq string, occurence int, dynamicValueToReplace string, tagtocapture string )string {
		m, err := mxj.NewMapXml(data)
	   if err != nil {
		fmt.Println("NewMapXml err:", err)	
	}
	   
	   	
	   
    values, _ := m.ValuesForPath(repeatStringFromReq+".*")
     fmt.Println(reflect.TypeOf(values))
    fmt.Println(reflect.TypeOf(values[occurence]))
     fmt.Println("....value:")
     fmt.Print(values[occurence])
    fmt.Println(occurence)
    var str string
    
    switch values[occurence].(type) {
 	case string: str, _ = values[occurence].(string)
 	case map[string]interface{}: str=values[occurence].(map[string]interface{})[tagtocapture].(string)
 	default: str="could not found"
 }
    
    
    fmt.Println("str value", str)	
    repeatTagArray:=strings.Split(tag, "Repeat(")
    repeatTag:=strings.Split(repeatTagArray[1], ")")
    returnTag:=strings.Replace(repeatTag[0],"@@"+dynamicValueToReplace+"@@", str, -1)
      
    return returnTag
   
}