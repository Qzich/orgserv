package users

type (
	kind uint8
	Kind struct {
		*uint8 `validate:"required"`
	}
)

const (
	kindCustomer kind = iota + 1
	kindSupport
	// kindNew
)

func NewKind(v uint8) Kind {
	switch kind(v) {
	case kindCustomer, kindSupport:
		return Kind{&v}
	}

	panic("invalid kind value")
}

func KindCustomer() Kind {
	return NewKind(uint8(kindCustomer))
}

func KindSupport() Kind {
	return NewKind(uint8(kindSupport))
}

func KindFromString(s string) Kind {
	switch s {
	case "customer":
		return KindCustomer()
	case "support":
		return KindSupport()
	}

	panic("invalid kind string")
}

func (k Kind) String() string {
	switch kind(k.Value()) {
	case kindCustomer:
		return "customer"
	case kindSupport:
		return "support"
	}

	panic("invalid kind value")
}

func (k Kind) Value() uint8 {
	return *k.uint8
}
