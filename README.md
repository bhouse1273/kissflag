# kissflag
KISS flag helper to provide minimal support for overriding default flags with environmental variables.

If you need to have a configuration file or need to pull configuration values from etcd, etc., then Viper (github.com/spf13/viper) is the package you are looking for.  However, if all you want to do is override flag defaults with environmental variables (if they exist), then kissflag is all you need.

## Usage

Create your CLI application using Cobra, defining whatever configuration values you require.  

1) In initConfig() call kissflag.SetPrefix() to assign your environment variable prefix (good to have to avoid name collisions)
2) In either rootCmd or a subcommand, explicitly call kissflag.BindEVar() for each configuration value that you want set via environment variable
3) If you have a secret store (like Kubernetes) that returns secret values as Base64-encoded strings, call kissflag.DecodeBase64() to decode the value 

## API Reference
---
### SetPrefix

`SetPrefix(prefix string)`

#### Arguments:

<table border=2>
    <tr>
        <td>No.</td>
        <td>Name</td>
        <td>Type</td>
        <td>Comment</td>
    </tr>
    <tr>
        <td>1.</td>
        <td>prefix</td>
        <td>string</td>
        <td>The value of your variable prefix string.  It is recommended that this end with an underscore.</td>
    </tr>
</table>

#### Returns: 

nothing

#### Discussion: 

When creating a containerized micro-service, it is a good practice to have a service-specific prefix for your container's environmental variables defined. This allows orchestration tools to deploy containers without the possibility of environmental name collisions between containers.  

#### Example:

`kissflag.SetPrefix("MYPREFIX_")`

---
### BindEVar

`BindEVar(tag string, target interface{}) error`

#### Arguments:

<table border=2>
    <tr>
        <td>No.</td>
        <td>Name</td>
        <td>Type</td>
        <td>Comment</td>
    </tr>
    <tr>
        <td>1.</td> 
        <td>tag</td>
        <td>string</td> 
        <td>The name of the environmental variable to bind. This will be concatenated with the configured prefix and upper-cased. Thus, given a prefix FOO_ and a tag bar, BindEVar() will look for the FOO_BAR environmental variable.</td>
    </tr>
    <tr>
        <td>2.</td>
        <td>target</td>
        <td>interface{}</td>
        <td>A pointer to a variable that shall receive the value of the environment variable specified by tag.  Supported types are *string, *bool, *int, *int32, *int64, *float32 and *float64.</td>
    </tr>
</table>

#### Returns: 

nil error if successful  

Otherwise, error conditions are: 
* If tag is empty, returns a "tag may not be empty" error
* If target is nil, returns a "target may not be nil" error
* If value of the environmental variable cannot be converted, returns the error returned by the corresponding conversion function (from standard packages)
* If the type of the target is not recognized, returns a "unsupported target type" error 

#### Discussion: 

The BindEVar() function will test to see if the specified tag exists as an environment variable (using the prefix and uppercase naming convention).  If the variable exists, then the value referenced by the specified target is assigned the value of the environment variable.  Empty values are supported.

#### Example:

`if err := kissflag.BindEVar("port", &config.Listen); err != nil {`<br>
&nbsp;&nbsp;&nbsp;`   log.Fatalln("Config error:", err)`<br>
`}`<br>

---
### DecodeBase64

`DecodeBase64(value string, target *string, size int) error`

#### Arguments:

<table border=2>
    <tr>
        <td>No.</td>
        <td>Name</td>
        <td>Type</td>
        <td>Comment</td>
    </tr>
    <tr>
        <td>1.</td>
        <td>value</td>
        <td>string</td>
        <td>The potentially base64-encoded value of your configuration variable.</td>
    </tr>
    <tr>
        <td>2.</td>
        <td>target</td>
        <td>*string</td>
        <td>A pointer to the configuration variable you wish to receive the decoded value.</td>
    </tr>
    <tr>
        <td>3.</td>
        <td>size</td>
        <td>int</td>
        <td>The size to expect for the decoded value (as reported by the len() function), or zero.</td>
    </tr>
</table>

#### Returns:

nil error if successful

Otherwise

* if the len() of the decoded value is unequal to size, returns a "target size mismatch" error
* if the value cannot be decoded, returns the error reported by base64.StdEncoding.DecodeString()

#### Discussion:

Depending on the execution environment, some secrets might be base64 encoded or not.  To simplify this case, DecodeBase64 tries to decode the value.  If this is successful, then it compares the len() of the resulting string to the size argument's value.  If they match, then the decode was a success.

If the size argument is passed a zero value, then the len() test is not performed.  This is not recommended, since there may be false positives for non-encoded values that are nonetheless legal base64 strings.  

#### Example:

`// Test key for base64 and use decoded value if no error`<br>
`if err := kissflag.DecodeBase64(config.DarkKey, &config.DarkKey, 32); err != nil {`<br>
&nbsp;&nbsp;`log.Fatalln(err)`<br>
`}`<br>

Copyright 2018 William J House

