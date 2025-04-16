package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
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
		buf := make([]byte, 1024)
		for {
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Error reading from UDP:", err)
				continue
			}
			fmt.Printf("Received %d bytes from %s: %s\n", n, addr, string(buf[:n]))
		}
	}()

	wg.Add(1)
	go func() {
		for {
			var input string
			fmt.Print("送信するメッセージを入力: ")
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println("入力エラー:", err)
				continue
			}
			_, err = conn.WriteToUDP([]byte(input), remoteAddr)
			if err != nil {
				fmt.Println("送信中にエラー:", err)
			} else {
				fmt.Println("送信完了！")
			}
		}
	}()

	wg.Wait()
	fmt.Println("全てのゴルーチンが終了しました。")
}
