package entity

type Password string

// TOOD: add password specific validation rules and other error
func (pass Password) Validate() error {
	if len(pass) == 0 {
		return ErrPasswordIncorrect
	}

	return nil
}
