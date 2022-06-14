package udp_resolution

import (
	"fmt"
	"math/big"
	"net"
)

func InetNtoA(ip int64) string {
	fmt.Println(ip >> 24)
	return fmt.Sprintf("%d.%d.%d.%d",
		byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4())
	return ret.Int64()
}

//单独设置v4和v6的区分方法
// func InetAtoN6(ip string) int64 {
// 	res :=big.NewInt(0)
// 	res.SetBytes(net.ParseIP(ip).To16())
// }

func IP6toInt(ipv6 string) *big.Int { //可以直接用net.Dial方法，转换不用自己完成
	IPv6Address := net.ParseIP(ipv6)
	IPv6Int := big.NewInt(0)
	IPv6Int.SetBytes(IPv6Address.To16())
	return IPv6Int
}
func main2() { //函数or方法必须开头大写才能外部调用
	// ip := "192.168.78.123"
	ip := "220.181.38.148" //baidu
	ipInt := InetAtoN(ip)

	fmt.Printf("convert string ip [%s] to int: %d\n", ip, ipInt)
	fmt.Printf("convert int ip [%d] to string: %s\n", ipInt, InetNtoA(ipInt))
}
