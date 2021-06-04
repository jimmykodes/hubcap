package auth

type OAuth interface{
	AuthCodeURL() (state string, url string, err error)
	GetUsername(code string) (username string, err error)
}
