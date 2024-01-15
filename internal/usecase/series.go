package usecase

import (
	"bytes"
	"fmt"
	"html/template"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/pauleibye/urban-octo-bassoon/internal/db"
	"github.com/pauleibye/urban-octo-bassoon/internal/entity"
	"github.com/pauleibye/urban-octo-bassoon/pkg/util"
)

func (u *Usecase) GetSeries(seriesId int) (*entity.Series, template.HTML, error) {
	fmt.Println("paultest GetSeries")
	series, err := db.GetSeries(u.DB, seriesId)
	if err != nil {
		return nil, template.HTML(""), err
	}

	points, err := db.GetPointsOfSeries(u.DB, seriesId)
	if err != nil {
		return nil, template.HTML(""), err
	}

	bar := charts.NewLine()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: series.Name,
	}),
		charts.WithColorsOpts(opts.Colors{"red", "green"}),
		charts.WithInitializationOpts(
			opts.Initialization{
				Theme:           "dark",
				BackgroundColor: "transparent",
			},
		),
		charts.WithLegendOpts(
			opts.Legend{
				Show: true,
				Top:  "20%",
			},
		),
		charts.WithAnimation(),
		charts.WithTooltipOpts(
			opts.Tooltip{
				Show: true,
			},
		),
	)

	bar.SetXAxis(
		util.Map[entity.Point, time.Time](points, func(p entity.Point) time.Time {
			return p.Time
		}),
	).
		AddSeries(
			series.Name,
			util.Map[entity.Point, opts.LineData](points, func(p entity.Point) opts.LineData {
				return opts.LineData{Value: p.Value}
			}),
		).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	var buf bytes.Buffer
	err = bar.Render(&buf)
	if err != nil {
		return nil, template.HTML(""), err
	}

	return &series, template.HTML(buf.String()), nil

}
