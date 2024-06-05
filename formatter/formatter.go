package formatter

import "fmt"

type Formatter interface {
	Print(elements []fmt.Stringer)
}
