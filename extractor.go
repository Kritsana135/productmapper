package productmapper

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type ProductParts struct {
	FilmTypeId string
	TextureId  string
	ModelId    string
	Qty        int
}

func (p *ProductParts) ProductId() string {
	return fmt.Sprintf("%s-%s-%s", p.FilmTypeId, p.TextureId, p.ModelId)
}

func (p *ProductParts) MaterialId() string {
	return fmt.Sprintf("%s-%s", p.FilmTypeId, p.TextureId)
}

const (
	Seperator = '-'
	Splitter  = '/'
	QtySymbol = '*'
)

// - assume prefix contains only uppercase letters and numbers
// - assume texture contains only uppercase letters
func ExtractPlatformId(platformProductId string) ([]ProductParts, int, error) {
	products := []ProductParts{}
	lenId := len(platformProductId)

	var (
		prefixBuilder  strings.Builder
		textureBuilder strings.Builder
		modelBuilder   strings.Builder

		prefixLetterCount int
		prefixDigitCount  int

		totalQty     = 0
		qty          = 1
		hasQtySymbol = false
		qtyDigits    []rune

		state = 0 // parsing prefix, 1: parsing texture, 2: parsing model, 3: parsing quantity, 4: append products
	)

	for i, c := range platformProductId {
		switch state {
		case 0: // parsing prefix
			if unicode.IsDigit(c) {
				prefixBuilder.WriteRune(c)
				prefixDigitCount++
			} else if unicode.IsLetter(c) && unicode.IsUpper(c) {
				prefixBuilder.WriteRune(c)
				prefixLetterCount++
			} else {
				if prefixLetterCount > 0 && prefixDigitCount > 0 && c == Seperator {
					state = 1 // transition to parsing texture
				} else {
					prefixBuilder.Reset()
					prefixDigitCount = 0
					prefixLetterCount = 0
				}
			}
		case 1: // parsing texture
			if unicode.IsLetter(c) && unicode.IsUpper(c) {
				textureBuilder.WriteRune(c)
			} else if textureBuilder.Len() > 0 && c == Seperator {
				state = 2 // transition to parsing model
			} else {
				return products, 0, &ParseError{
					Message: "invalid texture id format",
					Input:   platformProductId,
					Index:   i,
				}
			}
		case 2: // parsing model
			if c == QtySymbol {
				hasQtySymbol = true
				state = 3 // transition to parsing quantity
			} else if c == Splitter {
				state = 4
			} else {
				modelBuilder.WriteRune(c)
			}
		case 3: // parsing quantity
			if unicode.IsDigit(c) {
				qtyDigits = append(qtyDigits, c)
			} else if c == Splitter {
				state = 4
			}
		}

		if (i == lenId-1 && state == 2) || (i == lenId-1 && state == 3) {
			state = 4
		}

		if state == 4 {
			if hasQtySymbol {
				if len(qtyDigits) == 0 {
					return products, 0, &ParseError{
						Message: "quantity symbol '*' found but no digits followed",
						Input:   platformProductId,
					}
				}
				qtyVal, err := strconv.Atoi(string(qtyDigits))
				if err != nil {
					return products, 0, &ParseError{
						Message: "failed to parse quantity",
						Input:   platformProductId,
					}
				}
				qty = qtyVal
			}

			if prefixBuilder.Len() == 0 || textureBuilder.Len() == 0 || modelBuilder.Len() == 0 {
				return products, 0, &ParseError{
					Message: "invalid format",
					Input:   platformProductId,
				}
			}

			products = append(products, ProductParts{
				FilmTypeId: prefixBuilder.String(),
				TextureId:  textureBuilder.String(),
				ModelId:    modelBuilder.String(),
				Qty:        qty,
			})

			totalQty += qty

			prefixBuilder.Reset()
			textureBuilder.Reset()
			modelBuilder.Reset()

			prefixLetterCount = 0
			prefixDigitCount = 0

			qty = 1
			hasQtySymbol = false
			qtyDigits = []rune{}
			state = 0
		}
	}

	if lenId > 0 && len(products) == 0 {
		return products, 0, &ParseError{
			Message: "can't extract product from input",
			Input:   platformProductId,
		}
	}

	return products, totalQty, nil
}

type ParseError struct {
	Message string
	Input   string
	Index   int
}

func (e *ParseError) Error() string {
	if e.Index != 0 {
		return "Parse Error: " + e.Message + " at index " + strconv.Itoa(e.Index) + " in '" + e.Input + "'"
	}
	return "Parse Error: " + e.Message + " in '" + e.Input + "'"
}
