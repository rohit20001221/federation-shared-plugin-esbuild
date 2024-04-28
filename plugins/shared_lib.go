package plugins

import (
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

func NewSharedPlugin() api.Plugin {
	plugin := api.Plugin{
		Name: "federation:shared_lib",
		Setup: func(build api.PluginBuild) {
			/*
				Intercept import paths called "react" or "react-dom" so esbuild doesn't attempt
				to map them to a file system location. Tag them with the "federation:shared_lib"
				namespace to reserve them for this plugin.
			*/

			build.OnResolve(api.OnResolveOptions{Filter: `^(react|react-dom)$`}, func(args api.OnResolveArgs) (api.OnResolveResult, error) {
				fmt.Println("[x] args.path:", args.Path)
				fmt.Println("[x] args.resolved_path:", args.Importer)

				return api.OnResolveResult{
					Path:      fmt.Sprintf("/__shared_lib/%s.js", args.Path),
					Namespace: "federation:shared_lib",
					External:  true,
				}, nil
			})
		},
	}

	return plugin
}
