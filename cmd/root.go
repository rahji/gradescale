/*
Copyright © 2024 Rob Duarte <me@robduarte.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"math"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool
var points float64

type Grade struct {
	Letter string
	Min    float64
}

var scale []Grade

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gradescale",
	Short: "gradescale creates a letter grade scale, given a max number of points",
	Long: `gradescale creates a letter grade scale based on a max number of points.
The new scale is proportional to an existing reference scale.
The original scale can include fractional numbers.
The new scale uses whole numbers.

`,
	Run: func(cmd *cobra.Command, args []string) {

		if scale == nil {
			fmt.Println("No scale defined. Check the configuration file.")
			os.Exit(1)
		}
		if debug {
			fmt.Printf("Reference scale: %v\n", scale)
			return
		}

		outputNewScale(scale, points)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is ./gradescale.yaml)")
	rootCmd.Flags().BoolVar(&debug, "debug", false, "output debug info instead of the grade scale")
	rootCmd.Flags().Float64Var(&points, "points", 0, "the total number of points in the new scale (required)")
	rootCmd.MarkFlagRequired("points")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name "gradescale" (without extension).
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("gradescale")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if debug {
			fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		}
	}

	// I don't know what can actually make this return an error
	// I've tried all manner of messed up config files and the printf never happens
	if err := viper.UnmarshalKey("scale", &scale); err != nil {
		fmt.Printf("Unable to decode scale struct from config file, %v", err)
	}
}

func outputNewScale(scale []Grade, points float64) {
	fmt.Print("# Letter Grade Scale\n\n")
	fmt.Println("| Letter | Low | High |")
	fmt.Println("| :----- | :-- | :--- |")

	for i, g := range scale {

		thismin := math.Round((float64(points) * g.Min) / 100)

		if i == 0 {
			fmt.Printf("| %s | ≥ %.0f | ≤ %.0f |\n", g.Letter, thismin, points)
			continue
		}

		thismax := math.Round((float64(points) * scale[i-1].Min) / 100)

		fmt.Printf("| %s | ≥ %.0f | < %.0f |\n", g.Letter, thismin, thismax)

	}
}
