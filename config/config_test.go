package config

import (
	"fmt"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	// wantPort := "3333"
	// t.Setenv("PORT", fmt.Sprintf(wantPort))

	// got, err := New()
	// if err != nil {
	// 	t.Fatalf("cannnot create config: %v", err)
	// }
	// intPort, err := strconv.Atoi(wantPort)
	// if err != nil {
	// 	t.Errorf("cannot convert %s to int", wantPort)
	// }
	// if got.Port != intPort {
	// 	t.Errorf("want %d, but %d", intPort, got.Port)
	// }
	// wantEnv := "dev"
	// if got.Env != wantEnv {
	// 	t.Errorf("want %s, but %s", wantEnv, got.Env)
	// }
	wantPort := 3333
	stringPort := strconv.Itoa(wantPort)
	t.Setenv("PORT", fmt.Sprintf(stringPort))

	got, err := New()
	if err != nil {
		t.Fatalf("cannnot create config: %v", err)
	}
	if got.Port != wantPort {
		t.Errorf("want %d, but %d", wantPort, got.Port)
	}
	wantEnv := "dev"
	if got.Env != wantEnv {
		t.Errorf("want %s, but %s", wantEnv, got.Env)
	}
}
