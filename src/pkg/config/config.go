package config

type ConfigBasic struct {
	GithubApiHost string
}

func NewBasic() *ConfigBasic {
	cb := &ConfigBasic{}
	cb.GithubApiHost = "https://api.github.com"
	return cb
}
