package Tests

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// WeatherService предсказывает погоду.
type WeatherService struct{}

// Forecast сообщает ожидаемую дневную температуру на завтра.
func (ws *WeatherService) Forecast() int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(31)
	sign := rand.Intn(2)
	if sign == 1 {
		value = -value
	}
	return value
}

// Weather выдает текстовый прогноз погоды.
type Weather struct {
	service Service
}

// Forecast сообщает текстовый прогноз погоды на завтра.
func (w Weather) Forecast() string {
	deg := w.service.Forecast()
	switch {
	case deg < 10:
		return "холодно"
	case deg >= 10 && deg < 15:
		return "прохладно"
	case deg >= 15 && deg < 20:
		return "идеально"
	case deg >= 20:
		return "жарко"
	}
	return "инопланетно"
}

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

type Service interface {
	Forecast() int
}

type WeatherServiceMock struct {
	i int
}

func (wsm *WeatherServiceMock) Forecast() int {
	res := tests[wsm.i].deg
	wsm.i += 1
	return res
}

func TestForecast(t *testing.T) {
	service := &WeatherServiceMock{i: 0}
	weather := Weather{service}
	for _, test := range tests {
		name := fmt.Sprintf("%v", test.deg)
		t.Run(name, func(t *testing.T) {
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}
