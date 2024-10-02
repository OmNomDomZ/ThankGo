package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"strings"
//	"time"
//)
//
//// начало решения
//
//// Duration описывает продолжительность фильма
//type Duration time.Duration
//
//func (d Duration) MarshalJSON() ([]byte, error) {
//	dur := time.Duration(d)
//	hours := int(dur.Hours())
//	minutes := int(dur.Minutes()) % 60
//	var durationString string
//	if minutes == 0 {
//		durationString = fmt.Sprintf("\"%dh\"", hours)
//	} else if hours == 0 {
//		durationString = fmt.Sprintf("\"%dm\"", minutes)
//	} else {
//		durationString = fmt.Sprintf("\"%dh%dm\"", hours, minutes)
//	}
//
//	return []byte(durationString), nil
//}
//
//// Rating описывает рейтинг фильма
//type Rating int
//
//func (r Rating) MarshalJSON() ([]byte, error) {
//	b, err := json.Marshal(strings.Repeat("★", int(r)) + strings.Repeat("☆", int(5-r)))
//	if err != nil {
//		return nil, err
//	}
//	return b, nil
//}
//
//// Movie описывает фильм
//type Movie struct {
//	Title    string
//	Year     int
//	Director string
//	Genres   []string
//	Duration Duration
//	Rating   Rating
//}
//
//// MarshalMovies кодирует фильмы в JSON.
////   - если indent = 0 - использует json.Marshal
////   - если indent > 0 - использует json.MarshalIndent
////     с отступом в указанное количество пробелов.
//func MarshalMovies(indent int, movies ...Movie) (string, error) {
//	if indent == 0 {
//		b, err := json.Marshal(movies)
//		if err != nil {
//			return "", err
//		}
//		return string(b), nil
//	} else {
//		b, err := json.MarshalIndent(movies, "", strings.Repeat(" ", indent))
//		if err != nil {
//			return "", err
//		}
//		return string(b), nil
//	}
//}
//
//// конец решения
//
//func main() {
//	m1 := Movie{
//		Title:    "Interstellar",
//		Year:     2014,
//		Director: "Christopher Nolan",
//		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
//		Duration: Duration(2*time.Hour + 49*time.Minute),
//		Rating:   5,
//	}
//	m2 := Movie{
//		Title:    "Sully",
//		Year:     2016,
//		Director: "Clint Eastwood",
//		Genres:   []string{"Drama", "History"},
//		Duration: Duration(time.Hour + 36*time.Minute),
//		Rating:   4,
//	}
//
//	s, err := MarshalMovies(4, m1, m2)
//	fmt.Println(err)
//	// nil
//	fmt.Println(s)
//	/*
//		[
//		    {
//		        "Title": "Interstellar",
//		        "Year": 2014,
//		        "Director": "Christopher Nolan",
//		        "Genres": [
//		            "Adventure",
//		            "Drama",
//		            "Science Fiction"
//		        ],
//		        "Duration": "2h49m",
//		        "Rating": "★★★★★"
//		    },
//		    {
//		        "Title": "Sully",
//		        "Year": 2016,
//		        "Director": "Clint Eastwood",
//		        "Genres": [
//		            "Drama",
//		            "History"
//		        ],
//		        "Duration": "1h36m",
//		        "Rating": "★★★★☆"
//		    }
//		]
//	*/
//}
