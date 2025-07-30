package main

import (
"strings"
)

type Morph interface {
	Name() string
	Generate(n *Name) string
}

var morphs = []Morph{
	&FirstMorph{},
	&FirstLastMorph{},
	&FirstILastIMorph{},
	&FirstILastMorph{},
	&FirstDotLastMorph{},
	&FirstIDotLastMorph{},
	&LastFirstMorph{},
	&LastIFirstMorph{},
	&LastDotFirstMorph{},
	&FirstUnderLastMorph{},
	&FirstIUnderLastMorph{},
	&LastUnderFirstMorph{},
	&FirstUnderLastIMorph{},
	&FirstNumberMorph{},
	&FirstLastYearMorph{},
	&FirstLastNoVowelsMorph{},
	&FirstThreeLastThreeMorph{},
	&FirstFourMorph{},
	&LastFourMorph{},
	&FirstAtLastMorph{},
	&LowerFirstLastMorph{},
	&UpperFirstLastMorph{},
	&CamelFirstLastMorph{},
}

type FirstMorph struct{}

func (p *FirstMorph) Name() string            { return "FirstName" }
func (p *FirstMorph) Generate(n *Name) string { return n.firstname }

type FirstLastMorph struct{}

func (p *FirstLastMorph) Name() string            { return "FirstnameLastname" }
func (p *FirstLastMorph) Generate(n *Name) string { return n.firstname + n.lastname }

type FirstILastMorph struct{}

func (p *FirstILastMorph) Name() string            { return "FirstinitialLastname" }
func (p *FirstILastMorph) Generate(n *Name) string { return n.firstInitial + n.lastname }

type FirstILastIMorph struct{}

func (p *FirstILastIMorph) Name() string            { return "FirstinitialLastInitial" }
func (p *FirstILastIMorph) Generate(n *Name) string { return n.firstInitial + n.lastInitial }

type FirstDotLastMorph struct{}

func (p *FirstDotLastMorph) Name() string            { return "First Name Dot LastName" }
func (p *FirstDotLastMorph) Generate(n *Name) string { return n.firstname + "." + n.lastname }

type FirstIDotLastMorph struct{}

func (p *FirstIDotLastMorph) Name() string            { return "First Initial Dot LastName" }
func (p *FirstIDotLastMorph) Generate(n *Name) string { return n.firstInitial + "." + n.lastname }


// Reverse patterns
type LastFirstMorph struct{}
func (p *LastFirstMorph) Name() string            { return "Lastname Firstname" }
func (p *LastFirstMorph) Generate(n *Name) string { return n.lastname + n.firstname }

type LastIFirstMorph struct{}
func (p *LastIFirstMorph) Name() string            { return "Lastname First Initial" }
func (p *LastIFirstMorph) Generate(n *Name) string { return n.lastname + n.firstInitial }


type LastDotFirstMorph struct{}
func (p *LastDotFirstMorph) Name() string            { return "Lastname.Firstname" }
func (p *LastDotFirstMorph) Generate(n *Name) string { return n.lastname + "." + n.firstname }

// Underscore patterns
type FirstUnderLastMorph struct{}
func (p *FirstUnderLastMorph) Name() string            { return "First_Last" }
func (p *FirstUnderLastMorph) Generate(n *Name) string { return n.firstname + "_" + n.lastname }

type FirstIUnderLastMorph struct{}
func (p *FirstIUnderLastMorph) Name() string            { return "FirstInitial_Last" }
func (p *FirstIUnderLastMorph) Generate(n *Name) string { return n.firstInitial + "_" + n.lastname }

type LastUnderFirstMorph struct{}
func (p *LastUnderFirstMorph) Name() string            { return "Last_First" }
func (p *LastUnderFirstMorph) Generate(n *Name) string { return n.lastname + "_" + n.firstname }

type FirstUnderLastIMorph struct{}
func (p *FirstUnderLastIMorph) Name() string            { return "First_LastInitial" }
func (p *FirstUnderLastIMorph) Generate(n *Name) string { return n.firstname + "_" + n.lastInitial }

type FirstNumberMorph struct{}
func (p *FirstNumberMorph) Name() string            { return "First + Numbers" }
func (p *FirstNumberMorph) Generate(n *Name) string { return n.firstname + "123" }

type FirstLastYearMorph struct{}
func (p *FirstLastYearMorph) Name() string            { return "FirstLast + Year" }
func (p *FirstLastYearMorph) Generate(n *Name) string { return n.firstname + n.lastname + "2024" }

// Abbreviation patterns
type FirstLastNoVowelsMorph struct{}
func (p *FirstLastNoVowelsMorph) Name() string            { return "FirstLast No Vowels" }
func (p *FirstLastNoVowelsMorph) Generate(n *Name) string { 
	removeVowels := func(s string) string {
		vowels := "aeiouAEIOU"
		result := ""
		for _, char := range s {
			if !strings.ContainsRune(vowels, char) {
				result += string(char)
			}
		}
		return result
	}
	return removeVowels(n.firstname + n.lastname)
}

type FirstThreeLastThreeMorph struct{}
func (p *FirstThreeLastThreeMorph) Name() string            { return "First3Last3" }
func (p *FirstThreeLastThreeMorph) Generate(n *Name) string { 
	first3 := n.firstname
	if len(first3) > 3 {
		first3 = first3[:3]
	}
	last3 := n.lastname
	if len(last3) > 3 {
		last3 = last3[:3]
	}
	return first3 + last3
}

type FirstFourMorph struct{}
func (p *FirstFourMorph) Name() string            { return "First 4 chars" }
func (p *FirstFourMorph) Generate(n *Name) string { 
	if len(n.firstname) > 4 {
		return n.firstname[:4]
	}
	return n.firstname
}

type LastFourMorph struct{}
func (p *LastFourMorph) Name() string            { return "Last 4 chars" }
func (p *LastFourMorph) Generate(n *Name) string { 
	if len(n.lastname) > 4 {
		return n.lastname[:4]
	}
	return n.lastname
}

type FirstAtLastMorph struct{}
func (p *FirstAtLastMorph) Name() string            { return "First@Last" }
func (p *FirstAtLastMorph) Generate(n *Name) string { return n.firstname + "@" + n.lastname }

// Case variations
type LowerFirstLastMorph struct{}
func (p *LowerFirstLastMorph) Name() string            { return "lowercase firstlast" }
func (p *LowerFirstLastMorph) Generate(n *Name) string { 
	return strings.ToLower(n.firstname + n.lastname)
}

type UpperFirstLastMorph struct{}
func (p *UpperFirstLastMorph) Name() string            { return "UPPERCASE FIRSTLAST" }
func (p *UpperFirstLastMorph) Generate(n *Name) string { 
	return strings.ToUpper(n.firstname + n.lastname)
}

type CamelFirstLastMorph struct{}
func (p *CamelFirstLastMorph) Name() string            { return "camelCase firstLast" }
func (p *CamelFirstLastMorph) Generate(n *Name) string { 
	return strings.ToLower(n.firstname) + strings.Title(n.lastname)
}

