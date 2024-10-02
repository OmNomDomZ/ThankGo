package main

import (
	"encoding/json"
	"fmt"
)

// начало решения
//type Rating string
//
//func (r *Rating) UnmarshalJSON(data []byte) error {
//	var v string
//	err := json.Unmarshal(data, &v)
//	if err != nil {
//		return err
//	}
//
//	rn := []rune(v)
//	for i := range rn {
//		if rn[i] == rune('★') {
//			*r += "★"
//		} else {
//			*r += "☆"
//		}
//	}
//
//	return nil
//}

// Movie описывает фильм
type Movie struct {
	Name       string
	ReleasedAt int     `json:"released_at"`
	Genres     []Genre `json:"tags"`
}

// Genre описывает жанр фильма
type Genre string

func (g *Genre) UnmarshalJSON(data []byte) error {
	var obj map[string]string
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return err
	}

	*g = Genre(obj["name"])
	return nil
}

// конец решения

func main() {
	const src = `{
		"name": "Interstellar",
		"released_at": 2014,
		"director": "Christopher Nolan",
		"tags": [
			{ "name": "Adventure" },
			{ "name": "Drama" },
			{ "name": "Science Fiction" }
		],
		"duration": "2h49m",
		"rating": "★★★★★"
	}`

	var m Movie
	err := json.Unmarshal([]byte(src), &m)
	fmt.Println(err)
	// nil
	fmt.Println(m)
	// {Interstellar 2014 [Adventure Drama Science Fiction]}
}
