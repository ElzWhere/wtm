/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net"
	"sync"
	"time"
)

// checkPort function that checks if a port is open
func checkPort(ip string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

// scanPorts function that scans ports
func scanPorts(ip string, startPort, endPort int, timeout time.Duration) {
	var wg sync.WaitGroup
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			if checkPort(ip, p, timeout) {
				fmt.Printf("Port %d is open\n", p)
			}
		}(port)
	}
	wg.Wait()
}

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans for open ports on a host",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		count := CountIpAddresses(fileName)
		var ip []string
		local1, local2 := "127.0.0.1", "0.0.0.1"
		// only gets ip addresses
		for key, _ := range count {
			if key == local1 || key == local2 {
				continue
			} else {
				ip = append(ip, key)
			}
		}

		for _, i := range ip {
			fmt.Printf("Scanning ports on host: %s \n", i)
			scanPorts(i, 1, 1024, 500*time.Millisecond)

		}
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)

}
