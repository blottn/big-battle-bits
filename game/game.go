package game

import (
	"big-battle-bits/bf"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
	"time"
)

type Game struct {
	PCs PlayerConfigs
	Bg  bf.BattleGround
}

const PlayerConfigsFile = "players.json"
const BattleGroundFile = "map.png"

type dirName time.Time

func Load(savesDir string) (*Game, error) {
	// read save dirs available
	subDirs, err := ioutil.ReadDir(savesDir)
	if err != nil {
		return nil, err
	}

	// Do the sort thing!
	sorter := fsInfoSorter{&subDirs}
	sort.Sort(sorter)
	subDirs = *(sorter.f)
	d := subDirs[0]
	fmt.Println("Reading save at: " + path.Join(savesDir, d.Name()))
	pFile, err := os.Open(path.Join(savesDir, d.Name(), PlayerConfigsFile))
	if err != nil {
		return nil, err
	}
	defer pFile.Close()

	bs, err := ioutil.ReadAll(pFile)
	if err != nil {
		return nil, err
	}

	var pcs PlayerConfigs
	err = json.Unmarshal(bs, &pcs)
	if err != nil {
		return nil, err
	}

	bg, err := bf.ReadFromFile(path.Join(savesDir, d.Name(), BattleGroundFile))
	if err != nil {
		return nil, err
	}
	return &Game{pcs, *bg}, nil
}

func Save(savesDir string, game *Game) error {
	saveName := fmt.Sprintf("%d", (time.Now().Unix()))
	fmt.Println("save dir: " + path.Join(savesDir, saveName))
	dir := path.Join(savesDir, saveName)
	err := os.Mkdir(dir, 0700)
	if err != nil {
		return err
	}

	// Write playerconfigs
	bs, err := json.Marshal(game.PCs)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(dir, PlayerConfigsFile), bs, 0644)
	if err != nil {
		return err
	}

	// Write map file
	mapFile, err := os.OpenFile(path.Join(dir, BattleGroundFile), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer mapFile.Close()
	game.Bg.Output(mapFile)
	return nil
}

type fsInfoSorter struct {
	f *[]os.FileInfo
}

func (fs fsInfoSorter) Len() int {
	return len(*fs.f)
}

func (fs fsInfoSorter) Less(i, j int) bool {
	return (*fs.f)[i].ModTime().After((*fs.f)[j].ModTime())
}

func (fs fsInfoSorter) Swap(i, j int) {
	temp := (*fs.f)[i]
	(*fs.f)[i] = (*fs.f)[j]
	(*fs.f)[j] = temp
}
