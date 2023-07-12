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
	solver     usecases.Solver
	roomKind   entities.RoomKind
	readState  byte
	startFound bool
	endFound   bool
	inout      string
}

func NewCLI() *CLI {
	return &CLI{
		builder:    usecases.NewBuilder(),
		solver:     usecases.NewSolver(),
		roomKind:   entities.Regular,
		readState:  ants,
		startFound: false,
		endFound:   false,
	}
}

func (c *CLI) Run(filename string) error {
	var err error
	if err = c.saveData(filename); err != nil {
		return fmt.Errorf("ERROR: invalid data format, %v", err)
	}
	solution, err := c.solver.Solve(c.builder.Anthill())
	if err != nil {
		return fmt.Errorf("ERROR: invalid data format, %v", err)
	}
	c.writeResult(solution)
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
			return errors.New("empty line")
		}
		if strings.Contains(c.inout, line) {
			return errors.New("duplicating input data")
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
				return errors.New("not valid ants number")
			}
			c.builder.SetAnts(num)
			c.readState = rooms
		case rooms:
			name, x, y, err := c.checkRoomData(line)
			if err != nil {
				if err = c.makeTunnel(line); err != nil {
					return err
				}
				c.readState = tunnels
				continue
			}

			if err = c.builder.CreateRoom(name, x, y, c.roomKind); err != nil {
				return err
			}

			if c.roomKind != entities.Regular {
				c.roomKind = entities.Regular
			}

		case tunnels:
			if err = c.makeTunnel(line); err != nil {
				return err
			}
		case comment:
			c.readState = prevState
		default:
			return errors.New("invalid read mode")
		}

		c.inout += line + "\n"

	}
	if !c.startFound || !c.endFound {
		return errors.New("no data start or end room found")
	}

	return nil
}

func (c *CLI) writeResult(queue entities.Queue) {
	fmt.Println(c.inout)
	for _, step := range queue {
		for i, move := range step {
			fmt.Printf("L%v-%v", move.Ant, move.Destination)
			if i < len(step)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (c *CLI) checkRoomData(line string) (string, int, int, error) {
	roomData := strings.Fields(line)
	if len(roomData) != 3 {
		return "", 0, 0, errors.New("invalid room format - wrong number of entries")
	}

	x, err := strconv.Atoi(roomData[1])
	if err != nil {
		return "", 0, 0, errors.New("invalid room format - incorrect x coord")
	}

	y, err := strconv.Atoi(roomData[2])
	if err != nil {
		return "", 0, 0, errors.New("invalid room format - incorrect y coord")
	}

	return roomData[0], x, y, nil
}

func (c *CLI) makeTunnel(line string) error {
	roomNames := strings.Split(line, "-")
	if len(roomNames) != 2 {
		return errors.New("incorrect tunnel info - not 2 rooms")
	}
	if err := c.builder.CreateTunnel(roomNames); err != nil {
		return err
	}
	return nil
}
