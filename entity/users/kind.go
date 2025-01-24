package users

type KindEnum uint8

const (
	KindCustomer KindEnum = iota + 1
	KindSupport
)

func ParseKindFromString(s string) (KindEnum, error) {
	switch s {
	case "customer":
		return KindCustomer, nil
	case "support":
		return KindSupport, nil
	}

	return 0, KindIsNotCorrect
}

func (k KindEnum) Validate() error {
	switch k {
	case KindCustomer, KindSupport:
		return nil
	}

	return KindIsNotCorrect
}

func (k KindEnum) String() string {
	switch k {
	case KindCustomer:
		return "customer"
	case KindSupport:
		return "support"
	}

	return ""
}

func (k KindEnum) Value() uint8 {
	return uint8(k)
}
