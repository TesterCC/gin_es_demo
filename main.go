package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
)

func Query(c *gin.Context) {
	c.HTML(200, "query.html", nil)
}

func ESQuery(keyword string) *esapi.Response {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Get client error: %s", err)
	}

	// query condition
	query := `{"query":{"match":{"Name":"` + keyword + `"}},"size":3}`

	fmt.Printf("query: %v\n", query)

	var b strings.Builder
	b.WriteString(query)
	read := strings.NewReader(b.String())

	// query request
	sr := &esapi.SearchRequest{
		Index: []string{"golang"},
		Body:  read,
	}

	// execute query
	r, _ := sr.Do(context.Background(), client)
	return r
}

func DoQuery(c *gin.Context) {
	keyword := c.PostForm("keyword")
	r := ESQuery((keyword))

	c.HTML(200, "result.html", gin.H{
		"result": r,
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/query", Query)
	r.POST("/query", DoQuery)

	// http://127.0.0.1:7777/query
	r.Run(":7777")
}
