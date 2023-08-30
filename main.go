package main

import "github.com/1412270/ggcache/cache"

func main() {
	opts := ServerOpts{
		ListenAddr: ":3000",
		IsLeader:   true,
	}
	server := NewServer(opts, cache.NewCache())
	server.Start()
}
