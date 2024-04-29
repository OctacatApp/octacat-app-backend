package key

type Key string

const (
	KeyClaims = Key("claims")
	KeyCfg    = Key("cfg")
	KeyHeader = Key("headers")
)
