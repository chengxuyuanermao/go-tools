package pprof

import "github.com/pkg/profile"

func TestNew() {
	for i := 0; i < 100; i++ {
		repeat(generate(100), 100)
	}
	// pkg/profile 封装好的
	//defer profile.Start().Stop()
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
}
