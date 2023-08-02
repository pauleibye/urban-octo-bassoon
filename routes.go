package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"sort"
	"strconv"

	"github.com/labstack/echo/v4"
)

type series struct {
	Id     string
	Render template.HTML
	Series []int
}

func NewSeries(id string, data []int) series {
	return series{
		Id:     id,
		Render: render(data),
		Series: data,
	}
}

func (s *series) setSeries(newSeries []int) {
	s.Series = newSeries
	s.Render = render(s.Series)
}

func fromFormParams(id string, params url.Values) series {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	data := make([]int, 0, len(params))
	fmt.Printf("paultest data len: %v", len(data))
	for _, k := range keys {
		v, _ := strconv.Atoi(params.Get(k))
		data = append(data, v)
	}

	return NewSeries(id, data)
}

var series1 = series{
	Id:     "1",
	Render: render([]int{1, 2, 3, 4, 5, 6, 7}),
	Series: []int{1, 2, 3, 4, 5, 6, 7},
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", series1)
}

func getSeries(c echo.Context) error {
	id := c.Param("id")
	if id == "1" {
		return c.JSON(http.StatusOK, series1)
	}
	return c.JSON(http.StatusNotFound, "series not found")
}

// type putSeriesRequest struct {
// 	Series []int `form:"putSeries"`
// }

func putSeries(c echo.Context) error {
	// var series map[string]string
	// err := c.Bind(&series)
	values, err := c.FormParams()
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	fmt.Printf("paultest values: %v\n", values)

	// if len(series.Series) == 0 {
	// 	return c.String(http.StatusBadRequest, "bad request")
	// }

	id := c.Param("id")
	if id == "1" {
		series1 = fromFormParams(id, values)
		return c.Render(http.StatusOK, "index.html", series1)
	}
	return c.JSON(http.StatusNotFound, "series not found")
}
