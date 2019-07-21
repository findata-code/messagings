package main_test

import (
	"os"
	"testing"
	"richmenu_creator"
)

func TestExecShouldPanicStopIfRequiredProgramArgumentIsAreMissing(t *testing.T){

	defer func () {
		if r := recover(); r != nil {
			if r.(error).Error() != "required field are missing" {
				t.Error("expect", "required field are missing", "actual", r)
			}
		}else{
			t.Error("expect", "error happened", "actual", r)
		}
	}()

	os.Args = []string{
		"",
		"-width=2500",
		"-height=1686",
		"-selected=true",
		"-name=Home",
		"-chatBarText=Home",
		"-area=$(pwd)/areas/Home.json",
		"-image=$(pwd)/images/Home.png",
	}

	main.Exec()
}