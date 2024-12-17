package calc

import (
	"github.com/golkity/calc_go/internal/Errors"
	"strconv"
	"strings"
)

func StringToFloat64(str string) float64 {
	degree := float64(1)
	var res float64 = 0
	var invers bool = false
	for i := len(str); i > 0; i-- {
		if str[i-1] == '-' {
			invers = true
		} else {
			res += float64(9-int('9'-str[i-1])) * degree
			degree *= 10
		}
	}
	if invers {
		res = 0 - res
	}
	return res
}

func IsSign(value rune) bool {
	return value == '+' || value == '-' || value == '*' || value == '/'
}

func Calc(expression string) (float64, error) {
	if len(expression) < 3 {
		return 0, Errors.ErrInvalidExpression
	}

	var res float64
	var b string
	var c rune = 0
	var resflag bool = false
	var isc int
	var countc int = 0

	for _, value := range expression {
		if IsSign(value) {
			countc++
		}
	}

	if IsSign(rune(expression[0])) || IsSign(rune(expression[len(expression)-1])) {
		return 0, Errors.ErrInvalidExpression
	}
	for i, value := range expression {
		if value == '(' {
			isc = i
		}
		if value == ')' {
			calc, err := Calc(expression[isc+1 : i])
			if err != nil {
				return 0, Errors.ErrMismatchedBrackets
			}
			calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
			i2 := i
			i -= len(expression[isc:i+1]) - len(calcstr)
			expression = strings.Replace(expression, expression[isc:i2+1], calcstr, 1)
		}
	}
	if countc > 1 {
		for i := 1; i < len(expression); i++ {
			value := rune(expression[i])

			if value == '*' || value == '/' {
				var imin int = i - 1
				if imin != 0 {
					for !IsSign(rune(expression[imin])) && imin > 0 {
						imin--
					}
					imin++
				}
				var imax int = i + 1
				if imax == len(expression) {
					imax--
				} else {
					for !IsSign(rune(expression[imax])) && imax < len(expression)-1 {
						imax++
					}
				}
				if imax == len(expression)-1 {
					imax++
				}
				calc, err := Calc(expression[imin:imax])
				if err != nil {
					return 0, Errors.ErrInvalidExpression
				}
				calcstr := strconv.FormatFloat(calc, 'f', 0, 64)
				i -= len(expression[isc:i+1]) - len(calcstr) - 1
				expression = strings.Replace(expression, expression[imin:imax], calcstr, 1)
			}
			if value == '+' || value == '-' || value == '*' || value == '/' {
				c = value
			}
		}
	}

	for _, value := range expression + "s" {
		switch {
		case value == ' ':
			continue
		case value > 47 && value < 58:
			b += string(value)
		case IsSign(value) || value == 's':
			if resflag {
				switch c {
				case '+':
					res += StringToFloat64(b)
				case '-':
					res -= StringToFloat64(b)
				case '*':
					res *= StringToFloat64(b)
				case '/':
					res /= StringToFloat64(b)
				}
			} else {
				resflag = true
				res = StringToFloat64(b)
			}
			b = strings.ReplaceAll(b, b, "")
			c = value

		case value == 's':
		default:
			return 0, Errors.ErrInvalidInput
		}
	}
	return res, nil
}
