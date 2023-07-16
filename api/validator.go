package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/techschool/simplebank/util"
)

var validCurrency validator.Func = func(fielfLevel validator.FieldLevel) bool {
	if currency, ok := fielfLevel.Field().Interface().(string); ok {
		//check if currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
