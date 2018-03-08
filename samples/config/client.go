package main

import (
	"github.com/henrylee2cn/cfgo"
	tp "github.com/henrylee2cn/teleport"
)

func main() {
	cfg := tp.PeerConfig{}

	// auto create and sync config from config/config.yaml
	cfgo.MustGet("config/config.yaml", true).MustReg("cfg_cli", &cfg)

	cli := tp.NewPeer(cfg)
	defer cli.Close()

	sess, err := cli.Dial(":9090")
	if err != nil {
		tp.Fatalf("%v", err)
	}

	var reply int
	rerr := sess.Pull("/math/add?push_status=yes",
		[]int{1, 2, 3, 4, 5},
		&reply,
	).Rerror()

	if rerr != nil {
		tp.Fatalf("%v", rerr)
	}
	tp.Printf("reply: 1+2+3+4+5 = %d", reply)
}
