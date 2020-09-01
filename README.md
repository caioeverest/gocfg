# Gocfg

Gocfg is an interpolative library for configuration. It reads from files (yml, toml, json, conf, ...) and sets the corresponding parameter values on a struct. The cool part of this library is that this reading process is interpolated, in other words this means that you could use variables to search for environment variables and define default values if they don’t exist, all in the file. eg:

...code...

## How is the code part?

On your code you must have a struct to be mapped by the lib like that:

    ```
    ...code...
    ```

Or if you want a more organized way of segregating your nested structures into different independent types:

    ```
    ...code...
    ```

> The "cfg" tag, like the json tag, acts like an alias when binding the keys in the file with your struct's parameters. They are helpful but not obligatory.

### Required

Required subtag is meant to sinalize for the lib when a field that doesn’t have a default value and your application expects some value to be setted.

    ```
    ...code...
    ```

## Starting the load process

To load the configurations you just need to use the function load from the lib. e.g.:

`err := gocfg.Load(&variable, gocfg.YAML, "/your/file/location.yml")`

>If you want to, you don't need to pass a file path. The lib will set by default an application.{your selected file type} on your current directory
The file types are selected as integers, but the library has an iota enum that helps you to easily set the load. Currently we have:

- YAML;
- JSON;
- ENV (this means full environment variables, not dotenv);
- TOML;

