package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

// Homoglyphs map for ASCII characters
var homoglyphs = map[rune][]rune{
	'a': {'а', 'ａ', '𝐚', '𝒂', '𝖆', '𝕒', 'a', 'ạ', 'ą', 'ä', 'à', 'á'},
	'b': {'Ь', 'ｂ', '𝐛', '𝑏', '𝖇', '𝕓', 'b', 'ḅ', 'ḃ', 'ƀ'},
	'c': {'с', 'ｃ', '𝐜', '𝑐', '𝖈', '𝕔', 'c', 'ç', 'ć', 'č'},
	'd': {'ԁ', 'ｄ', '𝐝', '𝑑', '𝖉', '𝕕', 'd', 'ḍ', 'ḋ', 'đ'},
	'e': {'е', 'ｅ', '𝐞', '𝑒', '𝖊', '𝕖', 'e', 'ẹ', 'ę', 'ë', 'è', 'é', '℮'},
	'f': {'ｆ', '𝐟', '𝑓', '𝖋', '𝕗', 'f', 'ḟ', 'ƒ'},
	'g': {'ɡ', 'ｇ', '𝐠', '𝑔', '𝖌', '𝕘', 'g', 'ġ', 'ğ', 'ĝ'},
	'h': {'һ', 'ｈ', '𝐡', 'ℎ', '𝖍', '𝕙', 'h', 'ḥ', 'ḣ', 'ĥ'},
	'i': {'і', '１', 'ｉ', '𝐢', '𝑖', '𝖎', '𝕚', 'i', 'ị', 'į', 'ï', 'ì', 'í', 'ı'},
	'j': {'ј', 'ｊ', '𝐣', '𝑗', '𝖏', '𝕛', 'j', 'ĵ'},
	'k': {'κ', 'ｋ', '𝐤', '𝑘', '𝖐', '𝕜', 'k', 'ḳ', 'ķ', 'ḱ'},
	'l': {'ｌ', '𝐥', '𝑙', '𝖑', '𝕝', 'l', 'ḷ', 'ļ', 'ĺ', '1', 'I'},
	'm': {'ｍ', '𝐦', '𝑚', '𝖒', '𝕞', 'm', 'ṃ', 'ṁ'},
	'n': {'ｎ', '𝐧', '𝑛', '𝖓', '𝕟', 'n', 'ṇ', 'ṅ', 'ń', 'ñ'},
	'o': {'о', '０', 'ｏ', '𝐨', '𝑜', '𝖔', '𝕠', 'o', 'ọ', 'ǫ', 'ö', 'ò', 'ó', 'ø'},
	'p': {'р', 'ｐ', '𝐩', '𝑝', '𝖕', '𝕡', 'p', 'ṗ', 'ṕ'},
	'q': {'ｑ', '𝐪', '𝑞', '𝖖', '𝕢', 'q', 'ʠ'},
	'r': {'ｒ', '𝐫', '𝑟', '𝖗', '𝐫', 'r', 'ṛ', 'ŕ', 'ř'},
	's': {'ѕ', 'ｓ', '𝐬', '𝑠', '𝖘', '𝕤', 's', 'ṣ', 'ś', 'š', '$'},
	't': {'ｔ', '𝐭', '𝑡', '𝖙', '𝕥', 't', 'ṭ', 'ṫ', 'ť', 'τ'},
	'u': {'ｕ', '𝐮', '𝑢', '𝖚', '𝕦', 'u', 'ụ', 'ų', 'ü', 'ù', 'ú', 'µ'},
	'v': {'ｖ', '𝐯', '𝑣', '𝖛', '𝕧', 'v', 'ṿ', 'ν'},
	'w': {'ｗ', '𝐰', '𝑤', '𝖜', '𝕨', 'w', 'ẉ', 'ẁ', 'ẃ'},
	'x': {'ｘ', '𝐱', '𝑥', '𝖝', '𝕩', 'x', 'ẋ', '×'},
	'y': {'ｙ', '𝐲', '𝑦', '𝖞', '𝕪', 'y', 'ỵ', 'ẏ', 'ý', 'ÿ'},
	'z': {'ｚ', '𝐳', '𝐳', '𝖟', '𝕫', 'z', 'ẓ', 'ż', 'ź', 'ž'},
	'A': {'А', 'Ａ', '𝐀', '𝐴', '𝕬', '𝔸'},
	'B': {'В', 'Ｂ', '𝐁', '𝐵', '𝕭', '𝔹'},
	'C': {'С', 'Ｃ', '𝐂', '𝐶', '𝕮', 'ℂ'},
	'E': {'Е', 'Ｅ', '𝐄', '𝐸', '𝕰', '𝔼'},
	'H': {'Н', 'Ｈ', '𝐇', '𝐻', '𝕳', 'ℍ'},
	'I': {'Ｉ', '𝐈', '𝐼', '𝕴', '𝕀', 'l', '1'},
	'J': {'Ｊ', '𝐉', '𝐽', '𝕵', '𝕁'},
	'K': {'Ｋ', '𝐊', '𝐾', '𝕶', '𝕂'},
	'M': {'М', 'Ｍ', '𝐌', '𝑀', '𝕸', '𝕄'},
	'N': {'Ｎ', '𝐍', '𝑁', '𝕹', 'ℕ'},
	'O': {'О', '０', 'Ｏ', '𝐎', '𝑂', '𝕺', '𝕆'},
	'P': {'Р', 'Ｐ', '𝐏', '𝑃', '𝕻', 'ℙ'},
	'S': {'Ｓ', '𝐒', '𝑆', '𝕾', '𝕊', '5'},
	'T': {'Т', 'Ｔ', '𝐓', '𝑇', '𝕿', '𝕋'},
	'X': {'Х', 'Ｘ', '𝐗', '𝑋', '𝕏', '𝕏'},
	'Y': {'Ｙ', '𝐘', '𝑌', '𝖄', '𝕐'},
	'Z': {'Ｚ', '𝐙', '𝑍', '𝖅', 'ℤ', '2'},
	'0': {'O', 'o', '０', '𝟎', '𝟘'},
	'1': {'l', 'I', '１', '𝟏', '𝟙'},
	'2': {'Z', '２', '𝟐', '𝟚'},
	'5': {'S', 's', '５', '𝟓', '𝟝'},
}

var zeroWidthChars = []rune{
	'\u200B', // Zero Width Space
	'\u200C', // Zero Width Non-Joiner
	'\u200D', // Zero Width Joiner
	'\u2060', // Word Joiner
	'\uFEFF', // Zero Width No-Break Space
}

type Variant struct {
	Text    string  `json:"text"`
	Entropy float64 `json:"entropy,omitempty"`
}

func main() {
	jsonFlag := flag.Bool("j", false, "JSON array output")
	entropyFlag := flag.Bool("e", false, "Entropy-sort")
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "Usage: confusio [options] <keyword>")
		os.Exit(1)
	}
	keyword := args[0]

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	variants := generateVariants(keyword, 100) // Generate plenty to ensure >= 30

	// Ensure unique
	uniqueMap := make(map[string]bool)
	var uniqueVariants []Variant

	for _, v := range variants {
		if !uniqueMap[v] && v != keyword {
			uniqueMap[v] = true
			uniqueVariants = append(uniqueVariants, Variant{
				Text:    v,
				Entropy: calculateEntropy(v),
			})
		}
	}

	if *entropyFlag {
		sort.Slice(uniqueVariants, func(i, j int) bool {
			// Higher entropy means more complexity/disorder
			return uniqueVariants[i].Entropy < uniqueVariants[j].Entropy
		})
	}

	// Output
	if *jsonFlag {
		outputList := make([]string, len(uniqueVariants))
		for i, v := range uniqueVariants {
			outputList[i] = v.Text
		}

		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		encoder.Encode(outputList)
	} else {
		for _, v := range uniqueVariants {
			fmt.Println(v.Text)
		}
	}
}

func generateVariants(s string, minCount int) []string {
	var res []string
	runes := []rune(s)
	seen := make(map[string]bool)

	attempts := 0
	maxAttempts := minCount * 50

	for len(res) < minCount && attempts < maxAttempts {
		attempts++

		// Decide mutation type:
		// 0: Substitute 1 char
		// 1: Substitute 2 chars
		// 2: Substitute all chars
		// 3: Insert zero-width
		// 4: Mixed

		mode := rand.Intn(5)
		newRunes := make([]rune, len(runes))
		copy(newRunes, runes)

		switch mode {
		case 0: // Sub 1
			if len(newRunes) > 0 {
				idx := rand.Intn(len(newRunes))
				if subs, ok := homoglyphs[newRunes[idx]]; ok {
					newRunes[idx] = subs[rand.Intn(len(subs))]
				}
			}
		case 1: // Sub multiple
			if len(newRunes) > 0 {
				num := rand.Intn(len(newRunes)) + 1
				for k := 0; k < num; k++ {
					idx := rand.Intn(len(newRunes))
					if subs, ok := homoglyphs[newRunes[idx]]; ok {
						newRunes[idx] = subs[rand.Intn(len(subs))]
					}
				}
			}
		case 2: // Sub all
			for idx := range newRunes {
				if subs, ok := homoglyphs[newRunes[idx]]; ok {
					newRunes[idx] = subs[rand.Intn(len(subs))]
				}
			}
		case 3: // Insert zero-width
			// Insert at random pos
			idx := rand.Intn(len(newRunes) + 1)
			zw := zeroWidthChars[rand.Intn(len(zeroWidthChars))]

			// Insert
			temp := make([]rune, 0, len(newRunes)+1)
			temp = append(temp, newRunes[:idx]...)
			temp = append(temp, zw)
			temp = append(temp, newRunes[idx:]...)
			newRunes = temp
		case 4: // Mixed sub + insert
			// First sub
			if len(newRunes) > 0 {
				num := rand.Intn(len(newRunes)) + 1
				for k := 0; k < num; k++ {
					idx := rand.Intn(len(newRunes))
					if subs, ok := homoglyphs[newRunes[idx]]; ok {
						newRunes[idx] = subs[rand.Intn(len(subs))]
					}
				}
			}
			// Then insert
			if rand.Float32() < 0.5 {
				idx := rand.Intn(len(newRunes) + 1)
				zw := zeroWidthChars[rand.Intn(len(zeroWidthChars))]
				temp := make([]rune, 0, len(newRunes)+1)
				temp = append(temp, newRunes[:idx]...)
				temp = append(temp, zw)
				temp = append(temp, newRunes[idx:]...)
				newRunes = temp
			}
		}

		variant := string(newRunes)
		if variant != s && !seen[variant] {
			seen[variant] = true
			res = append(res, variant)
		}
	}

	return res
}

func calculateEntropy(s string) float64 {
	if len(s) == 0 {
		return 0
	}
	counts := make(map[rune]int)
	for _, r := range s {
		counts[r]++
	}

	var entropy float64
	length := float64(len([]rune(s))) // Rune count

	for _, count := range counts {
		p := float64(count) / length
		entropy -= p * math.Log2(p)
	}

	return entropy
}
