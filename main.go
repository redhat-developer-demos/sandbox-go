package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// activity is a Sandbox activity.
type activity struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// list of activities.
var activities = []activity{
	{Title: "Event-Driven Architecture using Apache Kafka (with Red Hat OpenShift Streams for Apache Kafka)", URL: "https://developers.redhat.com/node/226211"},
	{Title: "Deploy a Java Application on Kubernetes in minutes", URL: "https://developers.redhat.com/node/226211"},
	{Title: "Learn Kubernetes", URL: "https://developers.redhat.com/node/226211"},
}

// list as JSON
func getActivities(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, activities)
}

func main() {
	router := gin.Default()
	router.GET("/", getActivities)

	router.Run("localhost:8080")
}
