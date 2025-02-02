// Code generated by fabricator-generate-plugin-go
//
// Modifications in code regions will be lost during regeneration!

package wire

// region CODE_REGION(version)
import (
	_ "embed"

	"code.cestus.io/libs/buildinfo"
)

//go:embed version.yml
var version string

func init() {
	buildinfo.GenerateVersionFromVersionYaml(GetVersionYaml(), "code.cestus.io/tools/wire")
}

func GetVersionYaml() []byte {
	return []byte(version)
}

//endregion
