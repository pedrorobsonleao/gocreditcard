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
=== RUN   TestCheckCard
--- PASS: TestCheckCard (0.00s)
=== RUN   TestBrandCard
--- PASS: TestBrandCard (0.62s)
=== RUN   TestValidateCard
--- PASS: TestValidateCard (0.69s)
PASS
ok      github.com/pedrorobsonleao/gocreditcard  1.334s
```

## how to use

`go get github.com/pedrorobsonleao/gocreditcard`

```golang
func main() {
	check := gocreditcard.Check("3841784397026343307")
	if check {
		fmt.Println("OK")
	}

	brand := gocreditcard.Brand("3841784397026343307")
	if brand == "hipercard" {
		fmt.Println("OK")
	}

	check, brand = gocreditcard.Validate("3841784397026343307")
	fmt.Println(check, brand)
}

```

[1]:https://go.dev/ (Build simple, secure, scalable systems with Go)
[2]:https://github.com/pedrorobsonleao/check-credit-card (a javascript creditcard number validator)
[3]:https://www.bincodes.com/ (BIN Codes Credit Card & Debit Card Number Tools)