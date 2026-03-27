package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}

type Cents uint64

var ErrInvalidDollarFormat = errors.New("invalid dollar format. Must have two decimal places and be presented as a string.")

// converts cents to dollars for responses
func (c Cents) MarshalJSON() ([]byte, error) {
    jsonValue := fmt.Sprintf("%d.%02d", c/100, c%100)

    quotedJSONValue := strconv.Quote(jsonValue)

    return []byte(quotedJSONValue), nil
}

// converts dollars to cents for input
func (c *Cents) UnmarshalJSON(jsonValue []byte) error {
    unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
    if err != nil {
        return ErrInvalidDollarFormat
    }

    centString := strings.ReplaceAll(unquotedJSONValue, ".", "")

    cents, err := strconv.ParseUint(centString, 10, 64)
    if err != nil{
        return ErrInvalidDollarFormat
    }
    *c = Cents(cents)

    return nil
}