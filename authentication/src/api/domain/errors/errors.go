package errors

type InputError struct {
	Msg string
}

func (e *InputError) Error() string {
	return e.Msg
}

func NewInputError(msg string) *InputError {
	return &InputError{
		Msg: msg,
	}
}

type ConstraingError struct {
	Msg string
}

func (e *ConstraingError) Error() string {
	return e.Msg
}

func NewConstraingError(msg string) *ConstraingError {
	return &ConstraingError{
		Msg: msg,
	}
}

type PasswordHashError struct {
	Msg string
}

func (e *PasswordHashError) Error() string {
	return e.Msg
}

func NewPasswordHashError(msg string) *PasswordHashError {
	return &PasswordHashError{
		Msg: msg,
	}
}

type UserNotFoundError struct {
	Msg string
}

func (e *UserNotFoundError) Error() string {
	return e.Msg
}

func NewUserNotFoundError(msg string) *UserNotFoundError {
	return &UserNotFoundError{
		Msg: msg,
	}
}

type UserTokenUpdateError struct {
	Msg string
}

func (e *UserTokenUpdateError) Error() string {
	return e.Msg
}

func NewUserTokenUpdateError(msg string) *UserTokenUpdateError {
	return &UserTokenUpdateError{
		Msg: msg,
	}
}

type JwtParseError struct {
	Msg string
}

func (e *JwtParseError) Error() string {
	return e.Msg
}

func NewJwtParseError(msg string) *JwtParseError {
	return &JwtParseError{
		Msg: msg,
	}
}
