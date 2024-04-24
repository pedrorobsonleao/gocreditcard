package gocreditcard

import (
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

var Flags = [...]string{
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

// Get Brand to CardNumber
func Brand(cardNumber string) string {
	for i := ELO; i <= HIPERCARD; i++ {
		match, _ := regexp.MatchString(Flags[i], cardNumber)
		if match {
			return i.String()
		}
	}
	return "unknown"
}

// Check verification digit
func Check(cardNumber string) bool {
	return luhn(cardNumber)
}

// Check verification digit and get Brand
func Validate(cardNumber string) (bool, string) {
	check := Check(cardNumber)
	brand := Brand(cardNumber)
	return check, brand
}
