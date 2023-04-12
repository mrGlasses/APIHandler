package bex

import (
	// "encoding/json"
	"encoding/json"
	"fmt"
	"net/http"

	// "strings"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

type goodAnswer struct {
	Value string `json:"value"`
}

type badAnswer struct {
	Err string `json:"error"`
}

func CallBanana (c *gin.Context) {
	coinB := c.DefaultQuery("base", "USD")
	coinE := c.DefaultQuery("exch", "BRL")
	value := c.DefaultQuery("valu", "1")

	client := resty.New()

    var gAsw goodAnswer
    var bAsw badAnswer

    resp, err := client.R().Get(fmt.Sprintf("http://localhost:25243/banana?base=%s&exch=%s&valu=%s", coinB, coinE, value))

    if err != nil {
        c.JSON(http.StatusServiceUnavailable, err.Error())

	} else if resp.StatusCode() > 299 {		
		json.Unmarshal(resp.Body(), &bAsw)
		
		c.JSON(resp.StatusCode(), bAsw)
    } else {
		json.Unmarshal(resp.Body(), &gAsw)
		
		c.JSON(resp.StatusCode(), gAsw)
    }
}