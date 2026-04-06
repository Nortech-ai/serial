//go:build darwin || linux || freebsd || openbsd || netbsd
// +build darwin linux freebsd openbsd netbsd

package serial

import (
	"os"
	"syscall"
	"unsafe"
)

func getModemConfig(fd int) (int, error) {
	// Do a TIOCMGET to get the current modem config
	var modemConfig int
	reterr, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TIOCMGET),
		uintptr(unsafe.Pointer(&modemConfig)),
	)
	if reterr != 0 && errno != 0 {
		return 0, os.NewSyscallError("SYS_IOCTL (TIOCMGET)", errno)
	}
	return modemConfig, nil
}

func setModemConfig(fd int, modemConfig int) error {
	// Do a TIOCMSET to set the modem config
	reterr, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		uintptr(syscall.TIOCMSET),
		uintptr(unsafe.Pointer(&modemConfig)),
	)
	if reterr != 0 && errno != 0 {
		return os.NewSyscallError("SYS_IOCTL (TIOCMSET)", errno)
	}
	return nil
}
