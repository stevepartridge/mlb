package mlb

import (
	"fmt"
)

func (m *Mlb) log(v ...interface{}) {
	if m.Debug {
		fmt.Println(append([]interface{}{"debug:"}, v...)...)
	}
}

func ifError(err error) {
	if err != nil {
		fmt.Println("error:", err.Error())
	}
}
