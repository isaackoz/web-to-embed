package cmd

import (
	"fmt"

	"github.com/isaackoz/web-to-embed/convert"
	"github.com/isaackoz/web-to-embed/convert/generateFile"
	"github.com/spf13/cobra"
)

var InputDir string
var OutputDir string
var Progmem bool
var Verbose bool

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert static web assets into embeddable C++ code",
	Run: func(cmd *cobra.Command, args []string) {
		// Get all assets from dir
		assets, err := convert.GetAssetsFromDir(InputDir)
		if err != nil {
			fmt.Println("Error getting assets from dir:", err)
			return
		}

		// Log the total size of all the assets
		if Verbose {
			size := (convert.GetTotalSizeOfAssets(assets)) / 1000
			fmt.Println("Total size of all assets:", size, "kb")
		}

		// generate a cpp header file containing all assets
		err = generateFile.Generate(assets, generateFile.Options{OutputDir: OutputDir, Progmem: Progmem})
		if err != nil {
			panic(err)
		}
		//
		fmt.Println("Created header file succesfully in", OutputDir)
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(&InputDir, "input", "i", "", "Input directory (required)")
	convertCmd.MarkFlagRequired("input")

	convertCmd.Flags().StringVarP(&OutputDir, "output", "o", "./", "Output Dir")

	convertCmd.Flags().BoolVar(&Progmem, "progmem", false, "Store data on flash instead of SRAM for Arduino (default False)")

	convertCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose output for debugging")

}
