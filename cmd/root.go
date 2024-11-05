package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "web-to-embed",
	Short: "Convert static web assets into embeddable C++ code",
	Long: `Web-To-Embed is a CLI tool for converting an entire static website folder into a single C++ header file.
	
	This tool generates embedded C++ arrays for each file, making it easy to include static web assets in embedded systems
	or applications that need to serve files without a filesystem.
	
	Usage examples:
	- Convert a directory of web files:
		web-to-embed convert ./my-site -o embedded_site.h

	This command will produce a header file with embedded byte arrays, paths, sizes, and MIME types of each web asset,
	all accessible from within your C++ application.
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
