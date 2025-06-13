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
			fmt.Println(menu.ColorGreen + "Goodbye!" + menu.ColorReset)
			return
		default:
			fmt.Println(menu.ColorRed, "Unknown option. Try again.", menu.ColorReset)
		}

		fmt.Print("\nPress Enter to continue...")

		_, _ = reader.ReadString('\n')
	}
}

func sayHello() {
	fmt.Println(menu.ColorGreen + "\nHello, world!" + menu.ColorReset)
}

func showTime() {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(menu.ColorGreen, "\nCurrent time is", now, menu.ColorReset)
}

func doSomethingElse() {
	fmt.Println(menu.ColorGreen + "\nDoing something else..." + menu.ColorReset)
}
