package middleware

import "github.com/gin-gonic/gin"

type AuthMiddleware struct {
	errorHandler func(c *gin.Context)
	//authService  service.AuthService
}

func NewAuthMiddleware(errorHandler func(c *gin.Context)) *AuthMiddleware {
	return &AuthMiddleware{errorHandler: errorHandler}
}

////
//func (a *AuthMiddleware) SetCurrentUser() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		ctx := c.Request.Context()
//
//		authHeader := c.GetHeader("Authorization")
//		if authHeader == "" {
//			fmt.Println("empty Authorization header")
//			return
//		}
//
//		token, err := a.ParseTokenFromAuthHeader(authHeader, "Bearer")
//		if err != nil {
//			fmt.Println("error parsing token from Authorization header", "err", err.Error())
//			return
//		}
//
//		tokenParts := strings.Split(authHeader, " ")
//		if len(tokenParts) != 2 {
//			fmt.Println(fmt.Sprintf("splited token contains %d parts, expected 2", len(tokenParts)))
//			return
//		}
//
//		tokenType, token := tokenParts[0], tokenParts[1]
//		if tokenType != "Bearer" {
//			fmt.Println(
//				"unexpected token type",
//				"tokenType", tokenType,
//				"expectedTokenType", "Bearer",
//			)
//			return
//		}
//
//		user, err := a.Authenticate(ctx, token)
//		if err != nil {
//			return
//		}
//		ctx = context.WithValue(ctx, "user", user)
//		c.Request = c.Request.WithContext(ctx)
//
//		c.Next()
//	}
//}
//
//func (a *AuthMiddleware) Auth(roles ...string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		currentUser, err := a.FromContext(c)
//		if err != nil {
//			c.Error(err)
//			a.errorHandler(c)
//			c.Abort()
//			return
//		}
//
//		if roles != nil && !currentUser.HasAllRoles(roles...) {
//			fmt.Errorf(
//				"user doesn't have enough privileges",
//				"userID", currentUser.UserID,
//				"roles", roles,
//			)
//
//			c.Error(errors.New("user doesn't have enough privileges"))
//			a.errorHandler(c)
//			c.Abort()
//			return
//		}
//
//		c.Next()
//	}
//}
//
//func (a *AuthMiddleware) ParseTokenFromAuthHeader(header string, prefix string) (string, error) {
//	tokenParts := strings.Split(header, " ")
//	if len(tokenParts) != 2 {
//		return "", fmt.Errorf("splited token contains %d parts, expected 2", len(tokenParts))
//	}
//
//	tokenType, token := tokenParts[0], tokenParts[1]
//	if tokenType != prefix {
//		return "", fmt.Errorf("unexpected token type. tokenType=%q, expectedTokenType=%q", tokenType, prefix)
//	}
//
//	return token, nil
//}
//
//func (a *AuthMiddleware) FromContext(c *gin.Context) (model.User, error) {
//	user := c.Request.Context().Value("user").(model.User)
//	user, err := a.authService.GetByID(c.Request.Context(), user.UserID)
//	if err != nil {
//		return model.User{}, err
//	}
//
//	return user, nil
//}
//
//func (a *AuthMiddleware) NewContext(ctx context.Context, user *model.User, authHeader string) context.Context {
//	ctx = context.WithValue(ctx, "user", user)
//	return context.WithValue(ctx, "authHeader", authHeader)
//}
//
//func (a *AuthMiddleware) Authenticate(ctx context.Context, token string, roles ...string) (model.User, error) {
//	user, err := a.authService.Me(ctx, token)
//	if err != nil {
//		return model.User{}, err
//	}
//
//	if roles != nil && !user.HasAllRoles(roles...) {
//		return model.User{}, errors.New("user doesn't have enough privileges")
//	}
//
//	return user, nil
//}
