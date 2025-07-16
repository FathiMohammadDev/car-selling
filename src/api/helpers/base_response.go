package helpers

import "github.com/FathiMohammadDev/car-selling/api/validations"

type BaseHttpResponse struct {
	Result           any
	Success          bool
	ResultCode       int
	Error            any
	VlaidationErrors *[]validations.ValidationError
}

func generateBaseRes(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		ResultCode: resultCode,
		Success:    success,
	}
}
func generateBaseResWithErr(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		ResultCode: resultCode,
		Success:    success,
		Error:      err.Error(),
	}
}
func generateBaseResWithValidationError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		ResultCode:       resultCode,
		Success:          success,
		VlaidationErrors: validations.GetValidationError(err),
	}
}
