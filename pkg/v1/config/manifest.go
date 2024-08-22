package config

//##############################################################################
// Manifest
//##############################################################################

type Manifest struct {
	Runtime          ManifestRuntime
	Name             string
	ContentTypePaths []string
	ContentTypes     []*ContentType
}

//##############################################################################
// ManifestRuntime
//##############################################################################

type ManifestRuntime string

const (
	ManifestRuntimeUnknown ManifestRuntime = "unknown"
	ManifestRuntimeGo      ManifestRuntime = "go"
)
