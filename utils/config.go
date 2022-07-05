package utils

//DeployableConfig - The required global config for the deployable
type DeployableConfig struct {
	Cellular   RPCConfig
	Powerline  RPCConfig
	JWTPubKeys JWTKeys
	S2S        S2SConfig
}

type RPCConfig struct {
	ClusterHostname string `json:"cluster_hostname"`
}

type S2SConfig struct {
	Host      string
	TokenPath string
}

type JWTKeys struct {
	Keys []string `json:"pub_keys"`
}

//Config - The global config object
var Config DeployableConfig

var Version string = "local"

func getConfig() {
	Config = DeployableConfig{}
}
