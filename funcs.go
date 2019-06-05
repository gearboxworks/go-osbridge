package osbridge

//
//	Go-OsBridge - Operating System Bridge package for use in Go apps
//	Copyright (C) 2019  NewClarity Consulting LLC
//
//	This program is free software: you can redistribute it and/or modify
//	it under the terms of the GNU Affero General Public License as published
//	by the Free Software Foundation, either version 3 of the License, or
//	(at your option) any later version.
//
//	This program is distributed in the hope that it will be useful,
//	but WITHOUT ANY WARRANTY; without even the implied warranty of
//	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//	GNU Affero General Public License for more details.
//
//	You should have received a copy of the GNU Affero General Public License
//	along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

import (
	"fmt"
	"log"
	"runtime"
)

func Get(args ...*Args) OsBridger {
	osb := NewOsBridge(args...)

	switch runtime.GOOS {
	case "darwin":
	case "windows":
	case "linux":
	default:
		msg := "Sadly, %s does not currently run on '%s.'\n" +
			"If you would like to offer us support to change " +
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
