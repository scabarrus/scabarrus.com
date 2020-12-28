package library

import (
	"net/http"
	//"io/ioutil"
	"crypto/tls"
	"fmt"
	"strings"
)

type CallAPI struct{
	URL string
	Method string
	Header map[string]string
	Payload *strings.Reader
	SSLCheck bool

}

func (c CallAPI) CallRest()(*http.Response,error){
	if c.SSLCheck == false {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	fmt.Println("URL :",c.URL)
	fmt.Println("METHOD:",c.Method)
	fmt.Println("HEADER :",c.Header)
	fmt.Println("PAYLOAD :",c.Payload)
	request, _ := http.NewRequest(c.Method, c.URL,c.Payload)
	for header,value:= range c.Header {
		request.Header.Add(header, value)
	}
	response, err := http.DefaultClient.Do(request)
	return response,err
}