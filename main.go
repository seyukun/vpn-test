package main

import (
	"fmt"
	"net"
	"sync"
	"vpn/visualizer"

	"golang.org/x/sys/unix"
)

func main() {
	tun, err := createTUN("tun0", unix.IFF_TUN|unix.IFF_MULTI_QUEUE|unix.IFF_NAPI)
	if err != nil {
		panic(err)
	}
	defer tun.Close()

	err = unix.SetNonblock(int(tun.Fd()), true)
	if err != nil {
		panic(err)
	}

	localAddr, err := net.ResolveUDPAddr("udp", ":43000")
	if err != nil {
		panic(err)
	}

	remoteAddr, err := net.ResolveUDPAddr("udp", "153.127.195.10:43000")
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		buf := make([]byte, 1500)
		for {
			n, _, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Error reading from UDP:", err)
				continue
			}
			// fmt.Printf("Received %d bytes from %s: %s\n", n, addr, string(buf[:n]))
			tun.Write(buf[:n])
		}
	}()

	wg.Add(1)
	go func() {
		buf := make([]byte, 1500)
		for {
			pollFds := []unix.PollFd{
				{Fd: int32(tun.Fd()), Events: unix.POLLIN},
			}
			n, err := unix.Poll(pollFds, 500)
			if err != nil {
				fmt.Printf("Error polling: %v\n", err)
				continue
			}
			if n == 0 {
				continue
			}

			if (pollFds[0].Revents & unix.POLLIN) != 0 {
				n, err := tun.Read(buf)
				if err != nil {
					if err == unix.EAGAIN || err == unix.EWOULDBLOCK {
					} else {
						fmt.Printf("Error reading from TUN: %v\n", err)
					}
					continue
				}
				if n > 0 {
					visualizer.IPDatagram(buf[:n])
					_, err = conn.WriteToUDP(buf[:n], remoteAddr)
					if err != nil {
						fmt.Println("Error sending to UDP:", err)
						continue
					}
				}
			}
		}
	}()

	wg.Wait()
	fmt.Println("全てのゴルーチンが終了しました。")
}
