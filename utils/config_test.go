// Copyright (c) 2016 sekisoft
// See License.txt

package utils

import "testing"

func TestLoadConfig(t *testing.T) {
	LoadConfig()

	if Cfg == nil {
		t.Fatal("Should have loaded")
	}
}
