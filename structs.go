package main

import "fmt"

type product struct {
	name            string
	medID           int
	NDC             [11]int
	isOutterPackage bool
	packageSize     int
}

func (p product) getNDC() string {
	return fmt.Sprintf("%d%d%d%d%d-%d%d%d%d-%d%d", p.NDC[0], p.NDC[1], p.NDC[2], p.NDC[3], p.NDC[4], p.NDC[5], p.NDC[6], p.NDC[7], p.NDC[8], p.NDC[9], p.NDC[10])
}

type inventoryLine struct {
	userName            string
	userID              string
	location            string
	transactionDateTime float64
	medication          product
	numFullPackages     int
	numPartialPackages  int
}

type inventoryLinesGroup []inventoryLine

type inventory []inventoryLinesGroup
