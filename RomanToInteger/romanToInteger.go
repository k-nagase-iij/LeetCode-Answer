package RomanToInteger

var allSymbolCombinations = map[string]int{
	"I": 1,
	"IV": 4,
	"V": 5,
	"IX": 9,
	"X": 10,
	"XL": 40,
	"L": 50,
	"XC": 90,
	"C": 100,
	"CD": 400,
	"D": 500,
	"CM":900,
	"M": 1000,
}

// 渡された一文字が引く数(I, X, C)に該当するかどうか
func isSubSymbol(symbol byte) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}

// 渡された二文字がsymbols mapに存在するかどうか
func isExistSymbols(symbolsWithSub ...byte) bool {
	strSymbols := string(symbolsWithSub)
	if _, ok := allSymbolCombinations[strSymbols]; ok {
		return true
	}

	return false
}

func splitSymbols(s string) []string {
	var ret []string

	for i := 0;i < len(s);i++ {
		upperSymbols := s[i]
		isEnd := i + 1 < len(s)

		if isEnd && isSubSymbol(upperSymbols) && isExistSymbols(upperSymbols, s[i+1]) {
			byteSymbols := []byte{upperSymbols, s[i+1]}
			ret = append(ret, string(byteSymbols))
			i++
		} else {
			ret = append(ret, string(upperSymbols))
		}
	}

	return ret
}

func romanToInt(s string) int {
	ret := 0
	arraySymbols := splitSymbols(s)

	for _, str := range arraySymbols {
		ret += allSymbolCombinations[str]
	}

    return ret
}