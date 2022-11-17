package fizzbuzz

import "fmt"

type Command struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func (c Command) ToString() string {
	template := `{int1: %d, int2: %d, limit: %d, str1: %s, str2 : %s }`
	return fmt.Sprintf(template, c.Int1, c.Int2, c.Limit, c.Str1, c.Str2)
}
