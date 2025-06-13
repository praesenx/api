package main

import (
	"bufio"
	"fmt"
	"github.com/oullin/cli/menu"
	"github.com/oullin/pkg/cli"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	panel := menu.Panel{Reader: reader}

	panel.PrintMenu()

	for {
		choice, err := panel.CaptureInput()

		if err != nil {
			cli.MakeTextColour(err.Error(), cli.Red).Println()
			continue
		}

		switch *choice {
		case 1:
			sayHello()
		case 2:
			showTime()
		case 3:
			doSomethingElse()
		case 0:
			fmt.Println(cli.Green + "Goodbye!" + cli.Reset)
			return
		default:
			fmt.Println(cli.Red, "Unknown option. Try again.", cli.Reset)
		}

		fmt.Print("\nPress Enter to continue...")

		_, _ = reader.ReadString('\n')
	}
}

func sayHello() {
	fmt.Println(cli.Green + "\nHello, world!" + cli.Reset)
}

func showTime() {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(cli.Green, "\nCurrent time is", now, cli.Reset)
}

func doSomethingElse() {
	fmt.Println(cli.Green + "\nDoing something else..." + cli.Reset)
}
