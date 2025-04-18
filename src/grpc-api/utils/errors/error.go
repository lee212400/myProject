package errors

import (
	"fmt"

	"github.com/lee212400/myProject/domain/entity"
	"github.com/lee212400/myProject/utils/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
)

type AppCode struct {
	Code        int
	Name        string
	GRPCCode    codes.Code
	LogLevel    LogLevel
	Description string
}

var (
	OK               = AppCode{0, "OK", codes.OK, InfoLevel, "Success"}
	InvalidArgument  = AppCode{1, "InvalidArgument", codes.InvalidArgument, WarnLevel, "Invalid input"}
	NotFound         = AppCode{2, "NotFound", codes.NotFound, WarnLevel, "Resource not found"}
	AlreadyExists    = AppCode{3, "AlreadyExists", codes.AlreadyExists, WarnLevel, "Already exists"}
	Internal         = AppCode{4, "Internal", codes.Internal, ErrorLevel, "Internal server error"}
	Unauthenticated  = AppCode{5, "Unauthenticated", codes.Unauthenticated, InfoLevel, "Login required"}
	PermissionDenied = AppCode{6, "PermissionDenied", codes.PermissionDenied, InfoLevel, "Permission denied"}
)

type AppError struct {
	Code    AppCode
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func New(ctx *entity.Context, code AppCode, msg string) *AppError {
	logMessage(ctx, code, msg)
	return &AppError{
		Code:    code,
		Message: msg,
	}
}

func WithError(ctx *entity.Context, code AppCode, err error) *AppError {
	logMessage(ctx, code, err.Error())
	return &AppError{
		Code: code,
		Err:  err,
	}
}

func (e *AppError) Generate(ctx *entity.Context) error {
	msg := ""
	if e.Message != "" || len(e.Message) > 0 {
		msg = e.Message
	} else {
		msg = e.Err.Error()
	}

	switch e.Code.LogLevel {
	case WarnLevel, ErrorLevel:
		logMessage(ctx, e.Code, msg)
	}

	return status.Error(e.Code.GRPCCode, msg)
}

func logMessage(ctx *entity.Context, code AppCode, msg string) {
	logger := logger.WithContext(ctx)
	switch code.LogLevel {
	case DebugLevel:
		logger.Debug(msg)
	case InfoLevel:
		logger.Info(msg)
	case WarnLevel:
		logger.Warn(msg)
	case ErrorLevel:
		logger.Error(msg)
	}
}
