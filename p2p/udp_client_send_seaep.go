package p2p

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	// "github.com/google/uuid"
)

const (
	// RESOLUTION_NA = "::1"
	// SERVER_IP       = "2400:dd01:1037:0003:192:168:47:198"  "0.0.0.0"

	SERVER_IP = "2400:dd01:101b:9:f120::34"
	// SERVER_IP = "fe80::215:5dff:fea2:630c" //本机ip
	// SERVER_PORT     = 2000 30000
	SERVER_PORT = 10061

	SERVER_RECV_LEN = 1000

	SEAEP_ACTION_REGISTER = 111
	SEAEP_ACTION_RESOLVER = 113
	RESOLVE_TRYTIMES      = 3
)

// func main() {
// 	enodestr := "8d2d99c7906097df60681fac79f2f5a622e0907374a8c5884336848bfe28041e16b8dc0cef465070895ff59b3bbd3950dcbb21d7500797b04bdf3d24f3c6d6c0"
// 	Seaep_Register_and_Resolve(enodestr)
// }

func Byte2int(b []byte) int16 {
	bin_buf := bytes.NewBuffer(b)
	var x int16
	binary.Read(bin_buf, binary.BigEndian, &x)
	fmt.Println(x)
	return x
}

func InetAtoN(ip string) int64 {
	ret := big.NewInt(0)
	ret.SetBytes(net.ParseIP(ip).To4()) // 可以parse ipv4或ipv6或IPv4-mapped IPv6
	return ret.Int64()
}

func InetNtoA(ip int64) string {
	// fmt.Println(ip>>24, byte(ip>>16), byte(ip>>8), ip)
	return fmt.Sprintf("%d.%d.%d.%d", byte(ip>>24), byte(ip>>16), byte(ip>>8), byte(ip))
}
func IpbytetoA(ipbyte []byte) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		ipbyte[0], ipbyte[1], ipbyte[2], ipbyte[3])
}
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)       //从conn里面得到localAddr()
	fmt.Println("localAddr:Port ", localAddr.String()) //ip+port
	ip = strings.Split(localAddr.String(), ":")[0]     //ip
	return
}
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
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

	// fmt.Printf("%d %d %d %d\n", a, b, c, d)
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
func TwoByte2Int(b []byte) int16 {
	bin_buf := bytes.NewBuffer(b)
	var x int16
	binary.Read(bin_buf, binary.BigEndian, &x)
	// fmt.Println(x)
	return x
}

func Get_random_request_id() []byte { //4Bytes
	// ran_uuid := uuid.New() //[16]byte
	ran_uuid := []byte("ae795f7ed38aef3a29763216ec3673bcffa38938a77bc9811a2d56977bc495d8d93d6b4536d21b01b5f67662cea07dd89feabb0354abc42ecc33bd918a4573c2")
	requestID := ran_uuid[0:4] //获取前4B
	fmt.Println("requestID: ", requestID)
	return requestID
}

func Get_delay_level(delayParam int) int {
	fmt.Println("delayParam:", delayParam)
	switch delayParam {
	case 100:
		fmt.Println("delayLevel:", 1)
		return 1
	case 50:
		fmt.Println("delayLevel:", 2)
		return 2
	case 10:
		fmt.Println("delayLevel:", 3)
		return 3
	}
	return 3
}

// 分割线输入
func DevideLine(input string) {
	for i := 0; i < 30; i++ {
		fmt.Printf("-")
	}
	fmt.Printf(input)
	for i := 0; i < 30; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

//main暂存
func Seaep_Register_and_Resolve(enodeid string) []string {
	DevideLine("udp_client_send_seaep start")
	eid := []byte{222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222, 222}
	eidstr := string(eid)

	// enode := "enode://ae795f7ed38aef3a29763216ec3673bcffa38938a77bc9811a2d56977bc495d8d93d6b4536d21b01b5f67662cea07dd89feabb0354abc42ecc33bd918a4573c2@127.0.0.1:30303"
	// enodeid := "ae795f7ed38aef3a29763216ec3673bcffa38938a77bc9811a2d56977bc495d8d93d6b4536d21b01b5f67662cea07dd89feabb0354abc42ecc33bd918a4573c2"

	// fmt.Println("enodeid len: ", len(enodeid)) //128
	// fmt.Println("eidstr", eidstr)
	//注册
	DevideLine("register start")
	// var cidsingle []byte
	// for i := 0; i < 32; i++ {
	// 	cidsingle = append(cidsingle, byte(1))
	// }
	// fmt.Println(cidsingle)
	// cidsingle_str := hex.EncodeToString(cidsingle)
	// fmt.Println(cidsingle_str)
	// na_local, _ := GetOutBoundIP()

	// Seaep_register_single(eidstr, cidsingle_str, na_local, 100, "", 1, 0, 0, 0, 0)
	Seaep_register_with_IP_enode(eidstr, enodeid, 100, "5", 1, 0, 0, 0, 0) //参数"5"没有用到

	//查询使用空cid

	cidstr := ""
	fmt.Println("cidstr", cidstr)
	tlv := []byte{5, 1, 1}
	// tlv := []byte{200, 1, 1}
	tlvstr := string(tlv)

	DevideLine("resolve start")
	//解析查询，查询类型4
	list, err := Seaep_resolve_with_tag_eid(eidstr, cidstr, 4, 100, tlvstr, 0, 0, 0, 1)
	CheckError(err)

	var map_enode_ip map[string]string
	map_enode_ip = Get_enodeip_from_list(list)
	var StaticNodeList []string
	StaticNodeList = Get_enodeip_str(map_enode_ip)
	//得到所有keys,即NA
	// keys := make([]string, 0, len(map_enode_ip))
	// for j, k := range StaticNodeList {
	// 	fmt.Println("ip: ", j)
	// 	fmt.Println("enode: ", k)
	// }
	return StaticNodeList
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

// 处理单个注册，测试
func Handle_register_single_packet(eid string, cid string, na string, delayParameter int, ttl int, flag int, tlvlen int) []byte {
	var data_buffer []byte
	//type 1
	typebyte := byte(SEAEP_ACTION_REGISTER)
	data_buffer = append(data_buffer, typebyte)
	fmt.Println("typebyte: ", typebyte)
	//reqeustID:4
	req := Get_random_request_id()
	data_buffer = append(data_buffer, req...)
	fmt.Println("Get_random_request_id: ", req)

	//eid:20
	eidbyte := []byte(eid)
	data_buffer = append(data_buffer, eidbyte...)
	fmt.Println("eidbyte: ", eidbyte)
	//cid:32bytes
	cidbyte, _ := hex.DecodeString(cid)
	// cidbyte := []byte(cid)
	data_buffer = append(data_buffer, cidbyte...)
	fmt.Println("cidbyte: ", cidbyte)
	//na:16,自己ip，先转数字然后转byte
	naInt := InetAtoN(na)
	// fmt.Println("naInt:", naInt)
	nabyte := Fill4(int32(naInt))
	// fmt.Println("len(nabyte):", len(nabyte))
	var zeros []byte

	for i := 0; i < 12; i++ {
		zeros = append(zeros, byte(0))
	}
	nabyte = append(zeros, nabyte...)
	data_buffer = append(data_buffer, nabyte...)
	fmt.Println("nabyte: ", nabyte)
	//delayParameter: 1byte 时延等级
	level := Get_delay_level(delayParameter)
	data_buffer = append(data_buffer, byte(level))
	fmt.Println("delay level: ", byte(level))
	//ttl 1 长期有效全一FF
	ttlbyte := byte(ttl)
	data_buffer = append(data_buffer, ttlbyte)

	//flag 1 ，0不移动1移动
	flagbyte := byte(flag)
	data_buffer = append(data_buffer, flagbyte)

	//timestamp 4，从19700101开始的秒数
	now := time.Now()
	timestamp1 := now.Unix() //int64
	// fmt.Printf("timestamp:")
	timestampbyte := Fill4(int32(timestamp1))
	data_buffer = append(data_buffer, timestampbyte...) //append []byte use ...
	fmt.Println("timestamp: ", timestampbyte)
	// tlv length tlv部分的总长度，单位byte，长度2bytes

	var tlvlength []byte
	if tlvlen == 0 {
		tlvlength = []byte{0, 0}
	} else {
		lenbyte := Fill2(int16(tlvlen))
		tlvlength = append(tlvlength, lenbyte...) //2Bytes
	}
	data_buffer = append(data_buffer, tlvlength...)

	// fmt.Println("data_buffer: ", data_buffer)
	// // tlv 共2B
	// tag 1B, length of value 1B 单位字节, value NB
	// tlvbyte := make([]byte, length)
	if tlvlen == 3 {
		tag := byte(200) //手动设置一个tag？
		length := byte(1)
		value := byte(1)
		data_buffer = append(data_buffer, tag)
		data_buffer = append(data_buffer, length)
		data_buffer = append(data_buffer, value)
	}
	// tlvbyte = append(tlvbyte, byte(tlv))
	// //长度为0
	fmt.Println("data_buffer: ", data_buffer)
	fmt.Println("data_buffer length: ", len(data_buffer))
	return data_buffer
}

// 处理成注册包提供发送
func Handle_register_packet(NO int, eid string, cid string, na string, delayParameter int, ttl int, flag int, tlvlen int) []byte {
	//处理好要发送的包到toWrite
	var data_buffer []byte
	//type 1
	typebyte := byte(SEAEP_ACTION_REGISTER)
	data_buffer = append(data_buffer, typebyte)
	fmt.Println("typebyte: ", typebyte)
	//reqeustID:4
	req := Get_random_request_id()
	data_buffer = append(data_buffer, req...)
	fmt.Println("Get_random_request_id: ", req)

	//eid:20
	eidbyte := []byte(eid)
	data_buffer = append(data_buffer, eidbyte...)
	fmt.Println("eidbyte: ", eidbyte)
	//cid:32bytes
	cidbyte, _ := hex.DecodeString(cid)
	// cidbyte := []byte(cid)
	data_buffer = append(data_buffer, cidbyte...)
	fmt.Println("cidbyte: ", cidbyte)
	//na:16,自己ip，先转数字然后转byte
	naInt := InetAtoN(na)
	// fmt.Println("naInt:", naInt)
	nabyte := Fill4(int32(naInt))
	// fmt.Println("len(nabyte):", len(nabyte))
	var zeros []byte
	if len(nabyte) == 4 {
		switch NO {
		case 1:
			for i := 0; i < 12; i++ {
				zeros = append(zeros, byte(0))
			}
		case 2: //将第9位赋值1，用于标记是cid的第二部分
			for i := 0; i < 8; i++ {
				zeros = append(zeros, byte(0))
			}
			zeros = append(zeros, byte(1))
			for i := 9; i < 12; i++ {
				zeros = append(zeros, byte(0))
			}
		}
	}
	nabyte = append(zeros, nabyte...)
	data_buffer = append(data_buffer, nabyte...)
	fmt.Println("nabyte: ", nabyte)
	//delayParameter: 1byte 时延等级
	level := Get_delay_level(delayParameter)
	data_buffer = append(data_buffer, byte(level))
	fmt.Println("delay level: ", byte(level))
	//ttl 1 长期有效全一FF
	ttlbyte := byte(ttl)
	data_buffer = append(data_buffer, ttlbyte)

	//flag 1 ，0不移动1移动
	flagbyte := byte(flag)
	data_buffer = append(data_buffer, flagbyte)

	//timestamp 4，从19700101开始的秒数
	now := time.Now()
	timestamp1 := now.Unix() //int64
	// fmt.Printf("timestamp:")
	timestampbyte := Fill4(int32(timestamp1))
	data_buffer = append(data_buffer, timestampbyte...) //append []byte use ...
	fmt.Println("timestamp: ", timestampbyte)
	// tlv length tlv部分的总长度，单位byte，长度2bytes

	var tlvlength []byte
	if tlvlen == 0 {
		tlvlength = []byte{0, 0}
	} else {
		lenbyte := Fill2(int16(tlvlen))           //查看c版源码长度要加2
		tlvlength = append(tlvlength, lenbyte...) //2Bytes
	}
	data_buffer = append(data_buffer, tlvlength...)

	// fmt.Println("data_buffer: ", data_buffer)
	// // tlv 共2B
	// tag 1B, length of value 1B 单位字节, value NB
	// tlvbyte := make([]byte, length)
	tag := byte(5) //手动设置一个tag？
	length := byte(1)
	value := byte(1)
	data_buffer = append(data_buffer, tag)
	data_buffer = append(data_buffer, length)
	data_buffer = append(data_buffer, value)
	// tlvbyte = append(tlvbyte, byte(tlv))
	// //长度为0
	fmt.Println("data_buffer: ", data_buffer)
	fmt.Println("data_buffer length: ", len(data_buffer))
	return data_buffer
}

// 处理注册接收到的响应包，打印响应信息
func Handle_register_response_packet(res []byte) {
	if res[0] != byte(112) {
		fmt.Println("Err: Not register response packet")
	}
	DevideLine("register response type correct!")

	requestid := res[1:5]
	fmt.Println("response requestid: ", requestid)
	var status string
	// fmt.Println("status response in byte: ", res[5])
	switch res[5] {
	case byte(1):
		status = "register succeed 注册成功"
	case byte(2):
		status = "parameter error"
	case byte(3):
		status = "internal error内部错误"
	case byte(4):
		status = "storage full error"
	case byte(5):
		status = "other error"
	default:
		status = "wrong response status"
	}
	fmt.Println("status: ", status)
	timestamp := res[6:10]
	fmt.Println("timestamp: ", timestamp)
}

func Seaep_register_single(eid string, cid string, na_ip string, delayParameter int64, tlv string, ttl uint64,
	isGlobalVisable int64, geoNeighborFlag int64, delayNeighborFlag int64, indexNeighborFlag int64) ([]byte, error) {
	// fmt.Println("Register:")
	DevideLine("Seaep_register_single")
	fmt.Println("EID:", eid)
	fmt.Println("CID:", cid)
	// fmt.Println("CID byte: ", cid)
	fmt.Println("NA:", na_ip)

	serverAddr := "[" + SERVER_IP + "]" + ":" + strconv.Itoa(SERVER_PORT)
	fmt.Println("serverAddr: ", serverAddr)
	//不移动，tlv长度0
	packet := Handle_register_single_packet(eid, cid, na_ip, int(delayParameter), int(ttl), 0, 0)
	conn, err := net.Dial("udp", serverAddr) //udp6 socket create
	localaddr := conn.LocalAddr().String()
	fmt.Println("local address ipv6: ", localaddr)
	// CheckError(err)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return []byte("err"), err
	}

	defer conn.Close()

	n, err := conn.Write(packet) //udp6 socket send packet out
	fmt.Println("udp write packet size: ", n)
	CheckError(err)

	msg := make([]byte, SERVER_RECV_LEN)

	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) //2 second timeout
	n, err = conn.Read(msg)
	msg = msg[:n]
	CheckError(err)
	fmt.Println("Response length:", n)
	fmt.Println("Response:", msg)
	Handle_register_response_packet(msg)
	// output := string(msg)
	return msg, err
}

func Seaep_register_with_IP_enode(eid string, enode string, delayParameter int64, tlv string, ttl uint64,
	isGlobalVisable int64, geoNeighborFlag int64, delayNeighborFlag int64, indexNeighborFlag int64) {
	DevideLine("Seaep_register_with_IP_enode")
	fmt.Println("enode register:", enode)
	//enode分成2部分
	na_ip, _ := GetOutBoundIP()
	// na_ip := "111.111.111.111"
	// na_ip := "233.233.233.233"
	cid1 := enode[:64]
	cid2 := enode[64:]
	Seaep_register_with_IP(1, eid, cid1, na_ip, delayParameter, tlv, ttl, isGlobalVisable, geoNeighborFlag, delayNeighborFlag, indexNeighborFlag)
	Seaep_register_with_IP(2, eid, cid2, na_ip, delayParameter, tlv, ttl, isGlobalVisable, geoNeighborFlag, delayNeighborFlag, indexNeighborFlag)

}

//发送两个包注册一个enode
func Seaep_register_with_IP(NO int, eid string, cid string, na_ip string, delayParameter int64, tlv string, ttl uint64,
	isGlobalVisable int64, geoNeighborFlag int64, delayNeighborFlag int64, indexNeighborFlag int64) ([]byte, error) {

	fmt.Println("Register:")
	fmt.Println("EID:", eid)
	fmt.Println("CID:", cid)
	fmt.Println("NA:", na_ip)

	serverAddr := "[" + SERVER_IP + "]" + ":" + strconv.Itoa(SERVER_PORT)
	fmt.Println("serverAddr: ", serverAddr)
	// na, err := net.ResolveUDPAddr("udp6", serverAddr) 包含在了Dial的底层
	//不可移动，倒数第二位是0
	packet := Handle_register_packet(NO, eid, cid, na_ip, int(delayParameter), int(ttl), 0, 3)
	// fmt.Println("packet size: ", len(packet))
	// handle_message(data_buffer, data_len, serverlist, listnum,1,seaep_process_register_msg_g)

	// socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
	// 	IP:   net.IPv4(0, 0, 0, 0),
	// 	Port: 30000,
	// })
	conn, err := net.Dial("udp", serverAddr) //udp6 socket create
	localaddr := conn.LocalAddr().String()
	fmt.Println("local address ipv6: ", localaddr)
	// CheckError(err)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return []byte("err"), err
	}

	defer conn.Close()

	n, err := conn.Write(packet) //udp6 socket send packet out
	fmt.Println("udp write packet size: ", n)
	CheckError(err)

	msg := make([]byte, SERVER_RECV_LEN)

	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) //2 second timeout
	n, err = conn.Read(msg)
	msg = msg[:n]
	CheckError(err)
	fmt.Println("Response length:", n)
	fmt.Println("Response:", msg)
	Handle_register_response_packet(msg)
	// output := string(msg)
	return msg, err
}

// 发送解析请求，使用tag进行查询，返回相应包数组
func Seaep_resolve_with_tag_eid(eid string, cid string, querytype int, delayParameter int, tlv string,
	geoNeighborFlag int, delayNeighborFlag int, indexNeighborFlag int, tlvFlag int) ([][]byte, error) {
	fmt.Println("Resolve:")
	start := time.Now()
	packet := Handle_resolve_packet(eid, cid, querytype, delayParameter, tlv, geoNeighborFlag, delayNeighborFlag, indexNeighborFlag, tlvFlag)
	// handle_resolve_packet(eid, cid,querytype,delayParameter,tlv,)
	var EidCidNaTLVlist [][]byte
	serverAddr := "[" + SERVER_IP + "]" + ":" + strconv.Itoa(SERVER_PORT)
	DevideLine("Connection dial")
	conn, err := net.Dial("udp", serverAddr) //udp6 socket create
	// CheckError(err)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
		return EidCidNaTLVlist, err
	}
	defer conn.Close()

	n, err := conn.Write(packet) //udp6 socket send packet out
	fmt.Println(n)
	CheckError(err)

	msg := make([]byte, SERVER_RECV_LEN)

	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) //2 second timeout
	DevideLine("Read")
	n, err = conn.Read(msg) //需要另外解析受到的msg包
	CheckError(err)
	fmt.Println("接收响应包长度: ", n)
	// output := string(msg)
	fmt.Println("Response:", msg)
	diff := time.Since(start)
	fmt.Println("resolve delay time: ", diff)

	EidCidNaTLVlist = Handle_resolve_response_packet(msg)
	// Get_enodeip_from_list(resolve_response_list)
	return EidCidNaTLVlist, err
}

// 生成发送到server的查询请求包，包含tlv
func Handle_resolve_packet(eid string, cid string, querytype int, delayParameter int, tlv string,
	geoNeighborFlag int, delayNeighborFlag int, indexNeighborFlag int, tlvFlag int) []byte {
	// total length: 65 Bytes
	var data_buffer []byte
	//type
	data_buffer = append(data_buffer, byte(SEAEP_ACTION_RESOLVER))
	// remote 1Byte default=0 有tlvflag则2
	// 2：普通，结果含自定义属性信息；
	t := byte(0)
	if tlvFlag == 1 {
		t = byte(2)
	}
	data_buffer = append(data_buffer, t)
	// querytype 4：tag查EID+CID+IP
	data_buffer = append(data_buffer, byte(querytype))
	//tlv 2B
	tlvbytetemp := []byte(tlv)
	tlvlen := len(tlvbytetemp)
	data_buffer = append(data_buffer, Fill2(int16(tlvlen))...)
	//reqeustID:4
	data_buffer = append(data_buffer, Get_random_request_id()...)

	//eid:20
	eidbyte := []byte(eid)
	data_buffer = append(data_buffer, eidbyte...)
	fmt.Println("eidbyte: ", eidbyte)
	//cid:32bytes
	var cidbyte []byte
	if cid != "" {
		cidbyte, _ = hex.DecodeString(cid)
	} else if cid == "" {
		for i := 0; i < 32; i++ {
			cidbyte = append(cidbyte, byte(0))
		}
	}
	// cidbyte := []byte(cid)
	data_buffer = append(data_buffer, cidbyte...)
	fmt.Println("cidbyte: ", cidbyte)

	//timestamp 4
	now := time.Now()
	timestamp1 := now.Unix() //int64
	timestampbyte := Fill4(int32(timestamp1))
	data_buffer = append(data_buffer, timestampbyte...) //append []byte use ...
	tlvbyte := []byte(tlv)
	tag := tlvbyte[0] //手动设置一个tag
	length := tlvbyte[1]
	value := tlvbyte[2]
	data_buffer = append(data_buffer, tag)
	data_buffer = append(data_buffer, length)
	data_buffer = append(data_buffer, value)

	fmt.Println("data_buffer: ", data_buffer)
	fmt.Println("data_buffer length: ", len(data_buffer))
	return data_buffer
}

// 处理查询响应，根据响应信息解析出各个enode+ip组合
func Handle_resolve_response_packet(res []byte) [][]byte {
	if res[0] != byte(114) {
		fmt.Println("Err: Not resolve response packet")
	}
	fmt.Println(res[0])
	var status string
	switch res[1] {
	case byte(1):
		status = "resolve succeed"
	case byte(0):
		status = "resolve failed, timeout!"
	default:
		status = "wrong response status 未设置的接收状态"
	}
	fmt.Println("status: ", status)
	//querytype1
	querytype := res[2]
	fmt.Println("querytype, 4 for TAG->(EID+CID+NA): ", querytype)
	//format1
	foramt := res[3]
	fmt.Println("foramt, 2 for TAG included: ", foramt)
	//more
	more := res[4]
	fmt.Println("more: ", more)
	//request_id
	requestid := res[5:9]
	fmt.Println("requestid: ", requestid)
	//timestamp
	timestamp := res[9:13]
	fmt.Println("timestamp: ", timestamp)
	//num
	num := res[13:15]
	fmt.Println("返回名字数量未解析: ", num)
	ecnNum := TwoByte2Int(num)
	// ecnNum, _ := strconv.ParseInt(string(num), 0, 64)
	fmt.Println("返回名字数量解析后: ", ecnNum)

	// fmt.Println(d)
	// ecnNum:=bytetoint(num)
	//eidcidna 68Bytes
	now := 15
	var ecnList [][]byte
	for i := 0; i < int(ecnNum); i++ { //一个组合的长度73
		eidcidna := res[now : now+68]
		now += 68
		fmt.Println("eidcidna: ", eidcidna)
		ecnList = append(ecnList, eidcidna)
		tlvlen := res[now : now+2]
		// fmt.Println("tlvlen in byte: ", string(tlvlen))
		now += 2
		tlvlenInt := Byte2int(tlvlen)
		// tlvlenInt, _ := strconv.ParseInt(string(tlvlen), 16, 16)
		fmt.Println("tlvlenInt: ", tlvlenInt)
		// tlv := res[now : now+int(tlvlenInt)]
		now += int(tlvlenInt) //3
		// ecnList = append(ecnList, tlv)

	}
	return ecnList
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

//得到ip-enode映射，节点可以直接用return的map发起连接
func Get_enodeip_from_list(list [][]byte) map[string]string {
	//得到ip:cid组合
	DevideLine("Get_enodeip_from_list")

	n := len(list)
	fmt.Println("len(list):", n)

	m := make(map[string]string)
	for i := 0; i < n; i++ {
		fmt.Printf("len(list[%d]):%d\n", i, len(list[i]))
		fmt.Println(list[i])
		now := 20
		cidbyte := list[i][now : now+32]
		fmt.Println("cidbyte: ", cidbyte)
		now += 32
		nabyte := list[i][now : now+16]
		fmt.Println("nabyte: ", nabyte)
		cid := hex.EncodeToString(cidbyte)
		//na先转int然后加:变ip
		NO := nabyte[8]
		fmt.Println("NO: ", NO)
		ipbyte := nabyte[12:16]
		fmt.Println("byte(ip): ", ipbyte)
		na := IpbytetoA(ipbyte)     //string
		if value, ok := m[na]; ok { //m[na]不为空，存在
			if len(value) == 128 {
				continue
			}
			fmt.Println("已有enode部分： ", value)
			if NO == byte(0) { // NO等于0时说明是enode1段
				m[na] = cid + m[na]
			} else if NO == byte(1) { // NO等于1说明是enode2段
				m[na] = m[na] + cid
			}
		} else { //empty then add cid
			m[na] = cid
		}
	}
	return m
	// fmt.Println(m])
}

//使用eid查询cid,不会获得ip
// 需要发起请求,等待获取消息
// func Seaep_resolve_with_eid() {

// 	Handle_resolve_response_packet()
// }
func Get_enodeip_str(m map[string]string) []string {
	// n := len(m)
	var list []string
	for j, k := range m {
		fmt.Println("ip: ", j)
		fmt.Println("enode: ", k)
		x := "enode://" + k + "@" + j + ":30303"
		list = append(list, x)
	}
	for _, y := range list {
		fmt.Println(y)
	}
	return list
}
