package Token

func IsTokenAvailable(token string) bool {
	for _, t := range availableTokens {
		if token == t {
			return true
		}
	}
	return false
}
