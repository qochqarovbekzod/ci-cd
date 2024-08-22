package main

import "testing"

func TestManin(t *testing.T)  {
	resp:=Hello("ali")
	if resp!="hello ali"{
        t.Error("expected hello ali but got",resp)
    }
}