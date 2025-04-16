package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	TUNSETIFF = 0x400454ca
	IFF_TUN   = 0x0001
	IFF_NO_PI = 0x1000
)

func createTUN(ifName string) (*os.File, error) {
	// /dev/net/tun を読み書きモードでオープン
	f, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to open /dev/net/tun: %v", err)
	}

	// ifreq 構造体のサイズはシステムによって異なりますが、
	// 一般的には 40 バイト程度確保しておきます。
	var ifr [40]byte

	// インターフェース名をコピー（null終端されることに注意）
	copy(ifr[:], ifName)

	// ifreq 構造体のオフセット 16 バイト目からフラグが設定される (C構造体のレイアウトに準拠)
	*(*uint16)(unsafe.Pointer(&ifr[16])) = IFF_TUN | IFF_NO_PI

	// ioctl で TUNSETIFF を実行し、TUN デバイスを生成
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(TUNSETIFF), uintptr(unsafe.Pointer(&ifr[0])))
	if errno != 0 {
		return nil, fmt.Errorf("ioctl TUNSETIFF failed: %v", errno)
	}

	err = syscall.SetNonblock(int(f.Fd()), true)
	if err != nil {
		f.Close()
		panic(err)
	}

	return f, nil
}
