/*
Copyright © 2023 Adam Debus
*/
package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/adebus/StoryEngine/card"
)

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

	//scanner := bufio.NewScanner(inFile)
	csvReader := csv.NewReader(inFile)
	
	// Define the Map we're going to use
	cardMap := map[string]map[string][]card.Card{}
	
	var (
		agentCount int = 0
		engineCount int = 0
		anchorCount int = 0
		conflictCount int = 0
		aspectCount int = 0
	)

	// Iterate through the file one line at a time
	for {

		line, err := csvReader.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			checkErr(err, "Encountered an error reading the line")
		}

		// Check and see if we've seen the current set before, if we haven't initialize the sub-map.
		if _, ok := cardMap[line[0]]; !ok {
			cardMap[line[0]] = map[string][]card.Card{}
		}

		// Validate that the type of card is valid
		if line[1] != "Agent" && line[1] != "Engine" && line[1] != "Anchor" && line[1] != "Conflict" && line[1] != "Aspect" {
			fmt.Println("Line does not have a valid card type: ", line[1])
		}

		// Add the card to the map
		var c card.Card
		if line[4] == "" && line[5] == "" {
			c, err = card.New(line[2:4])
		} else {
			c, err = card.New(line[2:])
		}
		
		if err != nil {
			checkErr(err, "Encountered an error creating the new card")
		}
		cardMap[line[0]][line[1]] = append(cardMap[line[0]][line[1]], c)

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

	// Iterate through the maps and format each array as JSON and write the appropriate file
	for cardSet, value := range cardMap {
		for cardType, value2 := range value {

			//fmt.Println("Processing "+cardSet+" - "+cardType+"...")

			//fmt.Printf("Number if items in cardMap[\"%v\"][\"%v\"]: %v\n", cardSet, cardType, len(value2))

			outfile := cardSet+"-"+cardType+".json"

			fileData, err := json.MarshalIndent(value2, "", "\t")
			if err != nil {
				checkErr(err, "Error marshalling JSON")
			}

			err = os.WriteFile(outputDir+"/"+outfile, fileData, 0644)
			if  err != nil {
				checkErr(err, "Error writing file")
			}

			//fmt.Println("...Done")
		}
	}

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
