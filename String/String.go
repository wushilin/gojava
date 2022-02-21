package String

import (
	"fmt"
	"regexp"
	"strings"
)

type String string

func (v String) CharAt(index int) rune {
	return []rune(string(v))[index]
}

func (v String) CompareToIgnoreCase(other String) int {
	thisv := strings.ToLower(string(v))
	otherv := strings.ToLower(string(other))
	if thisv < otherv {
		return -1
	}

	if thisv == otherv {
		return 0
	}

	return 1
}

func (v String) Concat(other String) String {
	result := string(v) + string(other)
	return String(result)
}

func (v String) Contains(other String) bool {
	return v.IndexOfString(other) != -1
}

func arrEqual(rune1 []rune, rune2 []rune, rune1Start int, rune2Start int, length int) bool {
	if rune1Start < 0 || rune2Start < 0 {
		return false
	}
	if length < 0 {
		return false
	}

	// hello, hello1, 0, 0, 3
	if rune1Start > len(rune1)-1 {
		return false
	}

	if rune2Start > len(rune2)-1 {
		return false
	}

	if rune1Start+length > len(rune1) {
		return false
	}

	if rune2Start+length > len(rune2) {
		return false
	}
	if length == 0 {
		return true
	}
	rune1Sub := rune1[rune1Start:]
	rune2Sub := rune2[rune2Start:]

	for i := 0; i < length; i++ {
		if rune1Sub[i] != rune2Sub[i] {
			return false
		}
	}
	return true
}
func (v String) EndsWith(other String) bool {
	rune1 := v.ToCharArray()
	rune2 := other.ToCharArray()
	return arrEqual(rune1, rune2, len(rune1)-len(rune2), 0, len(rune2))
}

func (v String) StartsWith(other String) bool {
	return v.StartsWithFrom(other, 0)
}

func (v String) StartsWithFrom(other String, startIndex int) bool {
	rune1 := v.ToCharArray()
	rune2 := v.ToCharArray()
	return arrEqual(rune1, rune2, 0, 0, len(rune2))
}
func (v String) SubStringWithLength(start int, length int) String {
	runes := v.ToCharArray()
	if length < 0 {
		runes = runes[start:]
	} else {
		runes = runes[start : start+length]
	}
	return String(string(runes))
}

func (v String) ToString() string {
	return string(v)
}

func (v String) SubString(start int) String {
	return v.SubStringWithLength(start, -1)
}

func (v String) Format(args ...interface{}) String {
	return String(fmt.Sprintf(string(v), args...))
}

func (v String) Bytes() []byte {
	return []byte(string(v))
}

func (v String) IndexOf(ch rune) int {
	return v.IndexOfFrom(ch, 0)
}

func (v String) IndexOfFrom(ch rune, fromIndex int) int {
	for pos, char := range []rune(string(v)) {
		if ch == char && pos >= fromIndex {
			return pos
		}
	}
	return -1
}

func (v String) IndexOfString(what String) int {
	return v.IndexOfStringFrom(what, 0)
}

func (v String) IndexOfStringFrom(what String, start int) int {
	rune1 := v.ToCharArray()
	rune2 := what.ToCharArray()

	if start < 0 {
		start = 0
	}

	for i := start; i <= len(rune1)-len(rune2); i++ {
		if arrEqual(rune1, rune2, i, 0, len(rune2)) {
			return i
		}
	}
	return -1
}
func (v String) IsEmpty() bool {
	return v.Length() == 0
}

func (v String) LastIndexOf(other String) int {
	if v.Length() < other.Length() {
		return -1
	}

	rune1 := v.ToCharArray()
	rune2 := other.ToCharArray()

	for pos := v.Length() - other.Length(); pos >= 0; pos-- {
		if arrEqual(rune1, rune2, pos, 0, len(rune2)) {
			return pos
		}
	}

	return -1
}

func (v String) Length() int {
	return len(v.ToCharArray())
}

func (v String) Matches(regex string) bool {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return false
	}

	return reg.MatchString(string(v))
}

func (v String) Replace(oldchar rune, newchar rune) String {
	mydata := []rune(string(v))
	newdata := make([]rune, len(mydata))
	for i := 0; i < len(mydata); i++ {
		if mydata[i] == oldchar {
			newdata[i] = newchar
		} else {
			newdata[i] = mydata[i]
		}
	}
	return String(string(newdata))
}

func (v String) ReplaceAll(regex string, replacement String) String {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return v
	}

	return String(reg.ReplaceAllString(string(v), string(replacement)))
}

func (v String) ReplaceFirst(regex string, replacement String) String {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return v
	}

	submatches := reg.FindStringSubmatch(string(v))
	if len(submatches) == 0 {
		return v
	}

	rune1 := v.ToCharArray()
	rune2 := replacement.ToCharArray()

	firststr := submatches[0]
	index := v.IndexOfString(String(firststr))
	if index < 0 {
		panic("This should not be the case")
	}
	before := rune1[:index]
	it := rune2
	after := rune1[index+len([]rune(firststr)):]
	return String(string(before) + string(it) + string(after))

}

func convertArray(in []string) []String {
	result := make([]String, len(in))
	for idx, v := range in {
		result[idx] = String(v)
	}
	return result
}

func (v String) Split(regex string) []String {
	return v.SplitLimit(regex, -1)
}

func (v String) SplitLimit(regex string, limit int) []String {
	reg, err := regexp.Compile(regex)
	if err != nil {
		return []String{v}
	}
	return convertArray(reg.Split(string(v), limit))
}

func (v String) ToCharArray() []rune {
	return []rune(string(v))
}

func (v String) Repeat(n int) String {
	rune1 := v.ToCharArray()
	resultRune := make([]rune, len(rune1)*n)
	for i := 0; i < n; i++ {
		for j := 0; j < len(rune1); j++ {
			resultRune[i*len(rune1)+j] = rune1[j]
		}
	}
	return String(string(resultRune))
}

func ConvertArray(data []string) []String {
	result := make([]String, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = String(data[i])
	}
	return result
}

func ConvertArrayBack(data []String) []string {
	result := make([]string, len(data))
	for i := 0; i < len(data); i++ {
		result[i] = string(data[i])
	}
	return result
}

func (v String) JoinString(data []String) String {
	return v.Join(ConvertArrayBack(data))
}

func (v String) Join(data []string) String {
	builder := strings.Builder{}
	for i := 0; i < len(data); i++ {
		builder.WriteString(data[i])
		if i != len(data)-1 {
			builder.WriteString(string(v))
		}
	}
	return String(builder.String())
}
