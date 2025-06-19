package main

import (
	"fmt"
)

type Instrument struct {
	Cost		float
	InstClass	string
	InstName	string
	NeedsRepair	bool	
}

type Class struct {
	Brass		string
	Woodwind	string
	Strings		string
	Percussion	string
}