package GM

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestSbox(t *testing.T) {
	type args struct {
		value uint8
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			name: "case1",
			args: args{
				value: '\xef',
			},
			want: '\x84',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sbox(tt.args.value); got != tt.want {
				t.Errorf("Sbox() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leftRotate(t *testing.T) {
	type args struct {
		in uint32
		r  int
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "01",
			args: args{
				in: 0xeeff0000,
				r:  4,
			},
			want: 0xeff0000e,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftRotate(tt.args.in, tt.args.r); got != tt.want {
				t.Errorf("leftRotate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestL(t *testing.T) {
	type args struct {
		in uint32
	}
	var tests = []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				in: 0x84848484,
			},
			want: 0x12121212,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := L(tt.args.in); got != tt.want {
				t.Errorf("L() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestL2(t *testing.T) {
	type args struct {
		in uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				in: 0x84848484,
			},
			want: 0x56565656,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := L2(tt.args.in); got != tt.want {
				t.Errorf("L2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tau(t *testing.T) {
	type args struct {
		in uint32
	}
	tests := []struct {
		name    string
		args    args
		wantRes uint32
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				in: 0xefefefef,
			},
			wantRes: 0x84848484,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := tau(tt.args.in); gotRes != tt.wantRes {
				t.Errorf("tau() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestT(t *testing.T) {
	type args struct {
		in uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := T(tt.args.in); got != tt.want {
				t.Errorf("T() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestT2(t *testing.T) {
	type args struct {
		in uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := T2(tt.args.in); got != tt.want {
				t.Errorf("T2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_round(t *testing.T) {
	type args struct {
		in uint32
		rk uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := round(tt.args.in, tt.args.rk); got != tt.want {
				t.Errorf("round() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cryption(t *testing.T) {
	type args struct {
		in []byte
		rk []uint32
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := cryption(tt.args.in, tt.args.rk); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("cryption() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_byte2uint32(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name    string
		args    args
		wantOut []uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := byte2uint32(tt.args.in); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("byte2uint32() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_keyExpansion(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name string
		args args
		want []uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := keyExpansion(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("keyExpansion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_encryption(t *testing.T) {
	type args struct {
		in []byte
		rk []byte
	}
	var in, _ = hex.DecodeString("0123456789ABCDEFFEDCBA9876543210")
	var want, _ = hex.DecodeString("681EDF34D206965E86B3E94F536E4246")
	//fmt.Printf("%x %x\n", in, want)
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				in: in,
				rk: in,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encryption(tt.args.in, tt.args.rk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("encryption() = %x, want %x", got, tt.want)
			}
		})
	}
}

func Test_decryption(t *testing.T) {
	type args struct {
		in []byte
		rk []byte
	}
	var in, _ = hex.DecodeString("681EDF34D206965E86B3E94F536E4246")
	var want, _ = hex.DecodeString("0123456789ABCDEFFEDCBA9876543210")
	//fmt.Printf("%x %x\n", in, want)
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				in: in,
				rk: want,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decryption(tt.args.in, tt.args.rk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decryption() = %v, want %v", got, tt.want)
			}
		})
	}
}
