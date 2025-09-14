package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
)

var a, b float64
var mode string
var separatingLine = "\n---------------------------------------------------------\n"

func clearTerminal() {
	cmd := exec.Command("clear") // Для Linux/macOS
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // Для Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func modeSelection() {
	fmt.Scanln(&mode)
	switch mode {
	case "-go":
		calculatorSaveMode()
	case "-exit":
		clearTerminal()
		os.Exit(0)
	case "-clear":
		clearTerminal()
	case "-help":
		clearTerminal()
		helpMode()
	default:
		println("\nВведите мод или выйдите...\n")
	}
}

func helpMode() {
	fmt.Print("Пока только три команды:\n'-go' - для запуска калькулятора\n'-clear' - для очистки терминала\n'-exit' - для выхода из программы\n")
}

func calculatorSaveMode() {

	fmt.Print(separatingLine, "\nВведите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	sum := a + b
	subtract := a - b
	multiply := a * b
	divide := a / b
	squared1 := a * a
	squared2 := b * b
	root1 := math.Sqrt(a)
	root2 := math.Sqrt(b)

	fmt.Println(separatingLine, "\nВаши числа:", a, "и", b, "\n\nСумма =", sum, "\nРазность =", subtract, "\nПроизведение =", multiply, "\nЧастное =", divide, "\n\nКвадрат первого введеного числа =", squared1, "\nКвадрат второго введнеого числа =", squared2, "\nКорень первого введнеого числа =", root1, "\nКорень второго введнеого числа =", root2, "\n", separatingLine)
	println("\nВведите мод или выйдите...\n")
	modeSelection()
}

func main() {
	for i := 0; i < 1; i++ {
		clearTerminal()
		fmt.Print("Введите мод или выйдите...\n'-help' - для изучения команд\n\n")
	}
	for {
		modeSelection()
	}
}
