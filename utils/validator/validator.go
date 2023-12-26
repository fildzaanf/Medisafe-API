package validator

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"talkspace/utils/constant"
)

func IsDataEmpty(data ...interface{}) error {
	for _, value := range data {
		switch v := value.(type) {
		case string:
			if v == "" {
				return errors.New(constant.ERROR_DATA_EMPTY)
			}
		case int:
			if v == 0 {
				return errors.New(constant.ERROR_DATA_EMPTY)
			}
		default:
			return errors.New(constant.ERROR_DATA_TYPE)
		}
	}
	return nil
}

func IsDataValid(data interface{}, validData []interface{}, caseSensitive bool) error {
    dataStr := fmt.Sprintf("%v", data)
    validDataStr := make([]string, len(validData))
    for i, v := range validData {
        validDataStr[i] = fmt.Sprintf("%v", v)
    }

    if !caseSensitive {
        dataStr = strings.ToLower(dataStr)
        for i, v := range validDataStr {
            validDataStr[i] = strings.ToLower(v)
        }
    }

    for _, validValue := range validDataStr {
        if dataStr == validValue {
            return nil
        }
    }

    return errors.New(constant.ERROR_DATA_INVALID + strings.Join(validDataStr, ", "))
}

func IsEmailValid(email string) error {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New(constant.ERROR_EMAIL_FORMAT)
	}
	return nil
}

func IsDateValid(date string) error {
	if date == "" {
		return nil
	}

	dateRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !dateRegex.MatchString(date) {
		return errors.New(constant.ERROR_DATE_FORMAT)
	}

	return nil
}

func IsMinLengthValid(data string, minLength int) error {
    if len(data) < minLength {
        return errors.New(fmt.Sprintf(constant.ERROR_MIN_LENGTH, minLength))
    }
    return nil
}

func IsMaxLengthValid(data string, maxLength int) error {
    if len(data) > maxLength {
        return errors.New(fmt.Sprintf(constant.ERROR_MAX_LENGTH, maxLength))
    }
    return nil
}