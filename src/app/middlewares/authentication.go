package middlewares

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware 是一个身份验证中间件函数
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		// 解析 Authorization 字段，获取令牌
		token := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		// 在这里进行令牌验证逻辑
		if isValidToken(token) {
			// 令牌验证通过，继续处理请求
			c.Next()
		} else {
			// 令牌验证失败，返回未授权错误
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		}
	}
}

// isValidToken 是一个示例的令牌验证函数
func isValidToken(token string) bool {
	// 在这里实现令牌验证逻辑
	// 可以调用验证服务或查询数据库进行验证
	// 返回 true 表示验证通过，返回 false 表示验证失败
	return true
}
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 从请求头部获取 JWT
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Missing Authorization header")
			return
		}

		// 解析和验证 JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名的密钥，确保与签名时使用的密钥相同
			// 这里使用一个示例的密钥，您应该根据实际情况更换为您自己的密钥
			signingKey := []byte("my-secret-key")
			return signingKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid JWT: %v", err)
			return
		}

		// 验证 JWT 的有效性
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Invalid JWT")
			return
		}

		// 提取 JWT 的载荷信息
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"].(int)
		username := claims["username"].(string)

		// 根据用户信息进行鉴权和授权的判断
		// 这里只是一个简单的示例，您可以根据实际需求进行相应的处理

		// 将用户信息存储到请求的上下文中，以便后续处理
		ctx := context.WithValue(r.Context(), "user_id", userID)
		ctx = context.WithValue(ctx, "username", username)
		r = r.WithContext(ctx)

		// 继续处理下一个中间件或处理程序
		next.ServeHTTP(w, r)
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// 从请求的上下文中获取用户信息
	userID := r.Context().Value("user_id").(int)
	username := r.Context().Value("username").(string)

	fmt.Fprintf(w, "Protected API: UserID=%d, Username=%s", userID, username)
}

func main() {
	// 创建一个简单的 HTTP 服务器
	http.HandleFunc("/protected", protectedHandler)

	// 使用 JWT 中间件保护受保护的 API
	http.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))

	// 启动服务器
	http.ListenAndServe(":8080", nil)
}
