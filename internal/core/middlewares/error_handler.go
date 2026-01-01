package middlewares

import (
	"address-book-server-v2/internal/common/fault"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors.Last().Err

		if ae, ok := err.(*fault.AppError); ok {

			data := gin.H{
				"error": ae.Code,
				"message": ae.Message,
			}

			if len(ae.Details) > 0 {
				data["details"] = ae.Details
			}

			response := gin.H{
				"status": "fail",
				"data": data,
			}
			
			ctx.JSON(ae.StatusCode, response)
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "fail",
			"data": gin.H{
				"error":   "INTERNAL_ERROR",
				"message": "Something went wrong",
			},
		})
	}
}