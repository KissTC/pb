package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html/v2"
)

type Post struct {
	Id          string
	Title       string `json:"title"`
	PublishDate string `json:"publish_date"`
	Content     string `json:"content"`
}

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views: engine,
		},
	)

	app.Static("/", "public")
	authMiddleware := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"Kiss": "thediablo",
		},
	})

	// render the home page
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index template
		var posts []Post

		files, err := os.ReadDir("./posts")
		if err != nil {
			return c.JSON("error at rendering template")
		}

		for _, file := range files {
			content, err := os.ReadFile("./posts/" + file.Name())
			if err != nil {
				log.Println("Error reading file:", file.Name(), err)
				continue
			}

			var post Post
			if err := json.Unmarshal(content, &post); err != nil {
				log.Println("Error unmarshalling file:", file.Name(), err)
				continue
			}
			post.Id = strings.TrimSuffix(file.Name(), ".json")

			posts = append(posts, post)

		}

		return c.Render("index", fiber.Map{
			"Data": posts,
		})
	})

	// render articles
	app.Get("/post/:id", func(c *fiber.Ctx) error {
		// TODO: Check if id is provided
		if c.Params("id") == "" {
			return c.SendString("No post id provided")
		}
		// TODO: Check if json file exist
		file := c.Params("id") + ".json"
		content, err := os.ReadFile("./posts/" + file)
		if err != nil {
			log.Println("Error al cargar post:", file, err)
		}

		var post Post
		err = json.Unmarshal(content, &post)
		if err != nil {
			log.Println("Error unmarshalling file:", file, err)
		}

		return c.Render("post", fiber.Map{
			"Data": post,
		})
	})

	// render the admin page where user can create, edit or delete new pages
	app.Get("/admin", authMiddleware, func(c *fiber.Ctx) error {
		// TODO: Check if the user is authenticated
		// TODO: Render the page showing the list of articles
		// TODO: In the page, redirect to /delete /new /edit with the id
		return c.JSON("asdasd")
	})

	// render the new page so admin can publish
	app.Get("/new", authMiddleware, func(c *fiber.Ctx) error {
		// TODO: Check if the user is authenticated. Perhaps create a middleware for that
		// TODO: Render the page with a form to submit a new article
		// TODO: In the form send a request to POST /publish or /new
		return c.JSON("asdasd")
	})

	// render the edit page so admin can edit a poist
	app.Get("/edit/:id", authMiddleware, func(c *fiber.Ctx) error {
		// TODO: Check if the user is authenticated. Perhaps create a middleware for that
		// TODO: Render the page with a form to submit a new article
		// TODO: In the form send a request to POST /publish or /new
		return c.JSON("asdasd")
	})

	// publish new posts
	app.Post("/publish", func(c *fiber.Ctx) error {
		// TODO: Create the json file
		// TODO: Return status created and redirect to /post/id
		return c.JSON("asdasd")
	})

	// edit selected article
	app.Put("/publish", func(c *fiber.Ctx) error {
		// TODO: Create the json file
		// TODO: Return status created and redirect to /post/id
		return c.JSON("asdasd")
	})

	app.Listen(":3000")
}
