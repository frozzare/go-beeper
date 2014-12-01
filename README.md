# go-beeper [![Build Status](https://travis-ci.org/frozzare/go-beeper.svg?branch=master)](https://travis-ci.org/frozzare/go-beeper)

 Make your terminal beep

 View the [docs](http://godoc.org/github.com/frozzare/go-beeper).

## Installation

```
$ go get github.com/frozzare/go-beeper
```

## Example

```go
package main

import (
	"github.com/frozzare/go-beeper"
)

func main() {
	// beep once
	beeper.Beep()

	// beep three times
	beeper.Beep(3)
}
```

Play a melody. Star is a beep and dash is the pause sign.

```go
package main

import (
	"github.com/frozzare/go-beeper"
)

func main() {
	beeper.Melody("**-**")
	// beep, beep, pause, beep, beep
}
```

# License

 MIT
