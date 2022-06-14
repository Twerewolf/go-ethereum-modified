package udp_resolution

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	RESOLUTION_NA   = "::1"
	SERVER_IP       = "2400:dd01:1037:0003:192:168:47:198"
	SERVER_PORT     = 2000
	SERVER_RECV_LEN = 1000

	SEAEP_ACTION_REGISTER = 111
	SEAEP_ACTION_RESOLVER = 113
	RESOLVE_TRYTIMES      = 3
)

type ServerInfo struct {
	server_addr string //128bit
	port        int
	result      int
}

//dial method and class to resolution system in SEANet

// func (u udpDialer) dial() (result, err) {

// 	dst, err := net.ResolveUDPAddr("udp", "192.0.2.1:2000")
// }

//conn, err := net.Dial("udp", "addressudp6")

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(1)
}
func Fill4(n int32) []byte { //将int32转为4bytes的slice输出
	// a := byte((n & 0xff000000) >> 24)
	// b := byte((n & 0x00ff0000) >> 16)
	// c := byte((n & 0x0000ff00) >> 8)
	// d := byte(n & 0x000000ff)
	// e := byte((n) >> 32)
	a := byte((n) >> 24) //注释掉的和此处的结果是一致的
	b := byte((n) >> 16)
	c := byte((n) >> 8)
	d := byte(n)

	fmt.Printf("%d %d %d %d\n", a, b, c, d)
	var res []byte
	res = append(res, a)
	res = append(res, b)
	res = append(res, c)
	res = append(res, d)
	return res
}
func Fill2(n int16) []byte { //将int16转为2bytes的slice输出
	c := byte((n) >> 8)
	d := byte(n)

	fmt.Printf("%d %d\n", c, d)
	var res []byte
	res = append(res, c)
	res = append(res, d)
	return res
}
func Get_random_request_id() []byte { //4Bytes
	ran_uuid := uuid.New()     //[16]byte
	requestID := ran_uuid[0:4] //获取前4B
	return requestID
}
func Get_delay_level(delayParam int) int {
	fmt.Println("delayParam:", delayParam)
	return 3
}

// func hash(char *str) uint
// {
//     unsigned int h;
//     unsigned char *p;

//     h = 0;
//     for (p = (unsigned char *)str; *p != '\0'; p++)
//         h = MULTIPLIER * h + *p;
//     return h;
// }

func Handle_register_packet(eid string, cid string, na string, delayParameter int, ttl int, flag int, tlvlen int) []byte {
	//处理好要发送的包到toWrite
	var data_buffer []byte
	//type
	data_buffer = append(data_buffer, byte(SEAEP_ACTION_REGISTER))

	//reqeustID:4
	data_buffer = append(data_buffer, Get_random_request_id()...)
	//eid:20
	eidbyte := byte(1)
	data_buffer = append(data_buffer, eidbyte)
	//cid:32bytes
	cidbyte := byte(1)
	data_buffer = append(data_buffer, cidbyte)
	//na:16,自己ip
	nabyte := byte(1)
	data_buffer = append(data_buffer, nabyte)
	//delayParameter: 1byte
	level := Get_delay_level(delayParameter)
	data_buffer = append(data_buffer, byte(level))
	//ttl 1
	ttlbyte := byte(ttl)
	data_buffer = append(data_buffer, ttlbyte)
	//flag 1
	flagbyte := byte(flag)
	data_buffer = append(data_buffer, flagbyte)
	//timestamp 4
	now := time.Now()
	timestamp1 := now.Unix() //int64
	timestampbyte := Fill4(int32(timestamp1))
	data_buffer = append(data_buffer, timestampbyte...) //append []byte use ...
	//tlv length 2bytes
	// var length int
	// if tlv != "" {
	// 	length = len(tlv)
	// }
	var tlvlength []byte
	if tlvlen == 0 {
		tlvlength = []byte{0, 0}
	}
	data_buffer = append(data_buffer, tlvlength...)
	// // tlv 0
	// tlvbyte := make([]byte, length)
	// tlvbyte = append(tlvbyte, byte(tlv))
	// //长度为0
	return data_buffer
}

//发送两个包注册
func Seaep_register_with_IP(eid string, cid string, na_ip string, delayParameter int64, tlv string, ttl uint64,
	isGlobalVisable int64, geoNeighborFlag int64, delayNeighborFlag int64, indexNeighborFlag int64) (int, error) {

	fmt.Println("Register:")
	fmt.Println("EID:", eid)
	fmt.Println("CID:", cid)
	fmt.Println("NA:", na_ip)
	//通用ipv4和ipv6，v4补齐0，转成binary ip addr
	// var networkType string
	// index := strings.Index(na_ip, ".")
	// if index != -1 { //ipv4
	// 	networkType = "udp4"
	// } else {
	// 	networkType = "udp6"
	// }

	serverAddr := SERVER_IP + ":" + strconv.Itoa(SERVER_PORT)
	// na, err := net.ResolveUDPAddr("udp6", serverAddr) 包含在了Dial的底层
	packet := Handle_register_packet(eid, cid, na_ip, int(delayParameter), int(ttl), 1, 0)
	// handle_message(data_buffer, data_len, serverlist, listnum,1,seaep_process_register_msg_g)
	conn, err := net.Dial("udp", serverAddr) //udp6 socket create
	// CheckError(err)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return -1, err
	}
	defer conn.Close()

	n, err := conn.Write(packet) //udp6 socket send packet out
	fmt.Println(n)
	CheckError(err)

	msg := make([]byte, SERVER_RECV_LEN)

	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) //2 second timeout
	n, err = conn.Read(msg)
	CheckError(err)
	output := string(msg)
	fmt.Println("Response:", output)
	return 1, err

}

func Seaep_resolve_with_eid(eid string, cid string, querytype int, delayParameter int, tlv string,
	geoNeighborFlag int, delayNeighborFlag int, indexNeighborFlag int, tlvFlag int) (string, error) {
	fmt.Println("Resolve:")
	start := time.Now()
	packet := Handle_resolve_packet(eid, cid, querytype, delayParameter, tlv, geoNeighborFlag, delayNeighborFlag, indexNeighborFlag, tlvFlag)
	// handle_resolve_packet(eid, cid,querytype,delayParameter,tlv,)

	serverAddr := SERVER_IP + ":" + strconv.Itoa(SERVER_PORT)
	conn, err := net.Dial("udp", serverAddr) //udp6 socket create
	// CheckError(err)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return "udp dial error", err
	}
	defer conn.Close()

	n, err := conn.Write(packet) //udp6 socket send packet out
	fmt.Println(n)
	CheckError(err)

	msg := make([]byte, SERVER_RECV_LEN)

	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) //2 second timeout
	n, err = conn.Read(msg)                               //需要另外解析受到的msg包
	CheckError(err)
	output := string(msg)
	fmt.Println("Response:", output)

	diff := time.Since(start)
	fmt.Println("resolve delay time: ", diff)
	return output, err
}
func Handle_resolve_packet(eid string, cid string, querytype int, delayParameter int, tlv string,
	geoNeighborFlag int, delayNeighborFlag int, indexNeighborFlag int, tlvFlag int) []byte {
	// total length: 65 Bytes
	var data_buffer []byte
	//type
	data_buffer = append(data_buffer, byte(SEAEP_ACTION_RESOLVER))
	//remote 1Byte default=0 有tlvflag则2
	t := byte(0)
	if tlvFlag == 1 {
		t = byte(2)
	}
	data_buffer = append(data_buffer, t)
	// querytype
	data_buffer = append(data_buffer, byte(querytype))
	//tlv 2B
	l := 0
	data_buffer = append(data_buffer, Fill2(int16(l))...)
	//reqeustID:4
	data_buffer = append(data_buffer, Get_random_request_id()...)
	//eid:20
	eidbyte := byte(1)
	data_buffer = append(data_buffer, eidbyte)
	//cid:32bytes
	cidbyte := byte(1)
	data_buffer = append(data_buffer, cidbyte)

	//timestamp 4
	now := time.Now()
	timestamp1 := now.Unix() //int64
	timestampbyte := Fill4(int32(timestamp1))
	data_buffer = append(data_buffer, timestampbyte...) //append []byte use ...
	//tlv length 2bytes
	// var length int
	// if tlv != "" {
	// 	length = len(tlv)
	// }
	// var tlvlength []byte
	if t == byte(0) {
		// tlvlength = []byte{0, 0}
		// do nothing
	}
	// if t == byte(2) {
	// 	data_buffer = append(data_buffer,[])
	// }

	return data_buffer
}

// func handle_messages(data_buffer []byte, serverlist []server_info, waitall int) {
// 	//建立多个socket，分别senddata到对应server address:port，select监听
// 	//FD_ZERO,FD_SET
// }
// func handle_message(data_buffer []byte, server server_info) {

// }

// net
// func ResolveUDPAddr(net, addr string) (*UDPAddr, os.Error)
// func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
// func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
// func (c *UDPConn) Read(b []byte) (int, error)
// func (c *UDPConn) Write(b []byte) (int, error)
