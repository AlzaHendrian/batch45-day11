package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// root statis untuk mengakses folder public
	e.Static("/public", "public") //public

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	// renderer
	e.Renderer = t

	// routing
	e.GET("/", home)
	e.GET("/contact", contactMe)
	e.GET("/project", myProject)
	e.GET("/project-detail/:id", projectDetail)
	e.POST("/add-project", addProject)

	fmt.Println("localhost: 5000 sucssesfully")
	e.Logger.Fatal(e.Start("localhost: 5000"))
}

func home(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func contactMe(c echo.Context) error {
	return c.Render(http.StatusOK, "contact-me.html", nil)
}

func myProject(c echo.Context) error {
	return c.Render(http.StatusOK, "myProject.html", nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	data := map[string]interface{}{
		"id":      id,
		"title":   "Mobile-App 2019",
		"content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis tempus imperdiet lectus, eu tincidunt leo luctus a. Sed volutpat elit risus, eget laoreet nibh tincidunt nec. Morbi convallis metus ipsum. Vivamus elementum iaculis diam, ac laoreet elit mattis in. Aliquam vestibulum massa ante, eu luctus neque dignissim vel. Nunc facilisis erat eu eros vestibulum faucibus. Cras malesuada, eros eu varius lobortis, nulla est porta Mauris iaculis nisl in purus tempus ultrices. Suspendisse eu sem pharetra, mattis felis nec, eleifend metus. Nam bibendum lorem ex. Fusce eleifend gravida tincidunt. Praesent justo sapien, hendrerit eu massa in, congue aliquam ligula. Sed at orci arcu. Maecenas fermentum blandit orci, vel vehicula metus feugiat a.",
	}
	return c.Render(http.StatusOK, "projectDetail.html", data)
}

func addProject(c echo.Context) error {
	title := c.FormValue("project-name")
	sDate := c.FormValue("start-date")
	eDate := c.FormValue("end-date")
	desc := c.FormValue("description")
	tech1 := c.FormValue("tech_1")
	tech2 := c.FormValue("tech_2")
	tech3 := c.FormValue("tech_3")
	tech4 := c.FormValue("tech_4")

	println("name project: " + title)
	println("star date: " + sDate)
	println("end date: " + eDate)
	println("description: " + desc)
	println("technology: " + tech1)
	println("technology: " + tech2)
	println("technology: " + tech3)
	println("technology: " + tech4)

	return c.Redirect(http.StatusMovedPermanently, "/project")
}
