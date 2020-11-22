package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {
	fmt.Println("Part 49 Exec")
	fmt.Println("")

	fmt.Println("Part 49.1 Penggunaan exec")
	execA, _ := exec.Command("ls").Output()
	execB, _ := exec.Command("ls", "-la").Output()
	execC, _ := exec.Command("git", "config", "--global", "--list").Output()
	execD, _ := exec.Command("go", "version").Output()
	execE, _ := exec.Command("node", "--version").Output()
	execF, _ := exec.Command("vue", "--version").Output()
	execG, _ := exec.Command("adonis", "--version").Output()
	fmt.Printf("execA: %s\n", execA)
	fmt.Printf("execB: %s\n", execB)
	fmt.Printf("execC: %s\n", execC)
	fmt.Printf("execD golang: %s\n", execD)
	fmt.Printf("execE nodejs: %s\n", execE)
	fmt.Printf("execF vue: %s\n", execF)
	fmt.Printf("execG adonis: %s\n", execG)
	fmt.Println("")

	fmt.Println("Part 49.2 Rekomendasi penggunaan exec")
	if runtime.GOOS == "linux" {
		execRuntimeOsA, _ := exec.Command("bash", "-c", "git config user.name").Output()
		execRuntimeOsB, _ := exec.Command("bash", "-c", "git config user.email").Output()
		execRuntimeOsC, _ := exec.Command("bash", "-c", "go version").Output()
		execRuntimeOsD, _ := exec.Command("bash", "-c", "node --version").Output()
		execRuntimeOsE, _ := exec.Command("bash", "-c", "vue --version").Output()
		execRuntimeOsF, _ := exec.Command("bash", "-c", "adonis --version").Output()
		fmt.Printf("execRuntimeOsA: %s", execRuntimeOsA)
		fmt.Printf("execRuntimeOsB: %s", execRuntimeOsB)
		fmt.Printf("execRuntimeOsC: %s", execRuntimeOsC)
		fmt.Printf("execRuntimeOsD: %s", execRuntimeOsD)
		fmt.Printf("execRuntimeOsE: %s", execRuntimeOsE)
		fmt.Printf("execRuntimeOsF: %s", execRuntimeOsF)
	} else {
		execRuntimeOsA, _ := exec.Command("cmd", "/C", "git config user.name").Output()
		execRuntimeOsB, _ := exec.Command("cmd", "/C", "git config user.email").Output()
		execRuntimeOsC, _ := exec.Command("cmd", "/C", "go version").Output()
		execRuntimeOsD, _ := exec.Command("cmd", "/C", "node --version").Output()
		execRuntimeOsE, _ := exec.Command("cmd", "/C", "vue --version").Output()
		execRuntimeOsF, _ := exec.Command("cmd", "/C", "adonis --version").Output()
		fmt.Printf("execRuntimeOsA: %s", execRuntimeOsA)
		fmt.Printf("execRuntimeOsB: %s", execRuntimeOsB)
		fmt.Printf("execRuntimeOsC: %s", execRuntimeOsC)
		fmt.Printf("execRuntimeOsD: %s", execRuntimeOsD)
		fmt.Printf("execRuntimeOsE: %s", execRuntimeOsE)
		fmt.Printf("execRuntimeOsF: %s", execRuntimeOsF)
	}
}
