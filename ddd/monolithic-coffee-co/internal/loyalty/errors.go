package loyalty

import "errors"

var (
	ErrInvalidPurchaseCount = errors.New("invalid purchase count")
	ErrInsuffientFreeDriks  = errors.New("insuficient free driks")
)
