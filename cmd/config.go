package cmd

import "os"

const switchConfigPath = ".config/gitentry"
const ghConfigdir = ".config/gh"
const ghConfigFile = "hosts.yml"

// const regarding user credentials
const credentialFormatter = "https://%s:%s@github.com"
const credentialCacheFile = ".configcredentials/cache.out"

var home, _ = os.UserHomeDir()
