package main

import (
	"fmt"
	"os"
	"unsafe"

	"golang.org/x/sys/unix"
)

func createTUN(ifname string, flags uint16) (*os.File, error) {
	cloneSrc := "/dev/net/tun"
	fd, err := os.OpenFile(cloneSrc, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	ifr, err := unix.NewIfreq(ifname)
	if err != nil {
		return nil, err
	}
	ifr.SetUint16(flags)

	_, _, errno := unix.Syscall(unix.SYS_IOCTL, fd.Fd(), uintptr(unix.TUNSETIFF), uintptr(unsafe.Pointer(ifr)))
	if errno != 0 {
		return nil, fmt.Errorf("ioctl TUNSETIFF failed: %v", errno)
	}

	return fd, nil
}
