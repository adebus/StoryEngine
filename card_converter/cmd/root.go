/*
Copyright © 2023 Adam Debus
*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type cardAgent struct {
	side1 string
	side2 string
	side3 string
	side4 string
}

type cardEngine struct {
	side1 string
	side2 string
}

type cardAnchor struct {
	side1 string
	side2 string
	side3 string
	side4 string
}

type cardConflict struct {
	side1 string
	side2 string
}

type cardAspect struct {
	side1 string
	side2 string
	side3 string
	side4 string
}

var (
	// Input file
	inputFile string
	// Output Directory
	outputDir string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "card_converter",
	Short: "Convert the StoryEngine card CSV to JSON",
	Long: `A utility application designed to take the specified input file and
	parse it into multiple appropriate JSON files`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: rootRun,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func checkErr(e error, msg string) {
	if e != nil {
		log.Println(msg+": "+e.Error())
		os.Exit(1)
	}
}

func rootRun(cmd *cobra.Command, args []string) {
	inFile, err := os.Open(inputFile)
	checkErr(err, "Unable to open the input file")
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	
	var (
		agentCount int = 0
		engineCount int = 0
		anchorCount int = 0
		conflictCount int = 0
		aspectCount int = 0
	)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(),",")
		switch cardType := line[1]; cardType {
		case "Agent":
			agentCount++
		case "Engine":
			engineCount++
		case "Anchor":
			anchorCount++
		case "Conflict":
			conflictCount++
		case "Aspect":
			aspectCount++
		default:
			fmt.Println("Line doesn't have a card type: ", cardType)
		}
	}

	fmt.Println("Number of Agents: ", agentCount)
	fmt.Println("Number of Engines: ", engineCount)
	fmt.Println("Number of Anchors: ", anchorCount)
	fmt.Println("Number of Conflicts: ", conflictCount)
	fmt.Println("Number of Aspects: ", aspectCount)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.card_converter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&inputFile, "inputfile", "i", "", "The CSV to convert")
	rootCmd.Flags().StringVarP(&outputDir, "outputdir", "o", "", "The directory to write the JSON files")
	rootCmd.MarkFlagRequired("inputfile")
	rootCmd.MarkFlagRequired("outputdir")
}

// Create a new instance of the cardAgent struct
func newAgent(s1 string, s2 string, s3 string, s4 string) *cardAgent {
	c := cardAgent{side1: s1, side2: s2, side3: s3, side4: s4}
	return &c
}

// Create a new instance of the cardEngine struct
func newEngine(s1 string, s2 string) *cardEngine {
	c := cardEngine{side1: s1, side2: s2}
	return &c
}

// Create a new instance of the cardAnchor struct
func newAnchor(s1 string, s2 string, s3 string, s4 string) *cardAnchor {
	c := cardAnchor{side1: s1, side2: s2, side3: s3, side4: s4}
	return &c
}

// Create a new instance of the cardConflict struct
func newConflict(s1 string, s2 string) *cardConflict {
	c := cardConflict{side1: s1, side2: s2}
	return &c
}

// Create a new instance of the cardAspect struct
func newAspect(s1 string, s2 string, s3 string, s4 string) *cardAspect {
	c := cardAspect{side1: s1, side2: s2, side3: s3, side4: s4}
	return &c
}