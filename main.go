package main

import (
	"os"

	flag "github.com/spf13/pflag"
)

type values interface{}

var (
	helpFlag       bool
	envsFlag       bool
	valuesEnvName  string
	valuesFileName string
	outputFileName string
	envValues      map[string]interface{}
	fileValues     map[string]interface{}
)

func init() {
	flag.BoolVarP(&helpFlag, "help", "h", false, "Help message for envsub")
	flag.BoolVar(&envsFlag, "envs", true, "Use environment variables as template values")
	flag.StringVarP(&valuesEnvName, "json-env", "e", "", "Env variable name containing json values")
	flag.StringVarP(&valuesFileName, "file-values", "f", "", "Path to the file containing values (yaml)")
	flag.StringVarP(&outputFileName, "output", "o", "", "Output file name (!!ONLY WHEN PARSING SINGLE FILE!!)")
}

func main() {

	flag.Parse()

	if helpFlag || flag.NArg() == 0 {
		usage()
		flag.PrintDefaults()
		os.Exit(1)
	}

	filepaths := prepareFilepaths(flag.Args())

	values := map[string]interface{}{}
	envs := prepareEnvs()

	if envsFlag {
		values = mergeValues(values, convert(envs))
	}

	if valuesEnvName != "" {
		envValues = prepareTemplateValues(envs[valuesEnvName])
		values = mergeValues(values, envValues)
	}

	if valuesFileName != "" {
		fileValues = getConf(valuesFileName)
		values = mergeValues(values, fileValues)
	}

	output := ""

	if len(filepaths) == 1 && outputFileName != "" {
		output = outputFileName
	}

	for _, path := range filepaths {
		parseFile(path, values, output)
	}
}
