//go:build !local
// +build !local

package oxilib

// Default configuration for the oxilib testing will generate temporary db file and
// delete it after the test
const localTestFile = "" // check config_lacal.go for default local configuration
const useLocalTestFile = false
