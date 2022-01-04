package cboy

import (
	"io/ioutil"
	"log"
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
	return string(content)
}