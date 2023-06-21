package main

import (
	"encoding/json"
	"net/http"
)

type ArmIntrinsics []ArmIntrinsic

type ArmIntrinsic struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func GetIntrinsics() (ArmIntrinsics, error) {
	resp, err := http.Get("https://developer.arm.com/architectures/instruction-sets/intrinsics/data/intrinsics.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var intrins ArmIntrinsics
	if err := json.NewDecoder(resp.Body).Decode(&intrins); err != nil {
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
