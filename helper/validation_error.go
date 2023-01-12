package helper

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorHandle(err error) string {
	reports := []string{}

	castedObject, ok := err.(validator.ValidationErrors)
	if ok {
		for _, v := range castedObject {
			switch v.Tag() {
			case "required":
				reports = append(reports, fmt.Sprintf("%s is required", v.Field()))
			case "min":
				reports = append(reports, fmt.Sprintf("%s value must be greater than %s character", v.Field(), v.Param()))
			case "max":
				reports = append(reports, fmt.Sprintf("%s value must be lower than %s character", v.Field(), v.Param()))
			case "lte":
				reports = append(reports, fmt.Sprintf("%s value must be below %s", v.Field(), v.Param()))
			case "gte":
				reports = append(reports, fmt.Sprintf("%s value must be above %s", v.Field(), v.Param()))
			case "numeric":
				reports = append(reports, fmt.Sprintf("%s value must be numeic", v.Field()))
			}
		}

	}

	report := strings.Join(reports, ", ")

	return report
}
