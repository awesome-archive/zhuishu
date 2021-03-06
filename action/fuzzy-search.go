package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1"
)

type ResFuzzySearch struct {
	Ok    bool
	Books []struct {
		Id          string `json:"_id"`
		Title       string
		Cat         string
		Author      string
		Cover       string
		ShortIntro  string
		LastChapter string
	}
}

func ActionPostFuzzySearch(c *gin.Context) {
	query := c.PostForm("query")
	strStart := c.DefaultPostForm("start", "0")
	start, _ := strconv.Atoi(strStart)
	strLimit := c.DefaultPostForm("limit", "10")
	limit, _ := strconv.Atoi(strLimit)

	ret := fuzzySearch(query, start, limit)
	c.JSON(http.StatusOK, ret)
}

func fuzzySearch(query string, start int, limit int) gin.H {
	client := http.DefaultClient

	query = url.QueryEscape(query)
	url := fmt.Sprintf("http://api.zhuishushenqi.com/book/fuzzy-search?query=%s&start=%d&limit=%d", query, start, limit)

	resp, e := client.Get(url)
	if e != nil {
		log.Panic(e)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var ret ResFuzzySearch
	json.Unmarshal(body, &ret)

	return gin.H{
		"status": 0,
		"msg":    "ok",
		"data": gin.H{
			"books": ret.Books,
		},
	}

}
