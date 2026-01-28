package main

import "errors"

type Note int

const (
	A Note = iota
	ASharp
	B
	C
	CSharp
	D
	DSharp
	E
	F
	FSharp
	G
	GSharp
)

var noteNames = map[string]Note{
	"A":  Note(A),
	"A#": Note(ASharp),
	"Bb": Note(ASharp),
	"B":  Note(B),
	"C":  Note(C),
	"C#": Note(CSharp),
	"Db": Note(CSharp),
	"D":  Note(D),
	"D#": Note(DSharp),
	"Eb": Note(DSharp),
	"E":  Note(E),
	"F":  Note(F),
	"F#": Note(FSharp),
	"Gb": Note(FSharp),
	"G":  Note(G),
	"G#": Note(GSharp),
	"Ab": Note(GSharp),
}

var intervals = map[int]string{
	0:  "Perfect Unison",
	1:  "Minor Second",
	2:  "Major Second",
	3:  "Mintor Third",
	4:  "Major Third",
	5:  "Perfect Fourth",
	6:  "Augmented Fourth?",
	7:  "Perfect Fifth",
	8:  "Minor Sixth",
	9:  "Major Sixth",
	10: "Minor Seventh",
	11: "Major Seventh",
	12: "Perfect Octave",
}

type ChordType struct {
	Name      string
	Intervals []uint
}

var chordTypes0 = map[string][]uint{
	"Major Triad": {4, 7},
	"Minor Triad": {3, 7},
}

/*
TODO:
- implement bi-directional hashmap
- implement hashset

then chord types can be a bi-directional map of {set of notes : chord name}
*/

func initScale() *CircularList[Note] {
	var scale = NewCircularList[Note](12)

	scale.set(0, Note(A))
	scale.set(1, Note(ASharp))
	scale.set(2, Note(B))
	scale.set(3, Note(C))
	scale.set(4, Note(CSharp))
	scale.set(5, Note(D))
	scale.set(6, Note(DSharp))
	scale.set(7, Note(E))
	scale.set(8, Note(F))
	scale.set(9, Note(FSharp))
	scale.set(10, Note(G))
	scale.set(11, Note(GSharp))

	return scale
}

func calculateInterval(scale *CircularList[Note], lowerNote Note, higherNote Note) (string, error) {
	//todo: get both values out of modulo to determine differences in octave
	lowerNoteIndex, err := scale.getByVal(lowerNote)
	if err != nil {
		return "", errors.New("Could not locate lower note")
	}

	higherNoteIndex, err := scale.getByVal(higherNote)
	if err != nil {
		return "", errors.New("Could not locate higher note")
	}

	semitones := higherNoteIndex - lowerNoteIndex

	if semitones > 12 {
		//todo: improve this error message
		return "", errors.New("Determined difference in semitones to be invalid")
	}

	interval := intervals[int(semitones)]
	return interval, nil
}

func main() {
	scale := initScale()

	interval, err := calculateInterval(scale, Note(ASharp), Note(D))
	if err != nil {
		println(err)
	}
	println(interval)
}
