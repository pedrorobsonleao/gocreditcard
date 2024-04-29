# GoCreditCard

The [GoLang](https://go.dev/ (Build simple, secure, scalable systems with Go)) version to [check-credit-card](https://github.com/pedrorobsonleao/check-credit-card (a javascript creditcard number validator)) the code has with a base the specification [BIN (The Bank Number Identification)](https://www.bincodes.com/ (BIN Codes Credit Card & Debit Card Number Tools)).

This implementation have support to then BINS:

```golang
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
```
## test

```bash
$ go test -v
=== RUN   TestCreditcard
--- PASS: TestCreditcard (1.18s)
=== RUN   TestCreditcardFlag
--- PASS: TestCreditcardFlag (1.11s)
=== RUN   TestCreditcardInvalidValue
--- PASS: TestCreditcardInvalidValue (0.00s)
=== RUN   TestCreditcardSmallValue
--- PASS: TestCreditcardSmallValue (0.00s)
=== RUN   TestCreditcardLargeValue
--- PASS: TestCreditcardLargeValue (0.00s)
=== RUN   TestCreditcardInvalidCheckDigit
--- PASS: TestCreditcardInvalidCheckDigit (0.00s)
PASS
ok      github.com/pedrorobsonleao/gocreditcard 2.331s
```

## how to use

`go get github.com/pedrorobsonleao/gocreditcard`

```golang
func main() {

	number := "6011777162346388"
	card, err := gocreditcard.Parse(number)

	if err != nil {
		t.Errorf(`%s parse error [%s]`, number, err)
	} else {
		flag, err := card.Flag()
		if err != nil {
			t.Errorf(`%s flag error [%s]`, number, err)
		}

		if s.brand != flag {
			t.Errorf(`%s invalid flag [%s] [%s]`, card.Number(), flag, s.brand)
		}
	}

}

```

[1]:https://go.dev/ (Build simple, secure, scalable systems with Go)
[2]:https://github.com/pedrorobsonleao/check-credit-card (a javascript creditcard number validator)
[3]:https://www.bincodes.com/ (BIN Codes Credit Card & Debit Card Number Tools)