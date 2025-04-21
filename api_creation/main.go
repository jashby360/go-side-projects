// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"io"
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// type TemplateRenderer struct {
// 	templates *template.Template
// }

// // Render renders a template document
// func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

// 	// Add global methods if data is a map
// 	if viewContext, isMap := data.(map[string]interface{}); isMap {
// 		viewContext["reverse"] = c.Echo().Reverse
// 	}

// 	return t.templates.ExecuteTemplate(w, name, data)
// }

// func main() {
// 	e := echo.New()
// 	renderer := &TemplateRenderer{
// 		templates: template.Must(template.ParseGlob("templates/*.html")),
// 	}
// 	e.Renderer = renderer

// 	e.GET("/", func(c echo.Context) error {
// 		users := GetUsers()

// 		return c.Render(http.StatusOK, "test.html", users)
// 		// return c.String(http.StatusOK, "Hello, World!")
// 	})

// 	e.Logger.Fatal(e.Start(":1323"))
// 	fmt.Println("Server is running on port 1323")
// }

// type User struct {
// 	Name  string
// 	Email string
// }

// func GetUsers() []User {
// 	return []User{
// 		{Name: "John Doe", Email: "johndoe@example.com"},
// 		{Name: "Alice Smith", Email: "alicesmith@example.com"},
// 	}
// }