package validator

import (
	"fmt"

	"github.com/ariefsn/intrans/entities"
	val "github.com/go-playground/validator/v10"
)

var validate *val.Validate

func InitValidator() {
	validate = val.New(val.WithRequiredStructEnabled())
}

func Validator() *val.Validate {
	return validate
}

func errorByTag(fe val.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Namespace())
	case "email":
		return fmt.Sprintf("Invalid email")
	}
	return fe.Error() // default error
}

func parseErrors(err error) []string {
	errs := []string{}

	if err == nil {
		return errs
	}

	for _, v := range err.(val.ValidationErrors) {
		errs = append(errs, errorByTag(v))
	}

	return errs
}

func ValidateStruct(s interface{}) (error, []string) {
	err := validate.Struct(s)

	return err, parseErrors(err)
}

func ValidateVar(field interface{}, tag string) (error, []string) {
	err := validate.Var(field, tag)

	return err, parseErrors(err)
}

func ValidateVarMap(data, rules entities.M) entities.M {
	return validate.ValidateMap(data, rules)
}
