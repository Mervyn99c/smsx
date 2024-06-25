// Copyright 2024 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/Rosas99/smsx.
//

// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package zerrors

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 未知错误，服务器内部错误
func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Unknown.String() && e.Code == 500
}

// 未知错误，服务器内部错误
func ErrorUnknown(format string, args ...any) *errors.Error {
	return errors.New(500, ErrorReason_Unknown.String(), fmt.Sprintf(format, args...))
}

// 无效参数错误
func IsInvalidParameter(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_InvalidParameter.String() && e.Code == 400
}

// 无效参数错误
func ErrorInvalidParameter(format string, args ...any) *errors.Error {
	return errors.New(400, ErrorReason_InvalidParameter.String(), fmt.Sprintf(format, args...))
}

// 未找到错误
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NotFound.String() && e.Code == 400
}

// 未找到错误
func ErrorNotFound(format string, args ...any) *errors.Error {
	return errors.New(400, ErrorReason_NotFound.String(), fmt.Sprintf(format, args...))
}

// 未经授权错误
func IsUnauthorized(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Unauthorized.String() && e.Code == 401
}

// 未经授权错误
func ErrorUnauthorized(format string, args ...any) *errors.Error {
	return errors.New(401, ErrorReason_Unauthorized.String(), fmt.Sprintf(format, args...))
}

// 禁止访问错误
func IsForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_Forbidden.String() && e.Code == 403
}

// 禁止访问错误
func ErrorForbidden(format string, args ...any) *errors.Error {
	return errors.New(403, ErrorReason_Forbidden.String(), fmt.Sprintf(format, args...))
}

// 缺少幂等性令牌错误
func IsIdempotentMissingToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_IdempotentMissingToken.String() && e.Code == 400
}

// 缺少幂等性令牌错误
func ErrorIdempotentMissingToken(format string, args ...any) *errors.Error {
	return errors.New(400, ErrorReason_IdempotentMissingToken.String(), fmt.Sprintf(format, args...))
}

// 幂等性令牌已过期错误
func IsIdempotentTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_IdempotentTokenExpired.String() && e.Code == 400
}

// 幂等性令牌已过期错误
func ErrorIdempotentTokenExpired(format string, args ...any) *errors.Error {
	return errors.New(400, ErrorReason_IdempotentTokenExpired.String(), fmt.Sprintf(format, args...))
}
