package main

import (
	"fmt"
	"go-ssh/ssh"
	"io"
	"os"
)

const (
	user     = "user"
	ip_port  = "ip_port"
	password = "Passwd"
)

func main() {
	PassWd := []ssh.AuthMethod{ssh.Password(password)}
	Conf := ssh.ClientConfig{User: user, Auth: PassWd}
	Client, err := ssh.Dial("tcp", ip_port, &Conf)
	if err != nil {
		fmt.Println(nil)
	}
	defer Client.Close()
	if session, err := Client.NewSession(); err == nil {
		defer session.Close()
		go func() {
			Buf := make([]byte, 1024)
			w, _ := session.StdinPipe()
			defer w.Close()
			File, _ := os.Open("FilePath")
			info, _ := File.Stat()
			fmt.Fprintln(w, "C0644", info.Size(), "Des_Name")
			for {
				n, err := File.Read(Buf)
				fmt.Fprint(w, string(Buf[:n]))
				if err != nil {
					if err == io.EOF {
						return
					} else {
						panic(err)
					}
				}
			}
		}()
		if err := session.Run("/usr/bin/scp -qrt /mnt"); err != nil {
			if err != nil {
				if err.Error() != "Process exited with: 1. Reason was:  ()" {
					fmt.Println(err.Error())
				}
			}
		}
	}
}