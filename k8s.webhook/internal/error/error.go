package error

import (
	"errors"
	"reflect"
	"strings"
	"time"
)

//Error is a custom structure for API error
type Error struct{
	Message string `json:"message"`
	Details string `json:"details"`
	Timestamp string `json:"timestamp"`
	Path string `json:"path"`

}

//FormatError is a setter for all Error attributes.
func(e *Error)FormatError(message string,details string,path string){
	e.Message=message
	e.Details=details
	t := time.Now()
	e.Timestamp=t.String()
	e.Path=path
	//return e
}


//Unmarshal function check if a required parameter is missing
func (e *Error) Unmarshal(data interface{}) (string,string,error) {
	var err []string
	fields := reflect.ValueOf(data).Elem()
	//Parse attribute to find if when they have mandatory tag it set to true
	for i := 0; i < fields.NumField(); i++ {
		mandatoryTag := fields.Type().Field(i).Tag.Get("mandatory")
		if strings.Contains(mandatoryTag, "true") && fields.Field(i).IsZero() {
			err=append(err,fields.Type().Field(i).Name)
		}
	}
	//Return message, details and error
	if(len(err)>0){
		return "Bad Request","Missing parameter(s): "+strings.Join(err, ","),errors.New("Bad Request")
	}else{
	return "","",nil
	}
}