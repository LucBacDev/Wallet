package interceptor

import (
	"context"
	"errors"
	"log"
	"source-base-go/golang/service/authService/config"
	"source-base-go/golang/service/authService/infrastructure/repository/util"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userIDContextKey = contextKey("user_id")

func AuthUnaryInterceptor(verifier util.Verifier) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		jwtSecretKey := []byte(config.GetString("jwt.secretKey"))
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
		}

		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader[0], "Bearer ")
		log.Println("tokenString:", tokenString)
		if tokenString == "" {
			return nil, status.Errorf(codes.Unauthenticated, "authorization header format must be Bearer {token}")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return nil, status.Errorf(codes.Unauthenticated, "invalid claims")
		}

		getUserId, ok := claims["user_id"].(float64)
		userId := int32(getUserId)
		userIdstring := strconv.Itoa(int(userId))
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "user_id claim missing or invalid")
		}

		tokenExists, err := verifier.Verifier(userId, tokenString)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to verify token: %v", err)
		}
		if tokenExists {
			return nil, status.Error(codes.Unauthenticated, "invalid or revoked token")
		}

		header := metadata.Pairs("user_id", userIdstring)
		grpc.SendHeader(ctx, header)

		newCtx := context.WithValue(ctx, userIDContextKey, userId)

		return handler(newCtx, req)
	}
}

func UserIDFromContext(ctx context.Context) (int32, bool) {
	userId, ok := ctx.Value(userIDContextKey).(int32)
	return userId, ok
}
