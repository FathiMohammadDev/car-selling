package helpers

import "github.com/FathiMohammadDev/car-selling/api/validations"

type BaseHttpResponse struct {
	Result           any
	Success          bool
	ResultCode       int
	Error            any
	VlaidationErrors *[]validations.ValidationError
}

func GenerateBaseRes(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		ResultCode: resultCode,
		Success:    success,
	}
}
func GenerateBaseResWithErr(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		ResultCode: resultCode,
		Success:    success,
		Error:      err.Error(),
	}
}
func GenerateBaseResWithValidationError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		ResultCode:       resultCode,
		Success:          success,
		VlaidationErrors: validations.GetValidationError(err),
	}
}
