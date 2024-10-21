package main

import (
	"fmt"
	"server/controllers"
	"server/models"
	"server/routes"
)

func main() {
    // Initialise the videos array to keep track of all videos in memory
    controllers.Videos = models.Videos{
        {
            Title:       "Example 1 Minute Video",
            Description: "An example of a 1 minute video",
            Length:      60,
        },
    }

    // Start the server
    fmt.Println("Starting server on the port 8000")
    routes.HandleRequests()
}