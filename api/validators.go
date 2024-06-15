package api

import (
	"net/url"
	"strconv"
)

const (
	QueryParam_Value        string = "value"
	QueryParam_Algo         string = "algo"
	QueryParam_PixelPattern string = "pixel_pattern"
	QueryParam_Dimension    string = "dimension"
	QueryParam_DarkMode     string = "dark_mode"
)

func validateGenerateAvatar(qp url.Values) error {
	if qp.Get(QueryParam_Value) == "" {
		return ErrRequiredValue
	}
	// validate algorithm
	if qp.Has(QueryParam_Algo) {
		val, err := strconv.Atoi(qp.Get(QueryParam_Algo))
		if err != nil || (val != 0 && val != 1) {
			return ErrInvalidAlgo
		}
	}
	// validate pixel pattern
	if qp.Has(QueryParam_PixelPattern) {
		val, err := strconv.Atoi(qp.Get(QueryParam_PixelPattern))
		if err != nil || (val != 5 && val != 7 && val != 9) {
			return ErrInvalidPixelPattern
		}
	}
	// validate dimension
	if qp.Has(QueryParam_Dimension) {
		val, err := strconv.Atoi(qp.Get(QueryParam_Dimension))
		if err != nil || val <= 0 {
			return ErrInvalidDimension
		}
	}
	// validate dark mode
	if qp.Has(QueryParam_DarkMode) {
		if _, err := strconv.ParseBool(qp.Get(QueryParam_DarkMode)); err != nil {
			return ErrInvalidDarkMode
		}
	}
	return nil
}
