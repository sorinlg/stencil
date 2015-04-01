package opts

import (
	log "github.com/Sirupsen/logrus"
	"strings"
)

func Parse_argument(arg string) (string, string) {
	pair := strings.Split(arg, ":")
	if len(pair) != 2 {
		log.WithFields(log.Fields{
			"arg": arg,
		}).Fatal("bad template argument; expected \"/path/to/template:/path/to/destination\"")
	}

	return pair[0], pair[1]
}
