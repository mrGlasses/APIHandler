package main

import (
	"github.com/mrGlasses/APIHandler/bex"

    "log"

	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)



func main() {
    // Create GIN
    router := gin.Default()

    // Call MS

    // Call BEX
    router.GET("/bex", bex.CallBanana)

    // Run GIN
    // router.Run(":443")

    // Run with Let's Encrypt
    log.Fatal(autotls.Run(router, "bcookietech.duckdns.org"))
}