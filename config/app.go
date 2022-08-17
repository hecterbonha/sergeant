package config

func RehearsalURL() string {
	insteadToken := Config("REHEARSAL_URL")

	return insteadToken
}
