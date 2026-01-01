package middlewares

// import (
// 	"address-book-server-v2/internal/common/utils"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func EnsureUserExistsMiddleware(userRepo *repositories.UserRepository) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		userID := ctx.GetUint("user_id") 
// 		if userID == 0 {
// 			utils.Error(ctx, http.StatusUnauthorized, "user not found in context")
// 			ctx.Abort()
// 			return
// 		}

// 		exists, err := userRepo.ExistsByID(userID)
// 		if err != nil || !exists {
// 			utils.Error(ctx, http.StatusUnauthorized, "user does not exist")
// 			ctx.Abort()
// 			return
// 		}

// 		ctx.Next()
// 	}
// }