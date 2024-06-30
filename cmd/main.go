package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-fuego/fuego"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Topic struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	OwnerToken string `json:"ownerToken"`
	Items      []Item `json:"items"`
}

type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       *string `json:"image"`
}

func main() {
	s := fuego.NewServer(fuego.WithPort(8080))

	fs := http.FileServer(http.Dir("./client/dist"))

	fuego.Handle(s, "/", fs)

	api := fuego.Group(s, "/api")
	userRoutes := fuego.Group(api, "/user")
	topicRoutes := fuego.Group(api, "/topic")

	fuego.Get(api, "/hello", func(c fuego.ContextNoBody) (string, error) {
		return "Hello, World!", nil
	})

	fuego.Get(userRoutes, "/", func(c fuego.ContextNoBody) (User, error) {
		return User{
			ID:    "test-1",
			Name:  "John Doe",
			Email: "jd@test.com",
		}, nil
	})

	fuego.Get(topicRoutes, "/all", func(c fuego.ContextNoBody) ([]Topic, error) {
		topicByte := []byte(`[{"id":"1","title":"Best Neovim Colorscheme","ownerToken":"987","items":[{"id":"21","name":"Gruvbox","image":"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQbMpzqcq3rKrcDjxeV16U5az9WgPg86S-v6w&s","description":"Designed as a bright theme with pastel 'retro groove' colors and light/dark mode switching in the way of solarized."},{"id":"22","name":"Catpuccin","image":"https://raw.githubusercontent.com/catppuccin/catppuccin/main/assets/logos/exports/1544x1544_circle.png","description":"Soothing pastel theme for the high-spirited!"},{"id":"23","name":"Kanagawa","image":"https://w7.pngwing.com/pngs/506/832/png-transparent-emoji-the-great-wave-off-kanagawa-sticker-sun-on-the-lake-emoji-logo-car-smiley-thumbnail.png","description":"NeoVim dark colorscheme inspired by the colors of the famous painting by Katsushika Hokusai."},{"id":"24","name":"Tokyo Night","image":null},{"id":"25","name":"Nord","image":null}]},{"id":"2","title":"Best front-end framework","ownerToken":"123","items":[{"id":"1","name":"React","image":"https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/2300px-React-icon.svg.png","description":"The library for web and native user interfaces"},{"id":"2","name":"Next","image":"https://static-00.iconduck.com/assets.00/nextjs-icon-512x512-y563b8iq.png","description":"Used by some of the world's largest companies, Next.js enables you to create high-quality web applications with the power of React components."},{"id":"3","name":"Vue","image":"https://w7.pngwing.com/pngs/854/555/png-transparent-vue-js-hd-logo-thumbnail.png","description":"An approachable, performant and versatile framework for building web user interfaces."},{"id":"4","name":"Svelte","image":null},{"id":"5","name":"Solid","image":null},{"id":"6","name":"Nuxt","image":null},{"id":"7","name":"Laravel","image":null},{"id":"8","name":"Ruby On Rails","image":null},{"id":"9","name":"Astro","image":null}]}]`)

		topic := []Topic{}

		if err := json.Unmarshal(topicByte, &topic); err != nil {
			panic(err)
		}

		return topic, nil
	})

	s.Run()
}
