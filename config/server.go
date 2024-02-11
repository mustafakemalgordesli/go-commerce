package config

type ServerConfiguration struct {
	Host                       string
	Port                       string
	AccessTokenSecret          string
	RefreshTokenSecret         string
	AccessTokenExpireDuration  int
	RefreshTokenExpireDuration int
	LimitCountPerRequest       float64
}
