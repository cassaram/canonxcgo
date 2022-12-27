package canonxcgo

import (
	"bufio"
	"strings"
)

// Functions to handle parsing a HTTP response from a device

// Extracts a key-value pair from a single string with delimiters
func extractKeyValue(root string) (string, string) {
	// Find potential := / ==
	for i := 0; i < len(root)-1; i++ {
		substr := root[i : i+2]
		if substr == ":=" || substr == "==" {
			key := root[:i]
			value := root[i+2:]

			return key, value
		}
	}
	return "", ""
}

// Parses a HTTP response into key-value pairs
func readResponse(response string) map[string]string {
	result := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(response))
	for scanner.Scan() {
		line := scanner.Text()

		// Handle body header junk
		if strings.HasPrefix(line, "timestamp=") {
			continue
		}
		if strings.HasPrefix(line, "realtime=") {
			continue
		}

		// Extract key value pair
		key, val := extractKeyValue(line)
		result[key] = val
	}

	return result
}

// Utility function for handling lists of comma separated values
func parseListCSV(source string) []string {
	result := make([]string, 1)
	for _, char := range source {
		if char == ',' {
			// New string
			result = append(result, "")
		} else {
			// Append to last string
			result[len(result)-1] += string(char)
		}
	}
	return result
}
