// Package models struct
package models

// ConfigLogging shema of configparams from config file
type ConfigLogging struct {
	Directory string
	Filename  string
}

// ConfigAPI shema API from config file
type ConfigAPI struct {
	APIPort string
}
