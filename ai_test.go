package main

import (
	"gopkg.in/inconshreveable/log15.v2"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setupForTesting()
	retCode := m.Run()
	os.Exit(retCode)
}

func setupForTesting() {
	logger = log15.New()
	logger.SetHandler(log15.StreamHandler(os.Stdout, log15.LogfmtFormat()))

	engine = NewEngine()
}

func TestBoardStatsFullLines1(t *testing.T) {
	board = NewBoard()

	// minoO
	mino1 := &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 18
	mino1.SetOnBoard()

	// minoI
	mino1 = &Mino{
		minoRotation: minoBag[0],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 2
	mino1.y = 18

	// minoI
	mino2 := &Mino{
		minoRotation: minoBag[0],
	}
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = 6
	mino2.y = 18

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 1
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	expected = 0
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 1
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	// mino1.SetOnBoard()
	// mino2.SetOnBoard()
	// board.drawDebugBoard()
}

func TestBoardStatsFullLines2(t *testing.T) {
	board = NewBoard()

	// minoO
	mino1 := &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 18
	mino1.SetOnBoard()

	// minoO
	mino1 = &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 2
	mino1.y = 18
	mino1.SetOnBoard()

	// minoO
	mino1 = &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 4
	mino1.y = 18
	mino1.SetOnBoard()

	// minoO
	mino1 = &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 6
	mino1.y = 18
	mino1.SetOnBoard()

	// minoO
	mino1 = &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 8
	mino1.y = 18
	mino1.SetOnBoard()

	// minoO
	mino1 = &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 16
	mino1.SetOnBoard()

	// minoI
	mino1 = &Mino{
		minoRotation: minoBag[0],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 2
	mino1.y = 16

	// minoI
	mino2 := &Mino{
		minoRotation: minoBag[0],
	}
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = 6
	mino2.y = 16

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 3
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	expected = 0
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 1
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	// mino1.SetOnBoard()
	// mino2.SetOnBoard()
	// board.drawDebugBoard()
}

func TestBoardStatsFullLines3(t *testing.T) {
	board = NewBoard()

	// minoO
	mino1 := &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 18
	mino1.SetOnBoard()

	// minoI
	mino1 = &Mino{
		minoRotation: minoBag[0],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 16
	mino1.SetOnBoard()

	// minoO
	mino1 = &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 4
	mino1.y = 18
	mino1.SetOnBoard()

	// minoI
	mino1 = &Mino{
		minoRotation: minoBag[0],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 4
	mino1.y = 16
	mino1.SetOnBoard()

	// minoO
	mino1 = &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 8
	mino1.y = 18

	// minoO
	mino2 := &Mino{
		minoRotation: minoBag[3],
	}
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = 8
	mino2.y = 16

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 1
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	expected = 0
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 9
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	// mino1.SetOnBoard()
	// mino2.SetOnBoard()
	// board.drawDebugBoard()
}

func TestBoardStatsBumpy1(t *testing.T) {
	board = NewBoard()

	// minoO
	mino1 := &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 18

	// minoO
	mino2 := &Mino{
		minoRotation: minoBag[3],
	}
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = 0
	mino2.y = 16

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 0
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	expected = 0
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 4
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	// mino1.SetOnBoard()
	// mino2.SetOnBoard()
	// board.drawDebugBoard()
}

func TestBoardStatsBumpy2(t *testing.T) {
	board = NewBoard()

	// minoO
	mino1 := &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 18
	mino1.SetOnBoard()

	// minoI
	mino1 = &Mino{
		minoRotation: minoBag[0],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 16

	// minoI
	mino2 := &Mino{
		minoRotation: minoBag[0],
	}
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = 0
	mino2.y = 15

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 0
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	// 9 + 9 = 18
	expected = 18
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 4
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	// mino1.SetOnBoard()
	// mino2.SetOnBoard()
	// board.drawDebugBoard()
}

func TestBoardStatsBumpy3(t *testing.T) {
	board = NewBoard()

	// minoO
	mino1 := &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 0
	mino1.y = 18
	mino1.SetOnBoard()

	// minoO
	mino1 = &Mino{
		minoRotation: minoBag[3],
	}
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 1
	mino1.y = 16

	// minoO
	mino2 := &Mino{
		minoRotation: minoBag[3],
	}
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = 2
	mino2.y = 14

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 0
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	// 6 + 7 + 4 + 5 + 6 + 7 = 35
	expected = 35
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 10
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	// mino1.SetOnBoard()
	// mino2.SetOnBoard()
	// board.drawDebugBoard()
}

func TestBoardStatsBumpy4(t *testing.T) {
	board = NewBoard()

	// minoI
	mino1 := &Mino{
		minoRotation: minoBag[0],
	}
	mino1.rotation = 1
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = 6
	mino1.y = 16

	// minoI
	mino2 := &Mino{
		minoRotation: minoBag[0],
	}
	mino2.rotation = 1
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = 6
	mino2.y = 12

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 0
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	expected = 0
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 16
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	// mino1.SetOnBoard()
	// mino2.SetOnBoard()
	// board.drawDebugBoard()
}

func TestBoardStatsHoleDepth1(t *testing.T) {
	board = NewBoard()

	// minoJ
	mino1 := &Mino{
		minoRotation: minoBag[1],
	}
	mino1.rotation = 1
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = -1
	mino1.y = 17

	// minoJ
	mino2 := &Mino{
		minoRotation: minoBag[1],
	}
	mino2.rotation = 1
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = 1
	mino2.y = 17

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 0
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	// 3 + 4 + 3 + 4 = 14
	expected = 14
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 3
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	//	mino1.SetOnBoard()
	//	mino2.SetOnBoard()
	//	board.drawDebugBoard()
}

func TestBoardStatsHoleDepth2(t *testing.T) {
	board = NewBoard()

	// minoJ
	mino1 := &Mino{
		minoRotation: minoBag[1],
	}
	mino1.rotation = 1
	mino1.length = len(mino1.minoRotation[0])
	mino1.x = -1
	mino1.y = 17

	// minoJ
	mino2 := &Mino{
		minoRotation: minoBag[1],
	}
	mino2.rotation = 1
	mino2.length = len(mino2.minoRotation[0])
	mino2.x = -1
	mino2.y = 14

	fullLines, holeDepth, bumpy := board.boardStatsWithMinos(mino1, mino2)
	expected := 0
	if fullLines != expected {
		t.Error("fullLines expected", expected, "got", fullLines)
	}
	// 3 + 4 + 6 + 7 = 20
	expected = 20
	if holeDepth != expected {
		t.Error("holeDepth expected", expected, "got", holeDepth)
	}
	expected = 6
	if bumpy != expected {
		t.Error("bumpy expected", expected, "got", bumpy)
	}

	// for debuging
	//	mino1.SetOnBoard()
	//	mino2.SetOnBoard()
	//	board.drawDebugBoard()
}