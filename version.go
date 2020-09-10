package main

var (
	Version   = "unknown"
	BuildDate = "unknown"
)

var meta = map[string]string{
	"version":   Version,
	"buildDate": BuildDate,
}

func GetVersion() string {
	return meta["version"]
}

func GetBuildDate() string {
	return meta["buildDate"]
}
