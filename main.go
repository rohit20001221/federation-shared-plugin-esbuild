package main

import (
	"fmt"
	"os"
	"path"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/rohit20001221/js-bundler/plugins"
)

func main() {
	result := api.Build(api.BuildOptions{
		EntryPoints: []string{"example/app.jsx"},
		Bundle:      true,
		Format:      api.FormatESModule,
		Outdir:      path.Join(".", "dist"),
		Write:       true,
		Define: map[string]string{
			"process.env.NODE_ENV": `"production"`,
		},
		Plugins: []api.Plugin{
			plugins.NewSharedPlugin(),
		},
		JSX:       api.JSXAutomatic,
		Splitting: true,
	})

	if len(result.Errors) > 0 {
		for _, err := range result.Errors {
			fmt.Println("[x] bundle error:", err.Text)
		}

		os.Exit(1)
	}
}
