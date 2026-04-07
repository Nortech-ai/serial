/*
Package serial provides a cross-platform serial reader and writer.
*/
package serial

import (
	"errors"
	"io"
	"time"
)

var (
	// ErrTimeout is occurred when timing out.
	ErrTimeout = errors.New("serial: timeout")
	// ErrUnsupported is thrown when the configuration is unsupported
	ErrUnsupported = errors.New("serial: unsupported")
)

// Config is common configuration for serial port.
type Config struct {
	// Device path (/dev/ttyS0)
	Address string
	// Baud rate (default 19200)
	BaudRate int
	// Data bits: 5, 6, 7 or 8 (default 8)
	DataBits int
	// Stop bits: 1 or 2 (default 1)
	StopBits int
	// Parity: N - None, E - Even, O - Odd (default E)
	// (The use of no parity requires 2 stop bits.)
	Parity string
	// Read (Write) timeout.
	Timeout time.Duration
	// Configuration related to RS485
	RS485 RS485Config
	// Modem bits configurations
	Modem ModemConfig
}

// Enumeratess the possible pin states for various modem pins
type PinConfiguration int

const (
	// Ignores any changes to this pin
	PinConfigurationIgnored PinConfiguration = 0
	// Force disables this pin
	PinConfigurationDisabled PinConfiguration = 1
	// Force enables this pin
	PinConfigurationEnabled PinConfiguration = 2
)

// Platform independent Modem configurations, these allow us to set flags such as DTR/RTS/etc..
type ModemConfig struct {
	DTR PinConfiguration
	RTS PinConfiguration
}

// platform independent RS485 config. Thie structure is ignored unless Enable is true.
type RS485Config struct {
	// Enable RS485 support
	Enabled bool
	// Delay RTS prior to send
	DelayRtsBeforeSend time.Duration
	// Delay RTS after send
	DelayRtsAfterSend time.Duration
	// Set RTS high during send
	RtsHighDuringSend bool
	// Set RTS high after send
	RtsHighAfterSend bool
	// Rx during Tx
	RxDuringTx bool
}

// Port is the interface for controlling serial port.
type Port interface {
	io.ReadWriteCloser
	// Connect connects to the serial port.
	Open(*Config) error
}

// Open opens a serial port.
func Open(c *Config) (p Port, err error) {
	p = New()
	err = p.Open(c)
	return
}
