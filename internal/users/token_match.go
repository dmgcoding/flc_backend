package users

func IsTokenMatch(token string) bool {
	return token == SampleTokenHeader
}
