package grpc

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

var jwtSecret = []byte("my-key")

type UserClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"user_role"`
	jwt.RegisteredClaims
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	authHeader, exists := md["authorization"]
	if !exists || len(authHeader) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is required")
	}

	tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")

	claims, err := parseJWT(tokenString)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	ctx = context.WithValue(ctx, "user_id", claims.UserID)
	ctx = context.WithValue(ctx, "user_role", claims.Role)

	return handler(ctx, req)
}

func parseJWT(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}

func getUserInfoFromContext(ctx context.Context) (int, string, error) {
	userID, ok := ctx.Value("user_id").(int)
	if !ok {
		return 0, "", fmt.Errorf("user_id not found in context")
	}

	role, ok := ctx.Value("user_role").(string)
	if !ok {
		return 0, "", fmt.Errorf("user_role not found in context")
	}

	return userID, role, nil
}
