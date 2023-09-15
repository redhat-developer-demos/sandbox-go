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
	{Title: "More activities)", URL: "https://developers.redhat.com/developer-sandbox/activities"},
	{Title: "Create an OpenShift Serverless function", URL: "https://developers.redhat.com/developer-sandbox/activities/create-openshift-serverless-function"},
	{Title: "Deploy a Java Application on Kubernetes in minutes", URL: "https://developers.redhat.com/node/226211"},
	{Title: "Learn Kubernetes", URL: "https://developers.redhat.com/developer-sandbox/activities/learn-kubernetes-using-red-hat-developer-sandbox-openshift"},
}

// list as JSON
func getActivities(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, activities)
}

func main() {
	router := gin.Default()
	router.GET("/", getActivities)

	router.Run("0.0.0.0:8080")
}
