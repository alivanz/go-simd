package main

import (
	"bytes"
	"io/ioutil"
	"regexp"

	"github.com/alivanz/go-simd/generator/utils"
)

type Intrinsic struct {
	Name        string
	CpuID       string
	Description string
	Operation   string
}

var (
	regIntrinsic   = regexp.MustCompile(`<intrinsic .+?</intrinsic>`)
	regName        = regexp.MustCompile(`name="(.+?)"`)
	regDescription = regexp.MustCompile(`<description>(.+?)</description>`)
	regCpuID       = regexp.MustCompile(`<CPUID>(.+?)</CPUID>`)
)

func GetIntrinsic() ([]*Intrinsic, error) {
	if err := utils.Download(
		"data.xml",
		"https://www.intel.com/content/dam/develop/public/us/en/include/intrinsics-guide/data-3-6-6.xml",
	); err != nil {
		return nil, err
	}
	raw, err := ioutil.ReadFile("data.xml")
	if err != nil {
		return nil, err
	}
	raw = bytes.ReplaceAll(raw, []byte("\n"), []byte(""))
	intrins := regIntrinsic.FindAll(raw, -1)
	out := make([]*Intrinsic, len(intrins))
	for i, part := range intrins {
		var intrin Intrinsic
		if match := regName.FindSubmatch(part); match != nil {
			intrin.Name = string(match[1])
		}
		if match := regDescription.FindSubmatch(part); match != nil {
			intrin.Description = string(match[1])
		}
		if match := regCpuID.FindSubmatch(part); match != nil {
			intrin.CpuID = string(match[1])
		}
		out[i] = &intrin
	}
	return out, nil
}
