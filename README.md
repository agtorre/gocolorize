#Gocolorize
Gocolorize is a package that allows Go programs to provide ANSI coloring in either a stateful or stateless manner. Ideal for logging or output of cli applications.

![colored tests passing](https://raw.github.com/agtorre/gocolorize/master/screenshot/tests.png)

##Features
- Stateful and stateless ANSI coloring 
- Foreground and background colors
- Tests with 100% coverage
- Working examples
- Interface allows coloring of many different types, including complex types
- Can color multiple arguments
- Helpful paint module to simplify the basic use-case, but is slightly less efficient

##Install Gocolorize
To install

    $ go get github.com/agtorre/gocolorize

##Examples
See examples directory for examples

    $ cd examples/
    $ go run song.go
    $ go run logging.go

###Tests
Tests are another good place to see examples. In order to run tests:

    $ go test -cover
    $ cd paints/
    $ go test -cover
