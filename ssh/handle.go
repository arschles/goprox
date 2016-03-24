package ssh

import (
	"golang.org/crypto/ssh"
)

func handleServerConn(keyID string, chans <-chan ssh.NewChannel) {
	for newChan := range chans {
		if newChan.ChannelType() != "session" {
			newChan.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}

		ch, reqs, err := newChan.Accept()
		if err != nil {
			// handle error
			continue
		}

		go func(in <-chan *ssh.Request) {
			defer ch.Close()
			for req := range in {
				payload := cleanCommand(string(req.Payload))
				switch req.Type {
				case "exec":
					cmdName := strings.TrimLeft(payload, "'()")

					args := []string{"serv", "key-" + keyID, "--config=" + setting.CustomConf}
					cmd := exec.Command(setting.AppPath, args...)

					stdout, err := cmd.StdoutPipe()
					if err != nil {
						// handle error
						return
					}
					stderr, err := cmd.StderrPipe()
					if err != nil {
						// handle error
						return
					}
					input, err := cmd.StdinPipe()
					if err != nil {
						// handle error
						return
					}

					if err = cmd.Start(); err != nil {
						// handle error
						return
					}

					go io.Copy(input, ch)
					io.Copy(ch, stdout)
					io.Copy(ch.Stderr(), stderr)

					if err = cmd.Wait(); err != nil {
						// handle error
						return
					}

					ch.SendRequest("exit-status", false, []byte{0, 0, 0, 0})
					return
				default:
				}
			}
		}(reqs)
	}
}
