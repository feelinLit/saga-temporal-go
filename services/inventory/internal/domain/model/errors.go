package model

import "errors"

var ErrStockAlreadyReserved = errors.New("stock should not be already reserved by the same order")
