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

type card struct {
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
	
	// Define the Map we're going to use
	cardMap := map[string]map[string][]card{}
	
	var (
		agentCount int = 0
		engineCount int = 0
		anchorCount int = 0
		conflictCount int = 0
		aspectCount int = 0
	)

	// Iterate through the file one line at a time
	for scanner.Scan() {
		// Split the line on commas
		line := strings.Split(scanner.Text(),",")

		// Check and see if we've seen the current set before, if we haven't initialize the sub-map.
		if _, ok := cardMap[line[0]]; !ok {
			cardMap[line[0]] = map[string][]card{}
		}

		// Add the card to the map
		cardMap[line[0]][line[1]] = append(cardMap[line[0]][line[1]], newCard(line[2:]))

		// Just get me some basic counts for sanity checking
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

	// Show me my counts
	fmt.Println("Number of Agents: ", agentCount)
	fmt.Println("Number of Engines: ", engineCount)
	fmt.Println("Number of Anchors: ", anchorCount)
	fmt.Println("Number of Conflicts: ", conflictCount)
	fmt.Println("Number of Aspects: ", aspectCount)

	// Validate that data is getting where I think it should be
	fmt.Println("Number of items in cardMap['Base']['Agent']: ", len(cardMap["Base"]["Agent"]))
	fmt.Println("Number of items in cardMap['SciFi']['Agent']: ", len(cardMap["SciFi"]["Agent"]))
	fmt.Println("The first item in cardMap['Base']['Agent']: ", cardMap["Base"]["Agent"][0])
	fmt.Println("The first item in cardMap['SciFi']['Engine']: ", cardMap["SciFi"]["Engine"][0])

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
func newCard(side []string) card {
	var c card
	
	if len(side) > 2 {
		// This is a 4 sided card
		c = card{side1: side[0], side2: side[1], side3: side[2], side4: side[3]}
	} else {
		// This is a 2 sided card
		c = card{side1: side[0], side2: side[1]}
	}
	
	return c
}