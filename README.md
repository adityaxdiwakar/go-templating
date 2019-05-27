# go-templating
A template engine written in Golang for NetVPX.

This was written and open saourced by the [NetVPX](https://netvpx.com/) DevOps team for internal usage. This library is maintained by the team for the public.

## Install

To use this library, simply install it using:

```
go get github.com/netvpx/go-sendpulse
```

If you awnt to set it up as a submodule, use:

```
git submodule add https://github.com/netvpx/go-templating github.com/netvpx/go-templating
```

## Usage

The library is very basic but requires a strict type of usage dynamic.

#### Templating Syntax

```html
<body>
    <h1>Hello {{ firstname }} {{ lastname }}!</h1>
</body>
```

The templating uses mustaches ``{{ }}`` for templating.

- ``{{variable}}`` 
- ``{{ variable }}``

These are both valid methods for templating, you can also have a tab between the brackets and variable name, but not extra raw spaces.

#### Templating Storage

Store these files in whatever file you want to save them, as long as they're accessible. 

An example project structure:

```
src/
- main.go
- templates/
- - template1.html
- - template2.html
- - template3.html
- github.com/
- - netvpx/
- - - go-templating/ 
```

#### Templating Implementation

The builtin function ``ReplaceOccurances()`` takes two variables; prototyped as...

```go
func ReplaceOccurances(templateName string, items map[string]string) []byte
```

In order to use this, you need to provide the path to the template file as ``templateName`` which is a string, and a map in the name of ``items``.

The map can be defined as so:
```go
values := map[string]string {
    "firstname": "John",
    "lastname": "Doe",
}
```

## Example

You can use the library as so:

template.html
```html
<body>
    <h1>Hello {{ firstname }} {{ lastname }}!</h1>
</body>
```

main.go
```go
package main

import (
    "fmt"
    "github.com/netvpx/go-templating"
)

func main() {
    values := map[string]string {
        "firstname": "John",
        "lastname": "Doe",
    }
    result := templating.ReplaceOccurances(
        "template.html",
        values,
    ) //returns a byte slice, more useful

    fmt.Println(string(result))
}
```

## Support

If you need help, feel free to open a GitHub issue or email ``aditya@netvpx.com`` who is the Head of R&D at NetVPX.

## Contributing

Feel free to contribute, please follow ``gofmt`` for formatting and make a pull request. No guarantees that contributions will be accepted, thanks for your interest.