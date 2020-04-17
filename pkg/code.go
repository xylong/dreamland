package pkg

const (
	SUCCESS = 0
	ERROR   = 10000

	RedisError = 10001

	UNAUTHORIZED     = 20001
	TokenExpired     = 20002
	TokenNotValidYet = 20003
	TokenMalformed   = 20004
	TokenInvalid     = 20005
)
