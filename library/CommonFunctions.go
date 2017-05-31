package main

import (
	"fmt"
	"github.com/clbanning/mxj"
	"math/rand"
    "time"
    "os"
)
      
func tagextractor(data []byte, tag string) string{
	m, err := mxj.NewMapXml(data)
	if err != nil {
		fmt.Println("NewMapXml err:", err)
	}
	xval,xerr:=m.ValueForPath(tag)
	if xerr!=nil{
	fmt.Println("value not found")
	}
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
	ch <- true
}

func getRandomNumber(min int, max int) int{
	rand.Seed(time.Now().UnixNano())
	rnum:=rand.Intn(imax-imin)+imin
    return rnum
}

func getFormattedTimeStamp(format string) string{
	ts:=time.Now().Format(tag)
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
