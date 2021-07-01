package bf

import (
    "fmt"
    "os"
)
type Mutator func(*BattleGround) (error)

func Mutate(inputName, outputName string, mutator Mutator) {
	file, err := os.Open(inputName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	bg, err := From(file)
	if err != nil {
		fmt.Println(err.Error())
	}

    err = mutator(bg)

    if err != nil {
        fmt.Println(err.Error())
    }

	out, _ := os.OpenFile(outputName, os.O_RDWR|os.O_CREATE, 0755)
	defer out.Close()
	bg.Output(out)
}
