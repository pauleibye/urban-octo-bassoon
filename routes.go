package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pauleibye/urban-octo-bassoon/internal/usecase"
)

// func fromFormParams(id string, params url.Values) series {
// 	keys := make([]string, 0, len(params))
// 	for k := range params {
// 		keys = append(keys, k)
// 	}
// 	sort.Strings(keys)

// 	data := make([]int, 0, len(params))
// 	fmt.Printf("paultest data len: %v", len(data))
// 	for _, k := range keys {
// 		v, _ := strconv.Atoi(params.Get(k))
// 		data = append(data, v)
// 	}

//		return NewSeries(id, data)
//	}
type SeriesResponse struct {
	Id     string        `json:"id"`
	Render template.HTML `json:"render"`
}

type controller struct {
	uc usecase.Usecase
}

func (c controller) Index(ctx echo.Context) error {
	series, render, err := c.uc.GetSeries(1)
	if err != nil {
		fmt.Println("paultest err: ", err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.Render(http.StatusOK, "index.html", SeriesResponse{Id: series.Name, Render: render})
}

// type putSeriesRequest struct {
// 	Series []int `form:"putSeries"`
// }

// func putSeries(c echo.Context) error {
// 	// var series map[string]string
// 	// err := c.Bind(&series)
// 	values, err := c.FormParams()
// 	if err != nil {
// 		return c.String(http.StatusBadRequest, "bad request")
// 	}

// 	fmt.Printf("paultest values: %v\n", values)

// 	// if len(series.Series) == 0 {
// 	// 	return c.String(http.StatusBadRequest, "bad request")
// 	// }

// 	id := c.Param("id")
// 	if id == "1" {
// 		series1 = fromFormParams(id, values)
// 		return c.Render(http.StatusOK, "index.html", series1)
// 	}
// 	return c.JSON(http.StatusNotFound, "series not found")
// }
