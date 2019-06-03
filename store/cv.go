package store

import (
	"honnef.co/go/js/xhr"
)

const url string = "http://localhost:8090"

var cv string

func FetchCV() error {

	data, err := xhr.Send("GET", url+"/static/cv.json", nil)
	if err != nil {
		return err
	}

	cv = string(data)

	return nil
}

func GetCV() string {
	return cv
}
