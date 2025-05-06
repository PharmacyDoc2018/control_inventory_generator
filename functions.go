package main

import (
	"fmt"
	"strconv"
	"strings"
)

func inputNDC(p *product, NDCstring string) {
	var NDCparts []string
	NDCparts = strings.Split(NDCstring, "-")
	n := 0
	for i := 0; i < len(NDCparts); i++ {
		for _, runeValue := range NDCparts[i] {
			num, err := strconv.Atoi(string(runeValue))
			if err != nil {
				fmt.Errorf("%w", err)
			}
			p.NDC[n] = num
			n++
		}
	}
}
