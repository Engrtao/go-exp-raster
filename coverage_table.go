// Copyright 2011 The draw2d Authors. All rights reserved.
// created: 27/05/2011 by Laurent Le Goff
package raster

var SUBPIXEL_OFFSETS_SAMPLE_8 = [8]float64{
	5.0 / 8,
	0.0 / 8,
	3.0 / 8,
	6.0 / 8,
	1.0 / 8,
	4.0 / 8,
	7.0 / 8,
	2.0 / 8,
}

var SUBPIXEL_OFFSETS_SAMPLE_8_FIXED = [8]Fix{
	Fix(SUBPIXEL_OFFSETS_SAMPLE_8[0] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_8[1] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_8[2] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_8[3] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_8[4] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_8[5] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_8[6] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_8[7] * FIXED_FLOAT_COEF),
}

var SUBPIXEL_OFFSETS_SAMPLE_16 = [16]float64{
	1.0 / 16,
	8.0 / 16,
	4.0 / 16,
	15.0 / 16,
	11.0 / 16,
	2.0 / 16,
	6.0 / 16,
	14.0 / 16,
	10.0 / 16,
	3.0 / 16,
	7.0 / 16,
	12.0 / 16,
	0.0 / 16,
	9.0 / 16,
	5.0 / 16,
	13.0 / 16,
}

var SUBPIXEL_OFFSETS_SAMPLE_16_FIXED = [16]Fix{
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[0] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[1] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[2] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[3] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[4] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[5] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[6] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[7] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[8] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[9] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[10] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[11] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[12] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[13] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[14] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_16[15] * FIXED_FLOAT_COEF),
}

var SUBPIXEL_OFFSETS_SAMPLE_32 = [32]float64{
	28.0 / 32,
	13.0 / 32,
	6.0 / 32,
	23.0 / 32,
	0.0 / 32,
	17.0 / 32,
	10.0 / 32,
	27.0 / 32,
	4.0 / 32,
	21.0 / 32,
	14.0 / 32,
	31.0 / 32,
	8.0 / 32,
	25.0 / 32,
	18.0 / 32,
	3.0 / 32,
	12.0 / 32,
	29.0 / 32,
	22.0 / 32,
	7.0 / 32,
	16.0 / 32,
	1.0 / 32,
	26.0 / 32,
	11.0 / 32,
	20.0 / 32,
	5.0 / 32,
	30.0 / 32,
	15.0 / 32,
	24.0 / 32,
	9.0 / 32,
	2.0 / 32,
	19.0 / 32,
}
var SUBPIXEL_OFFSETS_SAMPLE_32_FIXED = [32]Fix{
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[0] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[1] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[2] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[3] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[4] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[5] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[6] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[7] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[8] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[9] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[10] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[11] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[12] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[13] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[14] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[15] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[16] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[17] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[18] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[19] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[20] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[21] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[22] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[23] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[24] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[25] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[26] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[27] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[28] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[29] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[30] * FIXED_FLOAT_COEF),
	Fix(SUBPIXEL_OFFSETS_SAMPLE_32[31] * FIXED_FLOAT_COEF),
}

var coverageTable = [256]uint8{
	pixelCoverage(0x00), pixelCoverage(0x01), pixelCoverage(0x02), pixelCoverage(0x03),
	pixelCoverage(0x04), pixelCoverage(0x05), pixelCoverage(0x06), pixelCoverage(0x07),
	pixelCoverage(0x08), pixelCoverage(0x09), pixelCoverage(0x0a), pixelCoverage(0x0b),
	pixelCoverage(0x0c), pixelCoverage(0x0d), pixelCoverage(0x0e), pixelCoverage(0x0f),
	pixelCoverage(0x10), pixelCoverage(0x11), pixelCoverage(0x12), pixelCoverage(0x13),
	pixelCoverage(0x14), pixelCoverage(0x15), pixelCoverage(0x16), pixelCoverage(0x17),
	pixelCoverage(0x18), pixelCoverage(0x19), pixelCoverage(0x1a), pixelCoverage(0x1b),
	pixelCoverage(0x1c), pixelCoverage(0x1d), pixelCoverage(0x1e), pixelCoverage(0x1f),
	pixelCoverage(0x20), pixelCoverage(0x21), pixelCoverage(0x22), pixelCoverage(0x23),
	pixelCoverage(0x24), pixelCoverage(0x25), pixelCoverage(0x26), pixelCoverage(0x27),
	pixelCoverage(0x28), pixelCoverage(0x29), pixelCoverage(0x2a), pixelCoverage(0x2b),
	pixelCoverage(0x2c), pixelCoverage(0x2d), pixelCoverage(0x2e), pixelCoverage(0x2f),
	pixelCoverage(0x30), pixelCoverage(0x31), pixelCoverage(0x32), pixelCoverage(0x33),
	pixelCoverage(0x34), pixelCoverage(0x35), pixelCoverage(0x36), pixelCoverage(0x37),
	pixelCoverage(0x38), pixelCoverage(0x39), pixelCoverage(0x3a), pixelCoverage(0x3b),
	pixelCoverage(0x3c), pixelCoverage(0x3d), pixelCoverage(0x3e), pixelCoverage(0x3f),
	pixelCoverage(0x40), pixelCoverage(0x41), pixelCoverage(0x42), pixelCoverage(0x43),
	pixelCoverage(0x44), pixelCoverage(0x45), pixelCoverage(0x46), pixelCoverage(0x47),
	pixelCoverage(0x48), pixelCoverage(0x49), pixelCoverage(0x4a), pixelCoverage(0x4b),
	pixelCoverage(0x4c), pixelCoverage(0x4d), pixelCoverage(0x4e), pixelCoverage(0x4f),
	pixelCoverage(0x50), pixelCoverage(0x51), pixelCoverage(0x52), pixelCoverage(0x53),
	pixelCoverage(0x54), pixelCoverage(0x55), pixelCoverage(0x56), pixelCoverage(0x57),
	pixelCoverage(0x58), pixelCoverage(0x59), pixelCoverage(0x5a), pixelCoverage(0x5b),
	pixelCoverage(0x5c), pixelCoverage(0x5d), pixelCoverage(0x5e), pixelCoverage(0x5f),
	pixelCoverage(0x60), pixelCoverage(0x61), pixelCoverage(0x62), pixelCoverage(0x63),
	pixelCoverage(0x64), pixelCoverage(0x65), pixelCoverage(0x66), pixelCoverage(0x67),
	pixelCoverage(0x68), pixelCoverage(0x69), pixelCoverage(0x6a), pixelCoverage(0x6b),
	pixelCoverage(0x6c), pixelCoverage(0x6d), pixelCoverage(0x6e), pixelCoverage(0x6f),
	pixelCoverage(0x70), pixelCoverage(0x71), pixelCoverage(0x72), pixelCoverage(0x73),
	pixelCoverage(0x74), pixelCoverage(0x75), pixelCoverage(0x76), pixelCoverage(0x77),
	pixelCoverage(0x78), pixelCoverage(0x79), pixelCoverage(0x7a), pixelCoverage(0x7b),
	pixelCoverage(0x7c), pixelCoverage(0x7d), pixelCoverage(0x7e), pixelCoverage(0x7f),
	pixelCoverage(0x80), pixelCoverage(0x81), pixelCoverage(0x82), pixelCoverage(0x83),
	pixelCoverage(0x84), pixelCoverage(0x85), pixelCoverage(0x86), pixelCoverage(0x87),
	pixelCoverage(0x88), pixelCoverage(0x89), pixelCoverage(0x8a), pixelCoverage(0x8b),
	pixelCoverage(0x8c), pixelCoverage(0x8d), pixelCoverage(0x8e), pixelCoverage(0x8f),
	pixelCoverage(0x90), pixelCoverage(0x91), pixelCoverage(0x92), pixelCoverage(0x93),
	pixelCoverage(0x94), pixelCoverage(0x95), pixelCoverage(0x96), pixelCoverage(0x97),
	pixelCoverage(0x98), pixelCoverage(0x99), pixelCoverage(0x9a), pixelCoverage(0x9b),
	pixelCoverage(0x9c), pixelCoverage(0x9d), pixelCoverage(0x9e), pixelCoverage(0x9f),
	pixelCoverage(0xa0), pixelCoverage(0xa1), pixelCoverage(0xa2), pixelCoverage(0xa3),
	pixelCoverage(0xa4), pixelCoverage(0xa5), pixelCoverage(0xa6), pixelCoverage(0xa7),
	pixelCoverage(0xa8), pixelCoverage(0xa9), pixelCoverage(0xaa), pixelCoverage(0xab),
	pixelCoverage(0xac), pixelCoverage(0xad), pixelCoverage(0xae), pixelCoverage(0xaf),
	pixelCoverage(0xb0), pixelCoverage(0xb1), pixelCoverage(0xb2), pixelCoverage(0xb3),
	pixelCoverage(0xb4), pixelCoverage(0xb5), pixelCoverage(0xb6), pixelCoverage(0xb7),
	pixelCoverage(0xb8), pixelCoverage(0xb9), pixelCoverage(0xba), pixelCoverage(0xbb),
	pixelCoverage(0xbc), pixelCoverage(0xbd), pixelCoverage(0xbe), pixelCoverage(0xbf),
	pixelCoverage(0xc0), pixelCoverage(0xc1), pixelCoverage(0xc2), pixelCoverage(0xc3),
	pixelCoverage(0xc4), pixelCoverage(0xc5), pixelCoverage(0xc6), pixelCoverage(0xc7),
	pixelCoverage(0xc8), pixelCoverage(0xc9), pixelCoverage(0xca), pixelCoverage(0xcb),
	pixelCoverage(0xcc), pixelCoverage(0xcd), pixelCoverage(0xce), pixelCoverage(0xcf),
	pixelCoverage(0xd0), pixelCoverage(0xd1), pixelCoverage(0xd2), pixelCoverage(0xd3),
	pixelCoverage(0xd4), pixelCoverage(0xd5), pixelCoverage(0xd6), pixelCoverage(0xd7),
	pixelCoverage(0xd8), pixelCoverage(0xd9), pixelCoverage(0xda), pixelCoverage(0xdb),
	pixelCoverage(0xdc), pixelCoverage(0xdd), pixelCoverage(0xde), pixelCoverage(0xdf),
	pixelCoverage(0xe0), pixelCoverage(0xe1), pixelCoverage(0xe2), pixelCoverage(0xe3),
	pixelCoverage(0xe4), pixelCoverage(0xe5), pixelCoverage(0xe6), pixelCoverage(0xe7),
	pixelCoverage(0xe8), pixelCoverage(0xe9), pixelCoverage(0xea), pixelCoverage(0xeb),
	pixelCoverage(0xec), pixelCoverage(0xed), pixelCoverage(0xee), pixelCoverage(0xef),
	pixelCoverage(0xf0), pixelCoverage(0xf1), pixelCoverage(0xf2), pixelCoverage(0xf3),
	pixelCoverage(0xf4), pixelCoverage(0xf5), pixelCoverage(0xf6), pixelCoverage(0xf7),
	pixelCoverage(0xf8), pixelCoverage(0xf9), pixelCoverage(0xfa), pixelCoverage(0xfb),
	pixelCoverage(0xfc), pixelCoverage(0xfd), pixelCoverage(0xfe), pixelCoverage(0xff),
}

func pixelCoverage(a uint8) uint8 {
	return a&1 + a>>1&1 + a>>2&1 + a>>3&1 + a>>4&1 + a>>5&1 + a>>6&1 + a>>7&1
}
