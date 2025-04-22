package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string
	Email string
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

var users []User

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	// GET ROUTE FIRST VISITING PAGE
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "users.html", map[string]interface{}{
			"name":  "Dolly!",
			"email": "dolly@example.com",
			"users": initialUsers(),
		})
	})

	e.POST("/users", func(c echo.Context) error {

		tmpl := template.Must(template.ParseFiles("user_row.html"))

		name := c.FormValue("name")
		email := c.FormValue("email")

		user := User{
			Name: name, Email: email,
		}

		users = append(users, user)
		// fmt.Println("Users:", users)

		err := tmpl.Execute(c.Response().Writer, user)
		if err != nil {
			panic(err)
		}

		return nil
	})

	e.DELETE("/users-delete/:name", func(c echo.Context) error {
		tmpl := template.Must(template.ParseFiles("user_row_delete.html"))

		name := c.Param("name")

		// fmt.Println("Name to delete:", name)

		users = removeObject(users, name)

		user := User{
			Name: "", Email: "",
		}
		err := tmpl.Execute(c.Response().Writer, user)
		if err != nil {
			panic(err)
		}

		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func removeObject(slice []User, value string) []User {
	var newSlice []User
	for _, user := range slice {
		if user.Name != value {
			newSlice = append(newSlice, user)
		}
	}
	return newSlice
}

func initialUsers() []User {
	users = []User{
		{Name: "John Doe", Email: "johndoe@example.com"},
		{Name: "Alice Smith", Email: "alicesmith@example.com"},
	}

	return users
}
