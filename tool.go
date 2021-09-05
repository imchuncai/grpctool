package grpctool

import (
	"context"
	"errors"
	"runtime/debug"

	"github.com/imchuncai/log"
	"google.golang.org/grpc"
)

func ErrorInterceptor(logger log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() { err = doRecover(logger) }()
		return handler(ctx, req)
	}
}

func doRecover(logger log.Logger) error {
	if err := recover(); err != nil {
		logger.Log(log.Error, err, string(debug.Stack()))
		return errors.New("server response an error")
	}
	return nil
}
