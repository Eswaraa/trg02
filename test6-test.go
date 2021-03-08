package main

import (
	"testing"
)

func TestDrive(t *testing.T){
	var testCar Car = Car{Make:"Hyundai",
		Model:"Santro 2020"
		Price: 5.6 
		Speed: 60
	}
	var expectedSpeed  int = 66
	if s:= stestCar.Drive(3); s!=expectedSpeed {
		t.Errorf("For the given acc %d Car Speed does not match expected value %d", 3, expectedSpeed)
	}


}