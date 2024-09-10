package main

type Value struct {
	AsciiChr []rune
	value    string
}

func (v *Value) conv() {
	if v.value == "" {
		v.value = string(v.AsciiChr)
	} else if v.AsciiChr == nil {
		v.AsciiChr = func() []rune {
			var temp []rune
			for _, b := range []byte(v.value) {
				temp = append(temp, rune(b))
			}
			return temp
		}()
	}
}
