package main

import (
    "github.com/mrGlasses/APIHandler/bex"

	"github.com/gin-gonic/gin"
)



func main() {
    // Create GIN
    router := gin.Default()

    // Call MS

    // Call BEX
    router.GET("/bex", bex.CallBanana)

    // Run GIN
    router.Run(":443")
}