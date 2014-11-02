package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func read(in *os.File) (string, error) {
	if in == nil {
		in = os.Stdin
	}
	reader := bufio.NewReader(in)
	fmt.Print(">>> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return text, nil
}

func (e *env) initFile() error {
	if _, err := os.Stat(e.TmpPath); err == nil {
		return nil
	}
	f, err := os.OpenFile(e.TmpPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond)

	f.WriteString("package main\n")
	f.Sync()
	f.Close()

	return nil
}

func (e *env) write(content string) error {
	f, err := os.OpenFile(e.TmpPath, os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	time.Sleep(time.Microsecond)

	f.WriteString(content)
	f.Sync()
	f.Close()

	return nil
}

func (e *env) shell() {
	if err := e.initFile(); err != nil {
		fmt.Printf("[error] %v", err)
		return
	}

	p := parser{}
	for {
		text, err := read(nil)
		if err != nil {
			cleanDirs(e.BldDir)
			break
		}
		p.parserImport(text)
		if err := e.write(text); err != nil {
			fmt.Printf("[error] %v", err)
			break
		}
	}
}