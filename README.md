# Gocfg

Gocfg is an interpolative library for configuration. It reads from files (yml, toml, json, conf, ...) and sets the corresponding parameter values on a struct. The cool part of this library is that this reading process is interpolated, in other words this means that you could use variables to search for 
environment variables and define default values if they don’t exist, all in the file. eg:

    some_key: "test"
    another_key: 1234
    sub_object:
      something: "a"
      something_else: 1234
      something_else_with_interpolation: "{SOMETHING_ELSE_WITH_INTERPOLATION:3432.8}"
    last_key: "{LAST_KEY:default value}"
    last_key_with_number: "{LAST_KEY_WITH_NUMBER:54}"

As you can see the if you set the value as `"{ENVIRONMENT_VARIABLE:default_value}"` or only `"{ENVIRONMENT_VARIABLE}"`  if you want it to panic if no value no value found for this environment variable is set.

## How is the code part?

On your code you must have a struct to be mapped by the lib like that:

	var SomeStruct = struct {
		SomeKey    string `cfg:"some_key"`
		AnotherKey int    `cfg:"another_key"`
		SubObject  struct {
			SomeThing                      string  `cfg:"something"`
			SomethingElse                  int     `cfg:"something_else"`
			SomethingElseWithInterpolation float64 `cfg:"something_else_with_interpolation"`
		} `cfg:"sub_object"`
		LastKey           string `cfg:"last_key"`
		LastKeyWithNumber uint   `cfg:"last_key_with_number"`
	}{}

Or if you want a more organized way of segregating your nested structures into different independent types:

  type SomeStruct struct {
		SomeKey           string
		AnotherKey        int
		SubObject  		  SubObjectType
		LastKey           string
		LastKeyWithNumber uint
	}
	
	type SubObjectType struct {
		SomeThing                      string
		SomethingElse                  int
		SomethingElseWithInterpolation float64
	}


> The "cfg" tag, like the json tag, acts like an alias when binding the keys in the file with your struct's parameters. They are helpful but not obligatory.

### Required

Required subtag is meant to sinalize for the lib when a field that doesn’t have a default value and your application expects some value to be setted.

    type Example struct {
        First  float64 `cfg:"required"`
        Second bool    `cfg:"SECOND,required"`
    }

## Starting the load process

To load the configurations you just need to use the function load from the lib. e.g.:

`err := gocfg.Load(&variable, gocfg.YAML, "/your/file/location.yml")`

>If you want to, you don't need to pass a file path. The lib will set by default an application.{your selected file type} on your current directory
The file types are selected as integers, but the library has an iota enum that helps you to easily set the load. Currently we have:

- YAML;
- JSON;
- ENV (this means full environment variables, not dotenv);
- TOML;

