/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

var fileName = "/var/log/auth.log"

// readLogFile function that reads from a file and returns the content
func readLogFile(f string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

// getIpAddresses function that returns a slice of strings that matches regex
func getIpAddresses(f string) []string {
	logs := readLogFile(fileName)
	ipRegex := `\b(?:\d{1,3}\.){3}\d{1,3}\b`
	re := regexp.MustCompile(ipRegex)
	matches := re.FindAllString(logs, -1)
	//      fmt.Printf("%T\n", matches)
	return matches
}

// countIpAddresses function that returns a map of strings and int that counts the number of times an ip address appears
func CountIpAddresses(f string) map[string]int {
	ipAddresses := getIpAddresses(fileName)
	ipCount := make(map[string]int)
	for _, ips := range ipAddresses {
		_, exists := ipCount[ips]
		if exists {
			ipCount[ips] += 1
		} else {
			ipCount[ips] = 1
		}
	}
	//for key, value := range ipCount {
	//      fmt.Printf("%s: %d\n", key, value)
	//}

	return ipCount
}

// countCmd represents the count command
var CountCmd = &cobra.Command{
	Use:   "count",
	Short: "Gets a count of how many times a host appears in the logs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		// Logic Here

		count := CountIpAddresses(fileName)
		for key, value := range count {
			fmt.Printf("%s: %d\n", key, value)
		}
	},
}

func init() {
	/*	CountCmd.Flags().StringVarP(&fileName, "file", "f", "", "File to read from")

		if err := CountCmd.MarkFlagRequired("file"); err != nil {
			log.Fatal(err)
		} */
	rootCmd.AddCommand(CountCmd)

}
