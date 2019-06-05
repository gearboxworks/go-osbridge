package osbridge

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"runtime"
)


func Get(args ...*Args) OsBridger {
	osb := NewOsBridge(args...)

	switch runtime.GOOS {
	case "darwin":
	case "windows":
	case "linux":
	default:
		msg := "Sadly, %s does not currently run on '%s.'\n"+
			"If you would like to offer us support to change "+
			"that please submit a request at %s.\n"
		msg = fmt.Sprintf(msg,
			osb.GetProjectName(),
			runtime.GOOS,
			osb.GetSupportRequestUrl(),
		)
		log.Fatal(msg)
	}
	return osb
}
