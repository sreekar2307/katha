package model

type Simplifier string

var Simplifiers = struct {
	OneLevel Simplifier
}{
	OneLevel: "ONE_LEVEL",
}
