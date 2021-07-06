package bf

import (
    "fmt"
    "os"
)

type Mutator func(*BattleGround) (error)

// Mutate provides utility for reading and writing a battleground in one go
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

func ReadFromFile(inputName string) (*BattleGround, error) {
    file, err := os.Open(inputName)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    return From(file)
}

func WriteTo(bg *BattleGround, outputName string) error {
    out, err := os.OpenFile(outputName, os.O_RDWR|os.O_CREATE, 0755)
    if err != nil {
        return nil
    }
    defer out.Close()
    (*bg).Output(out)
    return nil
}
