package CLI

import (
	"bufio"
	"fmt"
	"os"

	"lem-in/internal/usecases"
)

type CLI struct {
	builder usecases.Builder
	inout   string
}

func NewCLI() *CLI {
	builder := usecases.NewBuilder()
	return &CLI{builder: builder}
}

func (c *CLI) SaveData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScan := bufio.NewScanner(file)
	for fileScan.Scan() {
		line := fileScan.Text()
		// c.manageinout(line)
		c.inout += line + "\n"
	}
	fmt.Println(c.inout)
	return nil
}
