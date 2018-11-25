package main

import (
	"leeif.me/Go-fun/spider-training/domain/piaofang"
	"leeif.me/Go-fun/spider-training/infra/url"
)

func main() {
	piaofang.GetPiaofangRankInfo(url.PFangURL)
}
