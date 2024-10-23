package utils

import (
	"errors"
	"strconv"
	"strings"
)

func PaginationArgsValidation(offset string, limit string) error {
	if offset == "" || limit == "" {
		return errors.New("Offset or limit is empty")
	}
	_, err := strconv.Atoi(offset)
	if err != nil {
		return errors.New("Offset is not a valid number")
	}
	_, err = strconv.Atoi(limit)
	if err != nil {
		return errors.New("Limit is not a valid number")
	}
	return nil
}

func ArgsValidation(id string, onlyNumber bool) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("Identifier cannot be empty")
	}
	if onlyNumber {
		if _, err := strconv.Atoi(strings.TrimSpace(id)); err != nil {
			return errors.New("Identifier must be a number")
		}
	}
	return nil
}
