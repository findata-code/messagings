package model_test

import (
	"errors"
	. "richmenu_creator/model"
	"testing"
)

func TestConfigReadInputSomeRequiredProgramArgumentConfigShouldReturnErrorRequiredFieldAreMissing(t *testing.T) {
	expectedError := errors.New("required field are missing")
	inputArgs := []string{
		"",
	}

	config := Config{}
	actualError := config.Read(inputArgs)

	if expectedError.Error() != actualError.Error() {
		t.Error("expect", expectedError, "actual", actualError)
	}
}

func TestConfigReadInputCorrectProgramConfigShouldReturnNoError(t *testing.T) {
	expectedError := error(nil)
	expectedConfig := Config{
		Width:       2500,
		Height:      1686,
		Selected:    true,
		Name:        "Home",
		ChatBarText: "Home",
		AreaFile:    "$(pwd)/areas/Home.json",
		ImageFile:   "$(pwd)/images/Home.png",
	}
	inputArgs := []string{
		"",
		"-width=2500",
		"-height=1686",
		"-selected=true",
		"-name=Home",
		"-chatBarText=Home",
		"-areaFile=$(pwd)/areas/Home.json",
		"-imageFile=$(pwd)/images/Home.png",
	}

	actualConfig := Config{}
	actualError := actualConfig.Read(inputArgs)

	if expectedError != actualError {
		t.Error("expect", "no error", "actual", actualError)
	}
	if expectedConfig != actualConfig {
		t.Error("expect", expectedConfig, "actual", actualConfig)
	}
}
