package usecase

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/pauleibye/urban-octo-bassoon/internal/db"
	"github.com/pauleibye/urban-octo-bassoon/internal/entity"
	"github.com/pauleibye/urban-octo-bassoon/pkg/util"
)

func (u *Usecase) GetChart(chartId int) (template.HTML, error) {
	fmt.Println("paultest GetSeries")

	series, err := db.GetSeriesOfChart(u.DB, chartId)
	if err != nil {
		return template.HTML(""), err
	}

	pointsOfSeries := make(map[entity.Series][]entity.Point)
	for _, s := range series {
		points, err := db.GetPointsOfSeries(u.DB, s.Id)
		if err != nil {
			return template.HTML(""), err
		}
		pointsOfSeries[s] = points
	}

	plot := charts.NewLine()
	plot.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: strconv.Itoa(chartId),
		}),
		charts.WithInitializationOpts(
			opts.Initialization{
				Theme:           "dark",
				BackgroundColor: "transparent",
			},
		),
		charts.WithLegendOpts(
			opts.Legend{
				Show: true,
			},
		),
		charts.WithTooltipOpts(
			opts.Tooltip{
				Show:      true,
				Enterable: true,
			},
		),
		charts.WithXAxisOpts(
			opts.XAxis{
				Type: "time",
			},
		),
		charts.WithAnimation(),
		charts.WithDataZoomOpts(
			opts.DataZoom{
				Type:  "slider",
				Start: 0,
				End:   10,
			},
		),
	)

	for series, points := range pointsOfSeries {
		plot.AddSeries(
			series.Name,
			util.Map[entity.Point, opts.LineData](points, func(p entity.Point) opts.LineData {
				return opts.LineData{Name: p.Time.Format(time.RFC3339), Value: []interface{}{p.Time, p.Value}}
			}),
			charts.WithLineChartOpts(
				opts.LineChart{
					Smooth:     false,
					ShowSymbol: true,
					Symbol:     "circle",
					SymbolSize: 10,
				},
			),
		)
	}

	var buf bytes.Buffer
	err = plot.Render(&buf)
	if err != nil {
		return template.HTML(""), err
	}

	return template.HTML(buf.String()), nil

}
