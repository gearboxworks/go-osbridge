package osbridge

type OsBridger interface {
	GetUserHomeDir() Dir
	GetProjectName() Name
	GetProjectDir() Dir
	GetUserConfigDir() Dir
	GetAdminRootDir() Dir
	GetCacheDir() Dir
	GetSupportRequestUrl() Url
	GetNamedPath(pn Name) Url
	GetNamedDir(dn Name) Url
	GetNamedPathMap() PathMap
	GetNamedDirMap() DirMap
}
