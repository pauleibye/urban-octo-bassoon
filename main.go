package main

import (
	"html/template"
	"io"
	"log"
	"math/rand"
	"os"

	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/pauleibye/urban-octo-bassoon/internal/usecase"
)

// type Film struct {
// 	Title    string
// 	Director string
// }

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func convertToLineData(series []int) []opts.LineData {
	items := make([]opts.LineData, 0)
	for _, v := range series {
		items = append(items, opts.LineData{Value: v})
	}
	return items
}

func generateLineItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

var schema = `
CREATE TABLE if not exists series(
   id INT GENERATED ALWAYS AS IDENTITY,
   name VARCHAR(255),
   PRIMARY KEY(id)
);

CREATE TABLE if not exists point(
	id INT GENERATED ALWAYS AS IDENTITY,
	series_id INT,
	time TIMESTAMP,
	value DECIMAL,
	primary key(id),
	constraint fk_series
		foreign key(series_id) 
			references series(id)
);

INSERT INTO series (name) VALUES ('paultest1');
INSERT INTO point (series_id, time, value) VALUES (1, '2021-01-01', 2.0);
`

func main() {
	db, err := sqlx.Connect("postgres", "user=uob password=password dbname=postgres sslmode=disable host=localhost port=6543")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
	// 	fmt.Printf("paultest reqBody: %v\n", string(reqBody))
	// 	fmt.Printf("paultest resBody: %v\n", string(resBody))
	// }))

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t

	controller := controller{
		uc: usecase.Usecase{DB: db},
	}

	e.GET("/", controller.Index)
	// e.GET("/series/:id", getSeries)
	// e.PUT("/series/:id", putSeries)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Debug(e.Start(":" + httpPort))
}
