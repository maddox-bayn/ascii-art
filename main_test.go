package main

import (
	"ascii-art/functions"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestLoadBanner(t *testing.T) {
	banner, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	sliceBanner := strings.Split(string(banner), "\n")

	table, err := functions.LoadBanner("standard.txt")
	if err != nil {
		log.Fatal("error loading banner", err)
	}

	//	start :=    ()

	if !reflect.DeepEqual(table[0], sliceBanner[1:9]) {
		t.Errorf("LoadBanner()= %#v want = %#v", table[0], sliceBanner[1:9])
	}
}
