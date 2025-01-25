package users

type Name string

func (name Name) Validate() error {
	if len(name) < 4 || len(name) > 255 {
		return ErrNameIsNotCorrect
	}

	return nil
}
