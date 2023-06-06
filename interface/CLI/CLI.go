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
	comment byte = iota
	ants
	rooms
	tunnels
)

type CLI struct {
	builder    usecases.Builder
	pathfinder usecases.Pathfinder
	roomKind   entities.RoomKind
	readState  byte
	startFound bool
	endFound   bool
	inout      string
}

func NewCLI() *CLI {
	return &CLI{
		builder:    usecases.NewBuilder(),
		pathfinder: usecases.NewPathfinder(),
		roomKind:   entities.Regular,
		readState:  ants,
		startFound: false,
		endFound:   false,
	}
}

func (c *CLI) Run(filename string) error {
	var err error
	if err = c.saveData(filename); err != nil {
		return err
	}
	c.builder.Anthill().Show()
	if err = c.solve(); err != nil {
		return err
	}
	if err = c.writeResult(); err != nil {
		return err
	}

	return nil
}

func (c *CLI) saveData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScan := bufio.NewScanner(file)
	var prevState byte
	for fileScan.Scan() {
		line := fileScan.Text()
		if len(line) == 0 {
			return errors.New("ERROR: empty entry line")
		}
		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				c.roomKind = entities.Start
				c.startFound = true
			} else if line == "##end" {
				c.roomKind = entities.End
				c.endFound = true
			}
			prevState = c.readState
			c.readState = comment
		}

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

			if len(strings.Split(line, " ")) != 3 {
				c.readState = tunnels
				c.builder.CreateTunnel(line)
				continue
			}

			if err = c.builder.CreateRoom(line, c.roomKind); err != nil {
				return err
			}

			if c.roomKind != entities.Regular {
				c.roomKind = entities.Regular
			}
			fmt.Printf("writing room: %v\n", line)
		case tunnels:
			if err = c.builder.CreateTunnel(line); err != nil {
				return err
			}
			fmt.Printf("writing tunnel: %v\n", line)
		case comment:
			c.readState = prevState
		default:
			return fmt.Errorf("ERROR: invalid read mode")
		}

		c.inout += line + "\n"

	}
	if !c.startFound || !c.endFound {
		return errors.New("ERROR: no data start or end room found")
	}

	return nil
}

func (c *CLI) solve() error {
	paths := c.pathfinder.Find(c.builder.Anthill())
	if len(paths) < 1 {
		return errors.New("ERROR: no path found")
	}
	for _, p := range paths {
		p.PrintList()
	}
	return nil
}

// TODO
func (c *CLI) writeResult() error {
	// first fmt.Println(c.inout)
	// the data will be recieved in the format that
	return nil
}
