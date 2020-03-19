package main

import (
	"fmt"
)

func usage() {
	fmt.Println(`
    USAGE:

    Envsub app is a simple tool which use golang template engine to modify pointed files.
    If template file name contains .tmpl suffix it will be removed after parsing template.

      example1.conf.tmpl => example1.conf

    1. Use env variables for templates

      In template file use following placeholder {{ .ENV_VARIABLE_NAME }}
      to replace it with corresponding env variable

      EXAMPLE:
      envsub file.tmpl --envs

    2. Use extra values

      Prepare json values:

        {"rootValue":{"firstLevel":{"secondLevel":{"key1":"value1","key2":["listel1","listel2"]}}}}

      Export it as env variable:

        export BASE64_JSON_ENV_VARIABLE_NAME={"rootValue":{"firstLevel":{"secondLevel":{"key1":"value1","key2":["listel1","listel2"]}}}}

      In template file use following placeholder {{ .rootValue.firstLevel.secondLevel.key1 }}
      to replace it with corresponding json key (value1 in this case)

      EXAMPLE:
      envsub file.tmpl -e BASE64_JSON_ENV_VARIABLE_NAME

    3. Use yaml configuration file

      EXAMPLE:
      envsub file.tmpl -f config.yaml
  `)

	fmt.Printf("\n")
}
