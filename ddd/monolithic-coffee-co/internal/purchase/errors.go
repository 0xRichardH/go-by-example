package purchase

import "errors"

var (
	ErrInvalidPurchase       = errors.New("purchase must consist of at least one product")
	ErrInvalidPurchaseAmount = errors.New("invalid purchase amount")
	ErrUnknownPaymentMethod  = errors.New("unknown payment method")
)
