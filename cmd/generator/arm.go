package main

import (
	"encoding/json"
	"os"
)

type ArmIntrinsics []ArmIntrinsic

type ArmIntrinsic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetIntrinsics() (ArmIntrinsics, error) {
	if err := download(
		"intrinsics.json",
		"https://developer.arm.com/architectures/instruction-sets/intrinsics/data/intrinsics.json",
	); err != nil {
		return nil, err
	}
	f, err := os.Open("intrinsics.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var intrins ArmIntrinsics
	if err := json.NewDecoder(f).Decode(&intrins); err != nil {
		return nil, err
	}
	return intrins, nil
}

func (intrins ArmIntrinsics) Find(s string) *ArmIntrinsic {
	for _, intrin := range intrins {
		if intrin.Name == s {
			return &intrin
		}
	}
	return nil
}
