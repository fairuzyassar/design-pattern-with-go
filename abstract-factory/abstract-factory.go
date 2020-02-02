package main

import "fmt"

type Level interface {
	chooseLevel() string
}

type EasyLevel struct{}

type HardLevel struct{}

func (e *EasyLevel) chooseLevel() string {
	return "Easy Level"
}

func (e *HardLevel) chooseLevel() string {
	return "Hard Level"
}

type Arena interface {
	chooseArena() string
}

type EasyArena struct{}

type HardArena struct{}

func (e *EasyArena) chooseArena() string {
	return "Easy Arena"
}

func (e *HardArena) chooseArena() string {
	return "Hard Arena"
}

type GamesFactory interface {
	createLevel() Level
	createArena() Arena
}

type EasyGame struct{}

type HardGame struct{}

func (eg *EasyGame) createLevel() Level {
	return &EasyLevel{}
}

func (eg *EasyGame) createArena() Arena {
	return &EasyArena{}
}

func (hg *HardGame) createLevel() Level {
	return &HardLevel{}
}

func (hg *HardGame) createArena() Arena {
	return &HardArena{}
}

type Games struct {
	level Level
	arena Arena
}

func main() {
	easyGame := &EasyGame{}
	hardGame := &HardGame{}

	GamesA := &Games{easyGame.createLevel(), easyGame.createArena()}
	GamesB := &Games{hardGame.createLevel(), hardGame.createArena()}
	
	fmt.Println("Games A", GamesA.level.chooseLevel(), GamesA.arena.chooseArena())
	fmt.Println("Games B", GamesB.level.chooseLevel(), GamesB.arena.chooseArena())


}