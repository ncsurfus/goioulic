package main

import "testing"

func TestGns3License(t *testing.T) {
	license, err := getLicense("00000000", "gns3vm")
	if err != nil {
		t.Errorf("License had an error: %s", err.Error())
	}
	if license != "73635fd3b0a13ad0" {
		t.Errorf("License was incorrect got: %s, want: %s.", license, "73635fd3b0a13ad0")
	}
}

func TestRandLicense(t *testing.T) {
	license, err := getLicense("12345678", "asdgsdgsgsasdf")
	if err != nil {
		t.Errorf("License had an error: %s", err.Error())
	}
	if license != "c839f272e688e0f0" {
		t.Errorf("License was incorrect got: %s, want: %s.", license, "c839f272e688e0f0")
	}
}

func TestShortNameLicense(t *testing.T) {
	license, err := getLicense("789", "hosting")
	if err != nil {
		t.Errorf("License had an error: %s", err.Error())
	}
	if license != "6d9850f05d783319" {
		t.Errorf("License was incorrect got: %s, want: %s.", license, "6d9850f05d783319")
	}
}

func TestLongNameLicense(t *testing.T) {
	license, err := getLicense("15325325", "aLongerHostname")
	if err != nil {
		t.Errorf("License had an error: %s", err.Error())
	}
	if license != "ab5fc4bda4077261" {
		t.Errorf("License was incorrect got: %s, want: %s.", license, "ab5fc4bda4077261")
	}
}
