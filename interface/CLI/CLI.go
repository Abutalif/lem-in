package CLI

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/internal/entities"
	"lem-in/internal/usecases"
)

const (
	ants byte = iota
	rooms
	tunnels
)

type CLI struct {
	builder   usecases.Builder
	readState byte
	inout     string
}

func NewCLI() *CLI {
	builder := usecases.NewBuilder()
	return &CLI{
		builder:   builder,
		readState: ants,
	}
}

func (c *CLI) Run(filename string) error {
	var err error
	if err = c.readData(filename); err != nil {
		return err
	}

	// TODO
	if err = c.solve(); err != nil {
		return err
	}
	if err = c.writeResult(); err != nil {
		return err
	}

	return nil
}

// ***HERE***
func (c *CLI) readData(filename string) error {
	var roomKind entities.Kind

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScan := bufio.NewScanner(file)
	for fileScan.Scan() {
		line := fileScan.Text()

		switch c.readState {
		case ants:
			num, err := strconv.Atoi(line)
			if err != nil || num < 1 {
				return errors.New("ERROR: not valid ants number")
			}
			c.builder.SetAnts(uint(num))
			c.readState = rooms
		case rooms:
			if strings.Contains(c.inout, line) {
				return errors.New("ERROR: repeated rooms")
			}

			// error: creates only regular room. How to create start/end?
			if err = c.builder.CreateRoom(line, roomKind); err != nil {
				return err
			}
		case tunnels:
		default:
			return fmt.Errorf("ERROR: invalid read mode")
		}
		c.inout += line + "\n"
	}

	return nil
}

// TODO
func (c *CLI) solve() error {
	return nil
}

// TODO
func (c *CLI) writeResult() error {
	return nil
}
