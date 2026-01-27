package main

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

func initScale() *CircularList[Note] {
	var scale = NewCircularList[Note](12)

	if scale.Size != 12 {
		panic("scale is not the right size!")
	}

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

func calculateInterval(lowerNote Note, higherNote Note) string {
}

func main() {
	scale := initScale()
	head := scale
}
