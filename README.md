# GoCreditCard

The [GoLang](1) version to [check-credit-card](2) the code has with a base the specification [BIN (The Bank Number Identification)](3).

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
=== RUN   TestCheckCard
--- PASS: TestCheckCard (0.00s)
=== RUN   TestBrandCard
--- PASS: TestBrandCard (0.62s)
=== RUN   TestValidateCard
--- PASS: TestValidateCard (0.69s)
PASS
ok      githb.com/pedrorobsonleao/gocreditcard  1.334s
```

## how to use

[1]: https://go.dev/ (Build simple, secure, scalable systems with Go)
[2]: https://github.com/pedrorobsonleao/check-credit-card (a javascript creditcard number validator)
[3]: https://www.bincodes.com/ (BIN Codes Credit Card & Debit Card Number Tools)