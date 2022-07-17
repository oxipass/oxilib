package oxilib

import "embed"

var (

	//go:embed templates/*.json
	templatesResources embed.FS
)
