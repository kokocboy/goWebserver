package cboy

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)
//favicon.ico
func GetHtml(url string)  string{
	if strings.HasSuffix(url,".ico")==false{
		url+=".html"
	}
	content, err := ioutil.ReadFile("resourse"+url)
	if err != nil {
		log.Printf("read file failed, err:%v", err)
		content, err = ioutil.ReadFile("resourse"+"/404"+".html")
		return string(content)
	}
	if strings.HasSuffix(url,"ioc"){
		return string(content)
	}
	str:=string(content)
	slice:=strings.SplitAfter(str,"rand")
	if len(slice)>1{
		str=""
		age:=getAge(url)
		for i:=0;i<len(slice)-1;i++{
			str+=slice[i]+strconv.Itoa(age)
		}
		str+=slice[len(slice)-1]
	}
	return str
}