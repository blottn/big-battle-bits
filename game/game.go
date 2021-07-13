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

type Phase string

const (
	Init     = Phase("init")
	InCombat = Phase("in_combat")
)

type Game struct {
	gamePhase Phase
	PCs       PlayerConfigs
	Bg        bf.BattleGround `json:"-"` // ignore for marshalling
}

func NewDefaultGame() *Game {
	return NewGame(Init, 256, 256, int(time.Now().Unix()))
}

func NewGame(p Phase, w, h, seed int) *Game {
	pcs := PlayerConfigs{}
	bg := bf.NewBattleGround(w, h, seed)
	return &Game{p, pcs, bg}
}

const StateFile = "state.json"
const BattleGroundFile = "map.png"

type dirName time.Time

func Load(savesDir string) (*Game, error) {
	// read save dirs available
	subDirs, err := ioutil.ReadDir(savesDir)
	if err != nil {
		return nil, err
	}

	// Find newest save
	sorter := fsInfoSorter{&subDirs}
	sort.Sort(sorter)
	subDirs = *(sorter.f)
	d := subDirs[0]
	fmt.Println("Reading save at: " + path.Join(savesDir, d.Name()))
	stateFile, err := os.Open(path.Join(savesDir, d.Name(), StateFile))
	if err != nil {
		return nil, err
	}
	defer stateFile.Close()

	bs, err := ioutil.ReadAll(stateFile)
	if err != nil {
		return nil, err
	}

	var game Game
	err = json.Unmarshal(bs, &game)
	if err != nil {
		return nil, err
	}

	bg, err := bf.ReadFromFile(path.Join(savesDir, d.Name(), BattleGroundFile))
	if err != nil {
		return nil, err
	}
	game.Bg = *bg
	return &game, nil
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
	bs, err := json.Marshal(game)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path.Join(dir, StateFile), bs, 0644)
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
