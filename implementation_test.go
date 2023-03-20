package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrefixToPostfix(c *C) {
	var res string
	var err error

	res, _ = PrefixToPostfix("+ + 2 2 2")
	c.Assert(res, Equals, "2 2 + 2 +")

	res, _ = PrefixToPostfix("/ * - + 10 20 30 ^ 40 ^ 50 60 70")
	c.Assert(res, Equals, "10 20 + 30 - 40 50 60 ^ ^ * 70 /")

	_, err = PrefixToPostfix("- - 2 - -")
	c.Assert(err, ErrorMatches, "incorrect expression")

	_, err = PrefixToPostfix("- 2 2 2")
	c.Assert(err, ErrorMatches, "incorrect expression")
}

func ExamplePrefixToPostfix() {
	res, err := PrefixToPostfix("+ 2 2")
	if err == nil {
		fmt.Println(res)
	} else {
		panic(err)
	}

	// Output:
	// 2 2 +
}
