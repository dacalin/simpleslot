package domain

import "errors"

var ErrorRng = errors.New("RNG Error")
var ErrorMoneyCurrency = errors.New("Money::CurrencyError")
var ErrorMoneyAmount = errors.New("Money::AmountError")
var ErrorVisibleReel = errors.New("VisibleReel::Error")