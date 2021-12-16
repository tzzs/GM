package GM

// sBox S盒
var sBox = [256]uint8{
	0xd6, 0x90, 0xe9, 0xfe, 0xcc, 0xe1, 0x3d, 0xb7, 0x16, 0xb6, 0x14, 0xc2, 0x28, 0xfb, 0x2c, 0x05,
	0x2b, 0x67, 0x9a, 0x76, 0x2a, 0xbe, 0x04, 0xc3, 0xaa, 0x44, 0x13, 0x26, 0x49, 0x86, 0x06, 0x99,
	0x9c, 0x42, 0x50, 0xf4, 0x91, 0xef, 0x98, 0x7a, 0x33, 0x54, 0x0b, 0x43, 0xed, 0xcf, 0xac, 0x62,
	0xe4, 0xb3, 0x1c, 0xa9, 0xc9, 0x08, 0xe8, 0x95, 0x80, 0xdf, 0x94, 0xfa, 0x75, 0x8f, 0x3f, 0xa6,
	0x47, 0x07, 0xa7, 0xfc, 0xf3, 0x73, 0x17, 0xba, 0x83, 0x59, 0x3c, 0x19, 0xe6, 0x85, 0x4f, 0xa8,
	0x68, 0x6b, 0x81, 0xb2, 0x71, 0x64, 0xda, 0x8b, 0xf8, 0xeb, 0x0f, 0x4b, 0x70, 0x56, 0x9d, 0x35,
	0x1e, 0x24, 0x0e, 0x5e, 0x63, 0x58, 0xd1, 0xa2, 0x25, 0x22, 0x7c, 0x3b, 0x01, 0x21, 0x78, 0x87,
	0xd4, 0x00, 0x46, 0x57, 0x9f, 0xd3, 0x27, 0x52, 0x4c, 0x36, 0x02, 0xe7, 0xa0, 0xc4, 0xc8, 0x9e,
	0xea, 0xbf, 0x8a, 0xd2, 0x40, 0xc7, 0x38, 0xb5, 0xa3, 0xf7, 0xf2, 0xce, 0xf9, 0x61, 0x15, 0xa1,
	0xe0, 0xae, 0x5d, 0xa4, 0x9b, 0x34, 0x1a, 0x55, 0xad, 0x93, 0x32, 0x30, 0xf5, 0x8c, 0xb1, 0xe3,
	0x1d, 0xf6, 0xe2, 0x2e, 0x82, 0x66, 0xca, 0x60, 0xc0, 0x29, 0x23, 0xab, 0x0d, 0x53, 0x4e, 0x6f,
	0xd5, 0xdb, 0x37, 0x45, 0xde, 0xfd, 0x8e, 0x2f, 0x03, 0xff, 0x6a, 0x72, 0x6d, 0x6c, 0x5b, 0x51,
	0x8d, 0x1b, 0xaf, 0x92, 0xbb, 0xdd, 0xbc, 0x7f, 0x11, 0xd9, 0x5c, 0x41, 0x1f, 0x10, 0x5a, 0xd8,
	0x0a, 0xc1, 0x31, 0x88, 0xa5, 0xcd, 0x7b, 0xbd, 0x2d, 0x74, 0xd0, 0x12, 0xb8, 0xe5, 0xb4, 0xb0,
	0x89, 0x69, 0x97, 0x4a, 0x0c, 0x96, 0x77, 0x7e, 0x65, 0xb9, 0xf1, 0x09, 0xc5, 0x6e, 0xc6, 0x84,
	0x18, 0xf0, 0x7d, 0xec, 0x3a, 0xdc, 0x4d, 0x20, 0x79, 0xee, 0x5f, 0x3e, 0xd7, 0xcb, 0x39, 0x48,
}

// FK 系统参数 System parameters
var FK = [4]uint32{
	0xA3B1BAC6, 0x56AA3350, 0x677D9197, 0xB27022DC,
}

// CK 固定参数 Fixed parameters
var CK = [32]uint32{
	0x00070e15, 0x1c232a31, 0x383f464d, 0x545b6269,
	0x70777e85, 0x8c939aa1, 0xa8afb6bd, 0xc4cbd2d9,
	0xe0e7eef5, 0xfc030a11, 0x181f262d, 0x343b4249,
	0x50575e65, 0x6c737a81, 0x888f969d, 0xa4abb2b9,
	0xc0c7ced5, 0xdce3eaf1, 0xf8ff060d, 0x141b2229,
	0x30373e45, 0x4c535a61, 0x686f767d, 0x848b9299,
	0xa0a7aeb5, 0xbcc3cad1, 0xd8dfe6ed, 0xf4fb0209,
	0x10171e25, 0x2c333a41, 0x484f565d, 0x646b7279,
}

// Sbox 输出置换 Output substitution
func Sbox(value uint8) byte {
	var x, y int
	x = (int)(value >> 4 & 0xff)
	y = (int)(value & 0x0f)
	return sBox[x*16+y]
}

// leftRotate 循环左移 cyclic left shift operation
func leftRotate(in uint32, r int) uint32 {
	return in<<r | in>>(32-r)
}

// L 线性变换 linear transformations
func L(in uint32) uint32 {
	return in ^ leftRotate(in, 2) ^ leftRotate(in, 10) ^ leftRotate(in, 18) ^ leftRotate(in, 24)
}

// L2 密钥扩展线性变换 key linear transformations
func L2(in uint32) uint32 {
	return in ^ leftRotate(in, 13) ^ leftRotate(in, 23)
}

// tau 非线性变换 Nonlinear transformations
func tau(in uint32) (res uint32) {
	var inBytes [4]byte
	inBytes[0] = byte(in>>24) & 0xff
	inBytes[1] = byte(in>>16) & 0xff
	inBytes[2] = byte(in>>8) & 0xff
	inBytes[3] = byte(in) & 0xff

	for k := 24; k >= 0; k -= 8 {
		res |= uint32(Sbox(inBytes[(24-k)/8])) << k
	}
	return res
}

// T 合成置换 Synthetic substitution functions
func T(in uint32) uint32 {
	return L(tau(in))
}

// T2 密钥扩展合成置换 Key expansion synthetic substitution functions
func T2(in uint32) uint32 {
	return L2(tau(in))
}

// round 轮函数 Round function
func round(in uint32, rk uint32) uint32 {
	var inBytes [4]byte
	inBytes[0] = byte(in>>24) & 0xff
	inBytes[1] = byte(in>>16) & 0xff
	inBytes[2] = byte(in>>8) & 0xff
	inBytes[3] = byte(in) & 0xff

	return uint32(inBytes[0]) ^ T(uint32(inBytes[1])^uint32(inBytes[2])^uint32(inBytes[3])^rk)
}

// cryption 加解密函数
func cryption(in []byte, rk []uint32) []byte {

	// check rk number

	// check in

	var x [36]uint32
	out := make([]byte, 16)

	for i := 0; i < 4; i++ {
		x[i] = uint32(in[i*4])<<24 | uint32(in[i*4+1])<<16 | uint32(in[i*4+2])<<8 | uint32(in[i*4+3])
	}

	for i := 0; i < 32; i++ {
		x[i+4] = x[i] ^ T(x[i+1]^x[i+2]^x[i+3]^rk[i])
		//fmt.Printf("rk[%v] %x X[%v] %x\n", i, rk[i], i+4, x[i+4])
	}

	for i := 0; i < 4; i++ {
		out[i*4] = byte(x[35-i] >> 24)
		out[i*4+1] = byte(x[35-i]>>16) & 0xff
		out[i*4+2] = byte(x[35-i]>>8) & 0xff
		out[i*4+3] = byte(x[35-i]) & 0xff
	}

	return out
}

/**
byte2unit32 convert 16 bytes to 4 uint32
@param in: 16 bytes
@return out: 4 uint32
*/
func byte2uint32(in []byte) (out []uint32) {
	out = make([]uint32, 4)
	for i := 0; i < 4; i++ {
		out[i] = uint32(in[i*4])<<24 | uint32(in[i*4+1])<<16 | uint32(in[i*4+2])<<8 | uint32(in[i*4+3])
	}
	return
}

/**
keyExpansion 密钥扩展函数
@param in: 128 bits original keys
@return rk: 32 extended keys
*/
func keyExpansion(in []byte) []uint32 {
	var mk = byte2uint32(in)
	keys := make([]uint32, 36)
	for i := 0; i < 4; i++ {
		keys[i] = mk[i] ^ FK[i]
	}

	for i := 0; i < 32; i++ {
		keys[i+4] = keys[i] ^ T2(keys[i+1]^keys[i+2]^keys[i+3]^CK[i])
	}

	return keys[4:]
}

/**
encryption 加密函数
@param

*/
func encryption(in []byte, rk []byte) []byte {
	var rks = keyExpansion(rk)
	return cryption(in, rks)
}

/**
decryption 解密函数
*/
func decryption(in []byte, rk []byte) []byte {
	var rks = keyExpansion(rk)
	derk := make([]uint32, 32)
	for i := 0; i < 32; i++ {
		derk[32-i-1] = rks[i]
	}

	return cryption(in, derk)
}