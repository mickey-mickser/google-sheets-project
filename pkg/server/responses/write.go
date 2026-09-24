package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// BadRequest sends HTTP 400 with standardized error payload
func BadRequest(c *gin.Context, err ResponseError) {
	c.JSON(http.StatusBadRequest, err)
}

// Unauthorized sends HTTP 401 with standardized error payload
func Unauthorized(c *gin.Context, err ResponseError) {
	c.JSON(http.StatusUnauthorized, err)
}

// NotFound sends HTTP 404 with a generic body (can be ResponseError or other structure)
func NotFound(c *gin.Context, err ResponseError) {
	c.JSON(http.StatusNotFound, err)
}

// Conflict sends HTTP 409 with a generic body
func Conflict(c *gin.Context, err ResponseError) {
	c.JSON(http.StatusConflict, err)
}

// InternalError sends HTTP 500 with InternalErrorResponse
func InternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, InternalErrorResponse)
}

// OK sends HTTP 200 with any body
func OK(c *gin.Context, body interface{}) {
	c.JSON(http.StatusOK, body)
}
