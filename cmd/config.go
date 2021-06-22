package cmd

import "os"

const switchConfigPath = ".config/gitentry"
const ghConfigFile = ".config/gh/hosts.yml"

var home, _ = os.UserHomeDir()
