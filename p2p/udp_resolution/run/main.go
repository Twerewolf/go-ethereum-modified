package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/p2p/enode"
)

func (api *privateAdminAPI) AddPeer(url string) (bool, error) {
	// Make sure the server is running, fail otherwise
	server := api.node.Server()
	if server == nil {
		return false, ErrNodeStopped
	}
	// Try to add the url as a static peer and return
	node, err := enode.Parse(enode.ValidSchemes, url)
	if err != nil {
		return false, fmt.Errorf("invalid enode: %v", err)
	}
	server.AddPeer(node)
	return true, nil
}

// "github.com/ethereum/go-ethereum/p2p/udp_resolution"
func main() { //此部分测试通过后放入server.Start()方法中

	log.Info("test udp connect to seaep")
	// s := "sha256 this string"
	// 硬编码唯一eid
	eid := []byte{0b11111111, 0x00, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x11, 0x22}
	fmt.Println("eid: ", (eid)) //o
	fmt.Println("eid len: ", len(eid))
	//本地节点enode 160bits，前面补零，需要单独一个解析方法+
	cid := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	// 64B,拆成两部分发送
	enode := "ae795f7ed38aef3a29763216ec3673bcffa38938a77bc9811a2d56977bc495d8d93d6b4536d21b01b5f67662cea07dd89feabb0354abc42ecc33bd918a4573c2"
	for i := 0; i < len(enode); i++ {
		// charbyte := byte(enode[i : i+2])

		cid = append(cid, charbyte)
		i++
	}
	// cid = append(cid, []byte(enode)...)
	// b := []byte(s)
	fmt.Println("cid len: ", len(cid))
	fmt.Println("cid: ", cid)
	// delayParameter := 100
	// x, err := udp_resolution.Seaep_register_with_IP(eid, cid, na, delayParameter, tlv, ttl, isGlobalVisable, geoNeighborFlag, delayNeighborFlag, indexNeighborFlag)

	// if err != nil {
	// 	fmt.Fprintln(os.Stderr, err)
	// 	os.Exit(1)
	// }
	os.Exit(0)
}
