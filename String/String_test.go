package String

import (
	"log"
	"strings"
	"testing"

	"github.com/wushilin/gojava/common"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestString(t *testing.T) {
	testAscii := String("Hello,world,hello,new")
	test := String("我想说一句人话人民")
	if test.Length() != len([]rune(test)) {
		log.Fatal("Length wrong")
	}
	if testAscii.Length() != len(testAscii) {
		log.Fatal("Length ascii wrong")
	}

	common.AssertEq(t, test.CharAt(2), '说')
	common.AssertEq(t, testAscii.CharAt(1), 'e')
	common.AssertEq(t, test.Concat(testAscii), String(string(test)+string(testAscii)))
	common.AssertTrue(t, !test.EndsWith("1我想说一句人话人"))
	common.AssertTrue(t, test.EndsWith("人民"))

	common.AssertEq(t, test.LastIndexOf("人"), test.Length()-2)
	common.AssertEq(t, test.IndexOf('人'), strings.IndexRune(string(test), '人')/3)
	common.AssertEq(t, len(testAscii.Split(",")), 4)
	common.AssertEq(t, len(test.Split("")), test.Length())
	common.AssertTrue(t, test.Matches("人话"))
	common.AssertTrue(t, test.ReplaceFirst("人", "狗").IndexOfString("人") == 7)
	common.AssertTrue(t, test.ReplaceAll("人", "狗").IndexOf('人') == -1)
	common.AssertTrue(t, test.ReplaceAll(".", "牛") == String("牛").Repeat(test.Length()))
	common.AssertEq(t, len(testAscii.SplitLimit(",", 2)), 2)
	common.AssertEq(t, test.SubStringWithLength(test.Length()-4, 4), String("人话人民"))
	common.AssertArrEq(t, test.Bytes(), []byte(string(test)))
	common.AssertEq(t, test.Contains("人话"), true)
	testFmt := String("hello, %d")
	common.AssertEq(t, testFmt.Format(12), String("hello, 12"))

	joiner := String(";")
	arr := []string{"a", "b", "c"}
	common.AssertEq(t, joiner.Join(arr), "a;b;c")
	common.AssertTrue(t, !test.IsEmpty())
	common.AssertArrEq(t, test.ToCharArray(), []rune(string(test)))
}
