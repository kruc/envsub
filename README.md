# envsub

### What is envsub?
Envsub is an application which allows you parse templated files using values from env variables or yaml file.

### How it works?
Envsub use go template engine to parse pointed files with values from:
- env variables
- env variable contains values in JSON format
- YAML file
If template file name contains .tmpl suffix it will be removed after parsing template.

```
example1.conf.tmpl => example1.conf
```


---
You can also provide `--output`, `-o` flag to set output filename
##### CAVEAT: works only when parsing single file!

---
### Installation

MacOS Homebrew:
```
brew install kruc/homebrew-tap/envsub
```

Linux
```
 wget -qO- https://github.com/kruc/envsub/releases/latest/download/envsub_linux.tar.gz | tar -xvz -C /usr/local/bin && chmod +x /usr/local/bin/envsub
```

### Usage:

1. Use env variables for templates

      In template file use following placeholder `{{ .ENV_VARIABLE_NAME }}`
      to replace it with corresponding env variable

        envsub file.tmpl --envs

2. Use extra values

      Prepare json values:

        {"rootValue":{"firstLevel":{"secondLevel":{"key1":"value1","key2":["listel1","listel2"]}}}}

      Export it as env variable:

        export JSON_ENV_VARIABLE_NAME={"rootValue":{"firstLevel":{"secondLevel":{"key1":"value1","key2":["listel1","listel2"]}}}}

      In template file use following placeholder {{.rootValue.firstLevel.secondLevel.key1}}
      to replace it with corresponding json key (value1 in this case)

        envsub file.tmpl -e JSON_ENV_VARIABLE_NAME

3. Use yaml configuration file

    config.yaml:
    ```yaml
      example1: value1
      example2: test-user
      example3:
        - val1
        - val2
    ```

    file.tmpl

         1. Simple example:
         {{ .example1 }}

    After execute:

        envsub file.tmpl -f config.yaml

    We will get `file` contains

        1. Simple example:
        value1

4. Use glob patterns to specify template filenames

      ```
      envsub file.json tpl/one.yaml wild/*.tmpl --envs
      ```


#### Check `examples` directory for more use cases
