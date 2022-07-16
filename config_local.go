//go:build local
// +build local

package oxilib

// use your own local sqlite db file for testing purposes
// set the path to the file in the constant localTestFile
// add the build tag 'local' when building the package for testing
// like 'go test -tags local'
const localTestFile = "/Users/bkv/.oxipass/oxipass.sqlite"
const useLocalTestFile = true
