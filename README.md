# serial [![Build Status](https://travis-ci.org/Nortech-ai/serial.svg?branch=master)](https://travis-ci.org/Nortech-ai/serial) [![GoDoc](https://godoc.org/github.com/Nortech-ai/serial?status.svg)](https://godoc.org/github.com/Nortech-ai/serial)
## Example
```go
package main

import (
	"log"

	"github.com/Nortech-ai/serial"
)

func main() {
	port, err := serial.Open(&serial.Config{Address: "/dev/ttyUSB0"})
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	_, err = port.Write([]byte("serial"))
	if err != nil {
		log.Fatal(err)
	}
}
```
## Testing

### Linux and Mac OS
- `socat -d -d pty,raw,echo=0 pty,raw,echo=0`
- on Mac OS, the socat command can be installed using homebrew:
	````brew install socat````

### Windows
- [Null-modem emulator](http://com0com.sourceforge.net/)
- [Terminal](https://sites.google.com/site/terminalbpp/)
