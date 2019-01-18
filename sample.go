package main

import (
	"fmt"
	"github.com/cinic0101/go-esconn/api/template"
	"github.com/cinic0101/go-esconn/client"
	"github.com/cinic0101/go-esconn/config"
)

func main() {
	c, _ := client.New(&config.Config{ URL: "http://localhost:9200"})
	resp, _ := c.SearchTemplate("index-2019.01.17",
		template.New().
		WithFromSize(0, 100).
		MatchQuery("list.value", "value1 value2"))

	fmt.Println(resp)
}
