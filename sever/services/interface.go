package services

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service interface {
	Handle(c *gin.Context) (any, error)
}

// Handler is a controller bridge between router and service
func Handler(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// First bind the JSON to the service
		err := c.ShouldBindJSON(s)
		if err != nil && err != io.EOF {
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
			return
		}

		log.Printf("[Handler]: bind result: %v\n", s)

		// Then invoke the Handle method of the server
		res, err := s.Handle(c)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, ErrorResponse(err))
		} else {
			// log.Println("StatusOK")
			c.JSON(http.StatusOK, Response(res))
		}
	}
}

type errorResponse struct {
	ErrorStr string `json:"error"`
}

func ErrorResponse(err error) errorResponse {
	if err == nil {
		return errorResponse{}
	}
	return errorResponse{
		ErrorStr: err.Error(),
	}
}

type response struct {
	Data any `json:"data,omitempty"`
}

func Response(obj any) response {
	if obj == nil {
		return response{}
	}
	return response{
		Data: obj,
	}
}
