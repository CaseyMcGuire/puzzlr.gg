package build

import (
	"fmt"
	"os/exec"
)

func RunWebpack() {
	fmt.Println("Building webpack...")

	output, err := exec.Command("npm", "install").CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("Ran into error while running 'npm install': %s", err.Error()))
	}
	fmt.Println(string(output))

	output, err = exec.Command("npm", "run", "webpack").CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("Ran into error while running 'npm run webpack': %s", err.Error()))
	}
	fmt.Println(string(output))
}
