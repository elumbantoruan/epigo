package ch07strings

var digitMapper map[rune]string

func MnemonicPhoneNumber(s string) []string {
	initializeDigitMapper()
	mnemonics := []string{}
	var partialMnemonics []rune
	for i := 0; i < len(s); i++ {
		partialMnemonics = append(partialMnemonics, ' ')
	}
	digits := []rune(s)
	ms := BuildMnemonicPhoneNumber(digits, 0, partialMnemonics, mnemonics)
	return ms
}

func BuildMnemonicPhoneNumber(digits []rune, index int, partialMnemonics []rune, mnemonics []string) []string{
	if index == len(digits) {
		return append(mnemonics, string(partialMnemonics))
	} else {
		s := digitMapper[digits[index]]
		for r := 0; r < len(s); r++ {
			runes := []rune(s)
			partialMnemonics[index] = runes[r]
			ms := BuildMnemonicPhoneNumber(digits, index+1, partialMnemonics, mnemonics)
			mnemonics = ms
		}
		return mnemonics
	}
}

func initializeDigitMapper() {
	digitMapper = make(map[rune]string)
	digitMapper['0'] = "0"
	digitMapper['1'] = "1"
	digitMapper['2'] = "ABC"
	digitMapper['3'] = "DEF"
	digitMapper['4'] = "GHI"
	digitMapper['5'] = "JKL"
	digitMapper['6'] = "MNO"
	digitMapper['7'] = "PQRS"
	digitMapper['8'] = "TUV"
	digitMapper['9'] = "WXYZ"
}
