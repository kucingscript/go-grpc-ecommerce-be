package utils

import "github.com/kucingscript/go-grpc-ecommerce-be/pb/common"

func SuccessResponse(message ...string) *common.BaseResponse {
	msg := "Success"

	if len(message) > 0 {
		msg = message[0]
	}

	return &common.BaseResponse{
		StatusCode: 200,
		Message:    msg,
		IsError:    false,
	}
}

func ValidationErrorResponse(validationErrors []*common.ValidationError) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode:       400,
		Message:          "Validation error",
		IsError:          true,
		ValidationErrors: validationErrors,
	}
}
