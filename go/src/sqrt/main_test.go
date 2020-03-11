package main

import "testing"

func TestRocksSqrt(t *testing.T) {
    result := rocksSqrt()

    if result != "Code.education Rocks! " {
        t.Error("Incorrect rocksSqrt result")
    }
}