package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/scgolang/osc"
	"gopkg.in/ini.v1"
)

const (
	defaultX32IP   = "127.0.0.1"             // Default x32 IP
	defaultPort    = 10023                   // Default port for Behringer X32
	defaultChannel = 37                      // Default channel number
	configPath     = "/etc/default/x32-mute" // Config file
)

func main() {
	// Check if an argument ("yes" or "no") is provided
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run client.go [yes|no]")
	}

	// Parse the argument to determine whether to send "yes" or "no"
	state := os.Args[1]
	var oscValue osc.Int

	switch state {
	case "yes":
		oscValue = 0 // Mute (i.e. x32-mute yes) (0 = channel off)
	case "no":
		oscValue = 1 // Unmute (i.e. x32-mute no) (1 = channel on)
	default:
		log.Fatalf("Invalid argument: %s. Use 'yes' or 'no'.", state)
	}

	// Load the configuration from the config file if available
	x32IP, channel := loadConfig()

	// Resolve the server address
	addr := fmt.Sprintf("%s:%d", x32IP, defaultPort)
	raddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	// Setup the OSC client
	client, err := osc.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatalf("Failed to dial UDP: %v", err)
	}
	defer client.Close()

	// Create the OSC message for the channel mute state
	msg := osc.Message{
		Address: fmt.Sprintf("/ch/%02d/mix/on", channel),
		Arguments: osc.Arguments{
			oscValue, // 1 for unmute, 0 for mute
		},
	}

	// Send the OSC message to the server
	if err := client.Send(msg); err != nil {
		log.Fatalf("Failed to send OSC message: %v", err)
	}

	fmt.Printf("Successfully sent %s command to channel %d on server %s.\n", state, channel, x32IP)
}

// loadConfig attempts to load the config file at ~/.x32-muter.conf
// and returns the server IP and channel. If the config file does not
// exist or does not specify values, it returns default values.
func loadConfig() (string, int) {
	// Set defaults
	x32IP := defaultX32IP
	channel := defaultChannel

	// Load the config file if it exists
	cfg, err := ini.Load(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Config file not found, using defaults.")
		} else {
			log.Printf("Error loading config file: %v", err)
		}
		return x32IP, channel
	}

	// Read the server IP from the [server] section
	if cfg.Section("x32").HasKey("ip") {
		x32IP = cfg.Section("x32").Key("ip").String()
	}

	// Read the channel number from the [channel] section
	if cfg.Section("channel").HasKey("number") {
		if ch, err := strconv.Atoi(cfg.Section("channel").Key("number").String()); err == nil {
			channel = ch
		} else {
			log.Printf("Invalid channel value in config: %v", err)
		}
	}

	return x32IP, channel
}
