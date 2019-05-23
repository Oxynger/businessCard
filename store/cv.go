package store

const url string = "localhost:8080"

var cv string

func FetchCV() error {
	// data, err := xhr.Send("GET", url+"/cv", nil)
	// if err != nil {
	// 	return err
	// }
	// cv = string(data)
	cv = `name: "Oxynger",
skills: [],
hobbies: [
  "AuthoChess",
  "Anime"
]
`
	return nil
}

func GetCV() string {
	return cv
}
