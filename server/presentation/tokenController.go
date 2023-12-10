package presentation

import (
	"net/http"
	"oauth2_api/application"
	"oauth2_api/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Introspect(c *gin.Context) {
	id := c.Param("id")
	readerAny, ok := c.Get("tokenReader")
	if !ok {
		panic("Not wired up.")
	}
	reader, ok := readerAny.(application.TokenReader)
	if !ok {
		panic("Wiring is bad!")
	}
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id is not valid."})
	} else {
		c.JSON(http.StatusOK, application.MakeIntrospection(reader, domain.UserId(int(uintId))))
	}
}

/* func Introduce(c *gin.Context) {
	saverAny := interface adapters.StoreToken
	// if !ok {
		// panic("Not wired up.")
	// }
	saver := saverAny.(application.CreateToken)
	// if !okF {
		// panic("Wrong wiring type")
	// }
	tokenBound := new(domain.TokenData)
	err := c.Bind(tokenBound)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body is not valid."})
	} else {
		c.JSON(http.StatusOK, saver(*tokenBound))
	}
} */
