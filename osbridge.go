package osbridge

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
)

var NilOsBridge = (*OsBridge)(nil)
var _ OsBridger = NilOsBridge

type OsBridge struct {
	ProjectName       Name
	ProjectDir        Dir
	UserDataPath      Path
	CachePath         Path
	AdminPath         Path
	SupportRequestUrl Url
	NamedPathMap      map[Name]Path
	NamedDirMap       map[Name]Dir
}

func NewOsBridge(args ...*Args) *OsBridge {
	if len(args) == 0 {
		args = make([]*Args,1)
	}
	var b OsBridge
	b = OsBridge(*args[0])
	if b.ProjectName == "" {
		b.ProjectName = DefaultProjectName
	}
	if b.ProjectDir == "" {
		b.ProjectDir = DefaultProjectDir
	}
	if b.UserDataPath == "" {
		b.UserDataPath = DefaultUserDataPath
	}
	if b.CachePath == "" {
		b.CachePath = DefaultCachePath
	}
	if b.AdminPath == "" {
		b.AdminPath = DefaultAdminPath
	}
	if b.SupportRequestUrl == "" {
		b.SupportRequestUrl = DefaultSupportRequestUrl
	}
	if b.NamedPathMap == nil {
		b.NamedPathMap = make(PathMap,0)
	}
	if b.NamedDirMap == nil {
		b.NamedDirMap = make(DirMap,0)
	}
	return &b
}

func (me *OsBridge) GetUserHomeDir() Dir {
	homeDir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return Dir(homeDir)
}

func (me *OsBridge) GetProjectName() Name {
	return me.ProjectName
}

func (me *OsBridge) GetProjectDir() string {
	return fmt.Sprintf("%s%c%s",
		me.GetUserHomeDir(),
		os.PathSeparator,
		me.ProjectDir,
	)
}

func (me *OsBridge) GetUserConfigDir() Dir {
	return Dir(fmt.Sprintf("%s%c%s",
		me.GetUserHomeDir(),
		os.PathSeparator,
		me.UserDataPath,
	))
}

func (me *OsBridge) GetAdminRootDir() Dir {
	return Dir(fmt.Sprintf("%s%c%s",
		me.GetUserConfigDir(),
		os.PathSeparator,
		me.AdminPath,
	))
}
func (me *OsBridge) GetCacheDir() Dir {
	return Dir(fmt.Sprintf("%s%c%s",
		me.GetUserConfigDir(),
		os.PathSeparator,
		me.CachePath,
	))
}

func (me *OsBridge) GetSupportRequestUrl() Url {
	return me.SupportRequestUrl
}

func (me *OsBridge) GetNamedPath(pn Name) Url {
	p, _ := me.NamedPathMap[pn]
	return p
}
func (me *OsBridge) GetNamedDir(dn Name) Url {
	p, _ := me.NamedDirMap[dn]
	return p
}

func (me *OsBridge) GetNamedPathMap() PathMap {
	return me.NamedPathMap
}

func (me *OsBridge) GetNamedDirMap() DirMap {
	return me.NamedDirMap
}
