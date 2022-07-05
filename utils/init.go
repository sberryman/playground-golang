package utils

func init() {
	// get our config
	getConfig()

	// override the s2s path!
	Config.S2S.TokenPath = "/workspaces/playground-golang/.s2s_token"
	Config.S2S.Host = "https://api-auth.dev.bird.co"
	Config.Cellular.ClusterHostname = "https://api-cellular.dev.bird.co"
	Config.Powerline.ClusterHostname = "https://pl-manage.dev.bird.co"
}
