package logic

import (
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type ValidatorWrapper struct {
	Name  string
	Regex string
}

func newValidator() *validator.Validate {
	validator := validator.New()
	validator.RegisterValidation("text", validateText)
	validator.RegisterValidation("timeDate", validateDate)
	validator.RegisterValidation("timeHour", validateTime)
	validator.RegisterValidation("amount", validateAmount)
	return validator
}

func validateDate(fl validator.FieldLevel) bool {
	return parseTime("2006-01-02", fl)
}

func validateTime(fl validator.FieldLevel) bool {
	return parseTime("15:04", fl)
}

func parseTime(layout string, fl validator.FieldLevel) bool {
	if _, err := time.Parse(layout, fl.Field().String()); err != nil {
		return false
	}
	return true
}

func validateText(fl validator.FieldLevel) bool {
	return validateRegExp(`^[\w\s\-&]+$`, fl)
}

func validateAmount(fl validator.FieldLevel) bool {
	return validateRegExp(`^\d+\.\d{2}$`, fl)
}

func validateRegExp(rgxp string, fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(rgxp)
	return regex.MatchString(fl.Field().String())
}
