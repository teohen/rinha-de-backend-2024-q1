package handler

import (
	"github.com/gin-gonic/gin"
)

func respondWithError(c *gin.Context, code int) {
	//respondWithJSON(w, code)
}

func respondWithJSON(c *gin.Context, code int) {
	//w.Header().Add("Content-Type", "application/json")
	//w.WriteHeader(code)
}
