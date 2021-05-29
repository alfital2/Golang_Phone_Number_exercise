package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

func Number(s string) (string, error) {
	// get only the digits from tfhe string representation of the number
	num, error := parseNumber(s)
	if error != nil {
		return "", error
	}
	// if the length is 11 , i.e num has country code number, check if the first digit is 1
	if len(num) == MAX_DIGITS {
		if string(num[0]) != COUNTRY_CODE_VALUE {
			return "", errors.New(ERR_COUNTRY_CODE)
		}
		num = num[1:]
	}
	// check if a the number meets with the requirements of area code, exchange code and correct length
	error = validateNumber(num)

	if error != nil {
		return "", error
	}
	return num, nil
}

// returns an error message if the number is not in the correct format
// if no error occured , return nil
func validateNumber(num string) error {
	if validateAreaCode(num) {
		return errors.New(ERR_AREA_CODE_MIN_VAL)
	}
	if validateExchangeCode(num) {
		return errors.New(ERR_EXCHANGE_CODE_MIN_VAL)
	}
	if validateNumLen(num) {
		return errors.New(ERR_DIGITS_RANGE)
	}
	return nil
}

// check that the area code is in the correct format
func validateAreaCode(num string) bool {
	return num[0] == '1' || num[0] == '0'
}

// check that the exchange code is in the correct format
func validateExchangeCode(num string) bool {
	return num[3] == '1' || num[3] == '0'
}

// check that the length of the number is correct
func validateNumLen(num string) bool {
	return len(num) < MIN_DIGITS
}

func parseNumber(s string) (string, error) {
	reg, err := regexp.Compile("[^0-9]+")
	if err != nil {
		return "ERROR", errors.New("error creating regex")
	}
	processedString := reg.ReplaceAllString(s, "")
	if len(processedString) > MAX_DIGITS {
		return "", errors.New(ERR_MAX_DIGITS_EXCEEDED)
	}
	return processedString, nil
}

//extracts the area code part from the number
func AreaCode(s string) (string, error) {
	num, error := Number(s)
	if error != nil {
		return "", error
	}
	// Extract the area code from the number
	areaCode := num[0:3]
	return areaCode, nil
}

//outputs the number in the correct format
func Format(s string) (string, error) {
	num, error := Number(s)
	if error != nil {
		return "", error
	}
	areaCode, exchangeCode, subscriberNumber := num[0:3], num[3:6], num[6:]
	numberFormatted := fmt.Sprintf("(%s) %s-%s", areaCode, exchangeCode, subscriberNumber)
	return numberFormatted, nil
}
