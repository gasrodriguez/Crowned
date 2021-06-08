package server

import "os"

type Stdinout struct{}

func (Stdinout) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

func (Stdinout) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

func (Stdinout) Close() error {
	if err := os.Stdin.Close(); err != nil {
		return err
	}
	return os.Stdout.Close()
}
