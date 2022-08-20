[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bykovme_bslib&metric=alert_status)](https://sonarcloud.io/dashboard?id=bykovme_bslib)
[![Go Report Card](https://goreportcard.com/badge/github.com/oxipass/oxilib)](https://goreportcard.com/report/github.com/oxipass/oxilib)
[![codecov](https://codecov.io/gh/bykovme/bslib/branch/master/graph/badge.svg)](https://codecov.io/gh/bykovme/bslib)

## Oxi Lib

Oxi Lib is a library to work with encrypted personal data stored as SQLite file keeping 
items/fields safe and secure. This library is a part of the [OxiPass](https://oxipass.io) project 

*Use the following command to install it locally* 
```
go get github.com/oxipass/oxilib
```

**IMPORTANT!** This package is still work in progress and is in active development phase, 
use it at your own risk

You can use local sqlite db to test the package. Create the file config_local.go 
with the following content:
```go
//go:build local
// +build local

package oxilib

const localTestFile = "/Users/bkv/.oxipass/oxipass.sqlite" // create your own local db file for testing purposes
const useLocalTestFile = true
```
and add the tag 'local' when you run the tests for the package 
```
go test -tags local
```

Check other related packages:    
[Encryption library 'OxiCrypt'](https://github.com/oxipass/oxicrypt)  
[Mobile/Desktop App 'OxiPass'](https://github.com/oxipass/oxipass)   
[OxiPass Homepage](https://oxipass.io)   

**[Alex Bykov](https://profile.codersrank.io/user/bykovme) © 2015 - 2022**

