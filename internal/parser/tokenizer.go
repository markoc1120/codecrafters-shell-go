package parser

import (
	"bufio"
	"fmt"
	"io"
)

// internal state
type tokenizerState int
type runeTokenClass int

const (
	spaceRunes       = " \t\r\n"
	singleQuoteRunes = "'"
)

// Classes of rune token
const (
	unknownRuneClass runeTokenClass = iota
	spaceRuneClass
	singleQuoteRuneClass
)

// tokenizer states
const (
	defaultState tokenizerState = iota
	singleQuoteState
)

type tokenClassifier map[rune]runeTokenClass

func (typeMap tokenClassifier) addRuneClass(runes string, tokenType runeTokenClass) {
	for _, runeChar := range runes {
		typeMap[runeChar] = tokenType
	}
}

func newDefaultClassifier() tokenClassifier {
	t := tokenClassifier{}
	t.addRuneClass(spaceRunes, spaceRuneClass)
	t.addRuneClass(singleQuoteRunes, singleQuoteRuneClass)
	return t
}

func (t tokenClassifier) ClassifyRune(runeVal rune) runeTokenClass {
	if class, found := t[runeVal]; found {
		return class
	}
	return unknownRuneClass
}

type Tokenizer struct {
	input      *bufio.Reader
	classifier tokenClassifier
}

func newTokenizer(r io.Reader) *Tokenizer {
	return &Tokenizer{input: bufio.NewReader(r), classifier: newDefaultClassifier()}
}

func (t *Tokenizer) Next() (string, error) {
	state := defaultState
	var argument []rune

	for {
		nextRune, _, err := t.input.ReadRune()
		nextRuneType := t.classifier.ClassifyRune(nextRune)

		if err == io.EOF {
			if state == singleQuoteState {
				return "", fmt.Errorf("unclosed single quote")
			}
			if len(argument) == 0 {
				return "", io.EOF
			}
			return string(argument), io.EOF
		}

		switch state {
		case defaultState:
			switch nextRuneType {
			case spaceRuneClass:
				if len(argument) > 0 {
					return string(argument), nil
				}
			case singleQuoteRuneClass:
				state = singleQuoteState
			default:
				argument = append(argument, nextRune)
			}
		case singleQuoteState:
			switch nextRuneType {
			case singleQuoteRuneClass:
				state = defaultState
			default:
				argument = append(argument, nextRune)
			}
		}
	}
}
