package main

import (
	"log"

	"github.com/adwait-rao/adwaitrao-dev/internal/builder"
)

func main() {
	log.Println("Starting site compilation...")

	err := builder.BuildSite("content/blog", "public", "templates")
	if err != nil {
		log.Fatalf("Build failed %v", err)
	}

	log.Println("Build complete! Check yout /public folder.")
}
