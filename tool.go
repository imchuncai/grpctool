package grpctool

import (
	"context"
	"errors"
	"runtime/debug"

	"github.com/imchuncai/log"
	"google.golang.org/grpc"
)

var _logger log.Logger

func SetLogger(logger log.Logger) {
	_logger = logger
}

func chekcLogger() {
	if _logger == nil {
		panic("grpctool: logger is not set")
	}
}

func ErrorInterceptor() grpc.ServerOption {
	chekcLogger()
	return grpc.UnaryInterceptor(
		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
			handler grpc.UnaryHandler) (resp interface{}, err error) {
			defer func() {
				if err != nil {
					return
				}
				if err := recover(); err != nil {
					_logger.Log(log.Error, err, string(debug.Stack()))
					err = errors.New("server response an error")
				}
			}()
			return handler(ctx, req)
		})
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func MustForInit(err error) {
	chekcLogger()
	if err == nil {
		return
	}
	_logger.Log(log.Error, err, string(debug.Stack()))
	panic(err)
}

func Log(prefix log.Prefix, v ...interface{}) {
	chekcLogger()
	_logger.Log(prefix, v...)
}
