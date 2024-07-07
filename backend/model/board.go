package model

type Disc int

const (
	Empty Disc = iota
	Dark
	Light
	Wall
)

const E = Empty // 空：0
const D = Dark  // 黒：1
const L = Light // 白：2
const W = Wall  //　壁：3

type Discs [][]Disc // Discの2重配列

type Board struct {
	Discs
}

var INITIAL_BORD Discs = [][]Disc{
	{E, E, E, E, E, E, E, E},
	{E, E, E, E, E, E, E, E},
	{E, E, E, E, E, E, E, E},
	{E, E, E, L, D, E, E, E},
	{E, E, E, D, L, E, E, E},
	{E, E, E, E, E, E, E, E},
	{E, E, E, E, E, E, E, E},
	{E, E, E, E, E, E, E, E},
}

func NewBoard() *Board {
	return &Board{INITIAL_BORD}
}

func ReverseColor(disc Disc) Disc {
	if disc == D {
		return L
	} else {
		return D
	}
}
