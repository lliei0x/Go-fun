package main

import (
	"leeif.me/go_exercise/github-trending/engine"
	"leeif.me/go_exercise/github-trending/parse/github"
)

func main() {

	var simplerTest engine.Trending

	simplerTest.Run(
		engine.RequestForGithub{
			URL:       "https://github.com/trending",
			ParseFunc: github.ParseForGithub,
		},
	)
	simplerTest.Run(
		engine.RequestForGithub{
			URL:       "https://github.com/trending/developers",
			ParseFunc: github.ParseForDevelopers,
		},
	)

}
