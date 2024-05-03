package gocreditcard

import (
	"errors"
	"fmt"
	"regexp"
)

type Flag uint8

const (
	ELO Flag = iota
	JCB
	AMEX
	VISA
	DINERS
	MASTER
	VOYAGER
	ENROUTE
	DISCOVER
	UNIONPAY
	HIPERCARD
)

type card string

// Number return the creditcard number
func (c card) Number() string {
	return string(c)
}

// Flag return the flag to creditcard or error case unknow the flag of the cardnumber
func (c card) Flag() (f string, e error) {
	for i := ELO; i <= HIPERCARD; i++ {
		match, _ := regexp.MatchString(creditcardPatterns[i], c.Number())
		if match {
			f = i.String()

			if f == "unknow" {
				e = errors.New("unknow creditcard flag")
			}
		}
	}
	return
}

var creditcardPatterns = [...]string{
	"^(5067|509[0-9]|6277|6363|650[0-9]|651[67]|6550)[0-9]{12}$",
	"^35\\d{14}$",
	"^3[47][0-9]{13}$",
	"^4[0-9]{12}(?:[0-9]{3})?$",
	"^3(?:0[0-5]|[68][0-9])[0-9]{11}$",
	"^5[1-5][0-9]{14}$",
	"^8699\\d{12}$",
	"^(2014|2149)\\d{12}$",
	"^6011\\d{12}$",
	"^(622305|622698|621483|622202)\\d{10}$",
	"^(606282\\d{10}(\\d{3})?)|(3841\\d{15})$",
}

func (f Flag) String() string {
	switch f {
	case ELO:
		return "elo"
	case JCB:
		return "jcb"
	case AMEX:
		return "amex"
	case VISA:
		return "visa"
	case DINERS:
		return "diners"
	case MASTER:
		return "master"
	case VOYAGER:
		return "voyager"
	case ENROUTE:
		return "endroute"
	case DISCOVER:
		return "discover"
	case UNIONPAY:
		return "unionpay"
	case HIPERCARD:
		return "hipercard"
	}

	return "unknown"
}

const zero rune = 48

func luhn(cardNumber string) bool {
	charAt := []rune(cardNumber)

	var flag bool = false

	oddorEven := len(cardNumber) & 1
	var sum rune = 0

	for pos := 0; pos < len(charAt); pos++ {
		digit := charAt[pos] - zero

		if ((pos & 1) ^ oddorEven) == 0 {
			digit *= 2

			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	if sum != 0 {
		sum %= 10
		if sum == 0 {
			flag = true
		}
	}
	return flag
}

// Parse identify and parse the string cardnumber and return an card or error
func Parse(cardnumber string) (c card, e error) {
	match, err := regexp.MatchString("^[0-9]{14,19}$", cardnumber)

	if match && err == nil {
		if luhn(cardnumber) {
			c = card(cardnumber)
		} else {
			e = errors.New("invalid check digit")
		}
	} else {
		e = fmt.Errorf("invalid parameter")
	}

	return
}
