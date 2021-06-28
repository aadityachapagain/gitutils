package cmd

import "os"

const switchConfigPath = ".config/gitentry"
const ghConfigdir = ".config/gh"
const ghConfigFile = "hosts.yml"
const credentialFormatter = "https://%s:%s@github.com"

var home, _ = os.UserHomeDir()
