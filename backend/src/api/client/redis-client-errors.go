package client

type RedisUserNotFoundError struct {
	Msg string
}

func (e *RedisUserNotFoundError) Error() string {
	return e.Msg
}

func NewRedisUserNotFoundError(msg string) *RedisUserNotFoundError {
	return &RedisUserNotFoundError{
		Msg: msg,
	}
}

type RedisUserSaveError struct {
	Msg string
}

func (e *RedisUserSaveError) Error() string {
	return e.Msg
}

func NewRedisUserSaveError(msg string) *RedisUserSaveError {
	return &RedisUserSaveError{
		Msg: msg,
	}
}
