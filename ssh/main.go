package ssh

import (
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/crypto/ssh"
)

// taken from http://blog.scalingo.com/post/105010314493/writing-a-replacement-to-openssh-using-go-22
func handleChanReq(chanReq ssh.NewChannel) {
	if chanReq.ChannelType() != "session" {
		chanReq.Reject(ssh.Prohibited, "channel type is not a session")
		return
	}

	ch, reqs, err := chanReq.Accept()
	if err != nil {
		log.Println("fail to accept channel request", err)
		return
	}

	req := <-reqs
	if req.Type != "exec" {
		ch.Write([]byte("request type '" + req.Type + "' is not 'exec'\r\n"))
		ch.Close()
		return
	}

	handleExec(ch, req)
}

// handleExec filter the command which can be run.
// Payload: string: command
func handleExec(ch ssh.Channel, req *ssh.Request) {
	command := string(req.Payload)
	gitCmds := []string{"git-receive-pack", "git-upload-pack"}

	valid := false
	for _, cmd := range gitCmds {
		if strings.HasPrefix(command, cmd) {
			valid = true
		}
	}
	if !valid {
		ch.Write([]byte("command is not a GIT command\r\n"))
		ch.Close()
		return
	}

	ch.Write([]byte("well done!\r\n"))
	ch.Close()
}

// Run runs the Git SSH server
func Run(port int) error {
	config := ssh.ServerConfig{}
	portStr := fmt.Sprintf(":%d", port)
	socket, err := net.Listen("tcp", portStr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := socket.Accept()
		if err != nil {
			return err
		}

		// form a standard TCP connection to an encrypted SSH connection
		sshConn, chans, reqs, err := ssh.NewServerConn(conn, &config)
		if err != nil {
			log.Printf("Error creating new SSH conn (%s)", err)
			continue
		}
		go ssh.DiscardRequests(reqs)

		log.Println("Connection from", sshConn.RemoteAddr())
		sshConn.Close()
	}
}
