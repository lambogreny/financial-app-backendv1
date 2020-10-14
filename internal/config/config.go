package config

import "github.com/namsral/flag"

//DataDirectory is th path used for loading templates/database migrations
var DataDirectory = flag.String("data-directory", "", "Path for loading")
