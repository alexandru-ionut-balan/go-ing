package grants

const (
	clientCredentials = "client_credentials"
	grantType         = "grant_type"
)

func ClientCredentials() string {
	return grantType + "=" + clientCredentials
}
