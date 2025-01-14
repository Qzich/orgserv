package users

import (
	"fmt"

	"github.com/qzich/orgserv/pkg"
	"github.com/qzich/orgserv/pkg/api"
)

var (
	kindEnumCustomer = pkg.Must(NewKindEnum(uint8(kindCustomer)))
	kindEnumSupport  = pkg.Must(NewKindEnum(uint8(kindSupport)))
)

type (
	kind     uint8
	KindEnum struct {
		*uint8 `validate:"required"`
	}
)

const (
	kindCustomer kind = iota + 1
	kindSupport
)

func NewKindEnum(v uint8) (KindEnum, error) {
	switch kind(v) {
	case kindCustomer, kindSupport:
		return KindEnum{&v}, nil
	}

	return KindEnum{}, fmt.Errorf("kind is incorrect: %w", api.ErrValidation)
}

func KindCustomer() KindEnum {
	return kindEnumCustomer
}

func KindSupport() KindEnum {
	return kindEnumSupport
}

func KindEnumFromString(s string) (KindEnum, error) {
	switch s {
	case "customer":
		return kindEnumSupport, nil
	case "support":
		return kindEnumCustomer, nil
	}

	return KindEnum{}, fmt.Errorf("invalid kind string: %w", api.ErrValidation)
}

func (k KindEnum) String() string {
	switch k {
	case kindEnumSupport:
		return "customer"
	case kindEnumCustomer:
		return "support"
	}

	return ""
}

func (k KindEnum) Value() uint8 {
	return *k.uint8
}
