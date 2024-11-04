# Env ![Build](https://github.com/xyproto/env/workflows/Build/badge.svg) [![GoDoc](https://godoc.org/github.com/xyproto/env?status.svg)](http://godoc.org/github.com/xyproto/env) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/env)](https://goreportcard.com/report/github.com/xyproto/env)

Makes fetching and interpreting environment variables easy and safe.

Being able to provide default values when retrieving environment variables often makes program logic shorter and more readable.

## Functions

### Keys

Keys returns the all the environment variable names as a sorted string slice

`func Keys() []string`

### Map

Map returns the current environment variables as a map from name to value

`func Map() map[string]string`

### Str

`func Str(name string, optionalDefault ...string) string`

`Str` does the same as `os.Getenv`, but allows the user to provide a default value (optional).
Only the first optional value is used, if the environment variable value is empty or not set.

### Bool

`func Bool(envName string) bool`

`Bool` returns the bool value of the given environment variable name. Returns `false` if it is not declared or empty.

### Int

`func Int(envName string, defaultValue int) int`

`Int` returns the number stored in the environment variable, or the given default value.

### Environ

`func Environ() []string`

`Environ` returns the same as `os.Environ()`; a slice of strings with `=` between the key and value, except that the contents are cached by default.

### True

`func True(s string) bool`

Checks if the given string is a positive string that should be interpreted as `true`, such as `"yes"`, `"1"` or `"true"`.

### False

`func False(s string) bool`

Checks if the given string is a negative string that should be interpreted as `false`, such as `"no"`, `"0"` or `"false"`.

### AsBool

`func AsBool(s string) bool`

`AsBool` can be used to interpret a string value as either `true` or `false`. Examples of `true` values are "yes" and "1".

### AsBoolSimple

`func AsBoolSimple(s string) bool`

Checks if the given string is `1`.

### Has

`func Has(s string) bool`

`Has` return true if the given environment variable name is non-empty.

## No

`func No(s string) bool`

`No` returns true if the given environment variable name is not set or empty.

## CurrentUser

`func CurrentUser() string`

Returns the current LOGNAME or USER env name, or an empty string.

### DurationSeconds

Takes a default int64 value, for the number of seconds, interprets the environment variable as the number of seconds and returns a `time.Duration`.

### DurationMinutes

Takes a default int64 value, for the number of minutes, interprets the environment variable as the number of seconds and returns a `time.Duration`.

### DurationHours

Takes a default int64 value, for the number of hours, interprets the environment variable as the number of seconds and returns a `time.Duration`.

### Contains

Checks if the given environment variable contains the given string.

### Is

Checks if the given environment variable is the given value, with leading and trailing spaces trimmed before comparing both values.

### HomeDir

Returns the home directory of the current user, or `/tmp` if it is not available. `/home/$LOGNAME` or `/home/$USER` are returned if `$HOME` is not set.

### ExpandUser

Replaces `~` or `$HOME` at the start of a string with the home directory of the current user.

### File

Does the same as the `Str` function, but replaces a leading `~` or `$HOME` with the home directory of the current user.

### Dir

Does the same as the `Str` function, but replaces a leading `~` or `$HOME` with the home directory of the current user.

### Path

Returns the current `$PATH` as a slice of strings.

### EtcEnvironment

Lookup `KEY=VALUE` lines in `/etc/environment` and return the value, if found.

## WaylandSession

Returns true of `XDG_SESSION_TYPE` is `wayland` or if `DESKTOP_SESSION` contains `wayland`.

## XSession

Returns true if `DISPLAY` is set.

## OnlyXSession

Returns true if `DISPLAY` is set and `WaylandSession()` returns false.

## XOrWaylandSession

Returns true if `DISPLAY` is set or `WaylandSession()` returns true.

## Set

Takes the name of the environment variable to be set and an optional string value. If no string is given, `1` is used as the value.

## Unset

Takes the name of the environment variable to unset.

### Other functions

There are also: `Float64`, `Float32`, `Uint64`, `Uint32`, `Uint16`, `Uint8`, `Uint`, `Int64`, `Int32`, `Int16` and `Int8` functions available.


## Example

```go
package main

import (
    "fmt"
    "github.com/xyproto/env/v2"
)

func main() {
    fmt.Println(env.DurationSeconds("REQUEST_TIMEOUT", 1800))
}
```

Running the above program like this: `REQUEST_TIMEOUT=1200 ./main`, outputs:

    20m0s

## Cache

* Cache is enabled by default from version 2.0.0 and beyond. The first time a variable is read, all environment variables are read into the cache, to avoid unnecessary system calls.
* If `Unload()` is called, then the cache is disabled and cleared and `os.Getenv` and `os.Setenv` are used directly.
* If the cache functionality is disabled, then `Load()` can be used to enable it and to read the environment variables from the system again.
* The `Set` and `Unset` functions can be used to both set values with `os.Setenv` and also modify the cache, so that an additional call to `Load()` is not needed.

## General info

* Version: 2.5.2
* License: BSD-3
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
