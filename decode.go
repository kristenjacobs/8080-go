package main

import (
	"fmt"
)

func step(ms *machineState) {
	opcode := fetch(ms)
	decodeAndExecute(ms, opcode)
}

func fetch(ms *machineState) uint8 {
	return ms.readMem(ms.pc, 1)[0]
}

func decodeAndExecute(ms *machineState, opcode uint8) {
	Debug.Printf("decodeAndExecute, opcode: 0x%02x\n", opcode)
	switch opcode {
	case 0x00:
		instr_0x00_NOP(ms)
	case 0x01:
		instr_0x01_LXI_B_D16(ms)
	case 0x02:
		instr_0x02_STAX(ms)
	case 0x03:
		instr_0x03_INX_B(ms)
	case 0x04:
		instr_0x04_INR_B(ms)
	case 0x05:
		instr_0x05_DCR_B(ms)
	case 0x06:
		instr_0x06_MVI_B_D8(ms)
	case 0x07:
		instr_0x07_RLC(ms)
	case 0x08:
		invalid(0x08)
	case 0x09:
		instr_0x09_DAD_B(ms)
	case 0x0a:
		instr_0x0a_LDAX_B(ms)
	case 0x0b:
		instr_0x0b_DCX(ms)
	case 0x0c:
		instr_0x0c_INR_C(ms)
	case 0x0d:
		instr_0x0d_DCR_C(ms)
	case 0x0e:
		instr_0x0e_MVI_C_D8(ms)
	case 0x0f:
		instr_0x0f_RRC(ms)
	case 0x10:
		invalid(0x10)
	case 0x11:
		instr_0x11_LXI_D_D16(ms)
	case 0x12:
		instr_0x12_STAX(ms)
	case 0x13:
		instr_0x13_INX_D(ms)
	case 0x14:
		instr_0x14_INR_D(ms)
	case 0x15:
		instr_0x15_DCR_D(ms)
	case 0x16:
		instr_0x16_MVI_D_D8(ms)
	case 0x17:
		instr_0x17_RAL(ms)
	case 0x18:
		invalid(0x18)
	case 0x19:
		instr_0x19_DAD_D(ms)
	case 0x1a:
		instr_0x1a_LDAX_D(ms)
	case 0x1b:
		instr_0x1b_DCX(ms)
	case 0x1c:
		instr_0x1c_INR_E(ms)
	case 0x1d:
		instr_0x1d_DCR_E(ms)
	case 0x1e:
		instr_0x1e_MVI_E_D8(ms)
	case 0x1f:
		instr_0x1f_RAR(ms)
	case 0x20:
		instr_0x20_RIM(ms)
	case 0x21:
		instr_0x21_LXI_H_D16(ms)
	case 0x22:
		instr_0x22_SHLD(ms)
	case 0x23:
		instr_0x23_INX_H(ms)
	case 0x24:
		instr_0x24_INR_H(ms)
	case 0x25:
		instr_0x25_DCR_H(ms)
	case 0x26:
		instr_0x26_MVI_H_D8(ms)
	case 0x27:
		instr_0x27_DAA(ms)
	case 0x28:
		invalid(0x28)
	case 0x29:
		instr_0x29_DAD_H(ms)
	case 0x2a:
		instr_0x2a_LHLD(ms)
	case 0x2b:
		instr_0x2b_DCX(ms)
	case 0x2c:
		instr_0x2c_INR_L(ms)
	case 0x2d:
		instr_0x2d_DCR_L(ms)
	case 0x2e:
		instr_0x2e_MVI_L_D8(ms)
	case 0x2f:
		instr_0x2f_CMA(ms)
	case 0x30:
		instr_0x30_SIM(ms)
	case 0x31:
		instr_0x31_LXI_SP_D16(ms)
	case 0x32:
		instr_0x32_STA_adr(ms)
	case 0x33:
		instr_0x33_INX_SP(ms)
	case 0x34:
		instr_0x34_INR(ms)
	case 0x35:
		instr_0x35_DCR_M(ms)
	case 0x36:
		instr_0x36_MVI_M_D8(ms)
	case 0x37:
		instr_0x37_STC(ms)
	case 0x38:
		invalid(0x30)
	case 0x39:
		instr_0x39_DAD(ms)
	case 0x3a:
		instr_0x3a_LDA(ms)
	case 0x3b:
		instr_0x3b_DCX(ms)
	case 0x3c:
		instr_0x3c_INR_A(ms)
	case 0x3d:
		instr_0x3d_DCR_A(ms)
	case 0x3e:
		instr_0x3e_MVI_A_D8(ms)
	case 0x3f:
		instr_0x3f_CMC(ms)
	case 0x40:
		instr_0x40_MOV_B_B(ms)
	case 0x41:
		instr_0x41_MOV_B_C(ms)
	case 0x42:
		instr_0x42_MOV_B_D(ms)
	case 0x43:
		instr_0x43_MOV_B_E(ms)
	case 0x44:
		instr_0x44_MOV_B_H(ms)
	case 0x45:
		instr_0x45_MOV_B_L(ms)
	case 0x46:
		instr_0x46_MOV_B_M(ms)
	case 0x47:
		instr_0x47_MOV_B_A(ms)
	case 0x48:
		instr_0x48_MOV_C_B(ms)
	case 0x49:
		instr_0x49_MOV_C_C(ms)
	case 0x4a:
		instr_0x4a_MOV_C_D(ms)
	case 0x4b:
		instr_0x4b_MOV_C_E(ms)
	case 0x4c:
		instr_0x4c_MOV_C_H(ms)
	case 0x4d:
		instr_0x4d_MOV_C_L(ms)
	case 0x4e:
		instr_0x4e_MOV_C_M(ms)
	case 0x4f:
		instr_0x4f_MOV_C_A(ms)
	case 0x50:
		instr_0x50_MOV_D_B(ms)
	case 0x51:
		instr_0x51_MOV_D_C(ms)
	case 0x52:
		instr_0x52_MOV_D_D(ms)
	case 0x53:
		instr_0x53_MOV_D_E(ms)
	case 0x54:
		instr_0x54_MOV_D_H(ms)
	case 0x55:
		instr_0x55_MOV_D_L(ms)
	case 0x56:
		instr_0x56_MOV_D_M(ms)
	case 0x57:
		instr_0x57_MOV_D_A(ms)
	case 0x58:
		instr_0x58_MOV_E_B(ms)
	case 0x59:
		instr_0x59_MOV_E_C(ms)
	case 0x5a:
		instr_0x5a_MOV_E_D(ms)
	case 0x5b:
		instr_0x5b_MOV_E_E(ms)
	case 0x5c:
		instr_0x5c_MOV_E_H(ms)
	case 0x5d:
		instr_0x5d_MOV_E_L(ms)
	case 0x5e:
		instr_0x5e_MOV_E_M(ms)
	case 0x5f:
		instr_0x5f_MOV_E_A(ms)
	case 0x60:
		instr_0x60_MOV_H_B(ms)
	case 0x61:
		instr_0x61_MOV_H_C(ms)
	case 0x62:
		instr_0x62_MOV_H_D(ms)
	case 0x63:
		instr_0x63_MOV_H_E(ms)
	case 0x64:
		instr_0x64_MOV_H_H(ms)
	case 0x65:
		instr_0x65_MOV_H_L(ms)
	case 0x66:
		instr_0x66_MOV_H_M(ms)
	case 0x67:
		instr_0x67_MOV_H_A(ms)
	case 0x68:
		instr_0x68_MOV_L_B(ms)
	case 0x69:
		instr_0x69_MOV_L_C(ms)
	case 0x6a:
		instr_0x6a_MOV_L_D(ms)
	case 0x6b:
		instr_0x6b_MOV_L_E(ms)
	case 0x6c:
		instr_0x6c_MOV_L_H(ms)
	case 0x6d:
		instr_0x6d_MOV_L_L(ms)
	case 0x6e:
		instr_0x6e_MOV_L_M(ms)
	case 0x6f:
		instr_0x6f_MOV_L_A(ms)
	case 0x70:
		instr_0x70_MOV_M_B(ms)
	case 0x71:
		instr_0x71_MOV_M_C(ms)
	case 0x72:
		instr_0x72_MOV_M_D(ms)
	case 0x73:
		instr_0x73_MOV_M_E(ms)
	case 0x74:
		instr_0x74_MOV_M_H(ms)
	case 0x75:
		instr_0x75_MOV_M_L(ms)
	case 0x76:
		instr_0x76_HLT(ms)
	case 0x77:
		instr_0x77_MOV_M_A(ms)
	case 0x78:
		instr_0x78_MOV_A_B(ms)
	case 0x79:
		instr_0x79_MOV_A_C(ms)
	case 0x7a:
		instr_0x7a_MOV_A_D(ms)
	case 0x7b:
		instr_0x7b_MOV_A_E(ms)
	case 0x7c:
		instr_0x7c_MOV_A_H(ms)
	case 0x7d:
		instr_0x7d_MOV_A_L(ms)
	case 0x7e:
		instr_0x7e_MOV_A_M(ms)
	case 0x7f:
		instr_0x7f_MOV_A_A(ms)
	case 0x80:
		instr_0x80_ADD_B(ms)
	case 0x81:
		instr_0x81_ADD_C(ms)
	case 0x82:
		instr_0x82_ADD_D(ms)
	case 0x83:
		instr_0x83_ADD_E(ms)
	case 0x84:
		instr_0x84_ADD_H(ms)
	case 0x85:
		instr_0x85_ADD_L(ms)
	case 0x86:
		instr_0x86_ADD(ms)
	case 0x87:
		instr_0x87_ADD_A(ms)
	case 0x88:
		instr_0x88_ADC_B(ms)
	case 0x89:
		instr_0x89_ADC_C(ms)
	case 0x8a:
		instr_0x8a_ADC_D(ms)
	case 0x8b:
		instr_0x8b_ADC_E(ms)
	case 0x8c:
		instr_0x8c_ADC_H(ms)
	case 0x8d:
		instr_0x8d_ADC_L(ms)
	case 0x8e:
		instr_0x8e_ADC(ms)
	case 0x8f:
		instr_0x8f_ADC_A(ms)
	case 0x90:
		instr_0x90_SUB_B(ms)
	case 0x91:
		instr_0x91_SUB_C(ms)
	case 0x92:
		instr_0x92_SUB_D(ms)
	case 0x93:
		instr_0x93_SUB_E(ms)
	case 0x94:
		instr_0x94_SUB_H(ms)
	case 0x95:
		instr_0x95_SUB_L(ms)
	case 0x96:
		instr_0x96_SUB(ms)
	case 0x97:
		instr_0x97_SUB_A(ms)
	case 0x98:
		instr_0x98_SBB_B(ms)
	case 0x99:
		instr_0x99_SBB_C(ms)
	case 0x9a:
		instr_0x9a_SBB_D(ms)
	case 0x9b:
		instr_0x9b_SBB_E(ms)
	case 0x9c:
		instr_0x9c_SBB_H(ms)
	case 0x9d:
		instr_0x9d_SBB_L(ms)
	case 0x9e:
		instr_0x9e_SBB(ms)
	case 0x9f:
		instr_0x9f_SBB_A(ms)
	case 0xa0:
		instr_0xa0_ANA_B(ms)
	case 0xa1:
		instr_0xa1_ANA_C(ms)
	case 0xa2:
		instr_0xa2_ANA_D(ms)
	case 0xa3:
		instr_0xa3_ANA_E(ms)
	case 0xa4:
		instr_0xa4_ANA_H(ms)
	case 0xa5:
		instr_0xa5_ANA_L(ms)
	case 0xa6:
		instr_0xa6_ANA(ms)
	case 0xa7:
		instr_0xa7_ANA_A(ms)
	case 0xa8:
		instr_0xa8_XRA(ms)
	case 0xa9:
		instr_0xa9_XRA(ms)
	case 0xaa:
		instr_0xaa_XRA(ms)
	case 0xab:
		instr_0xab_XRA(ms)
	case 0xac:
		instr_0xac_XRA(ms)
	case 0xad:
		instr_0xad_XRA(ms)
	case 0xae:
		instr_0xae_XRA(ms)
	case 0xaf:
		instr_0xaf_XRA_A(ms)
	case 0xb0:
		instr_0xb0_ORA(ms)
	case 0xb1:
		instr_0xb1_ORA(ms)
	case 0xb2:
		instr_0xb2_ORA(ms)
	case 0xb3:
		instr_0xb3_ORA(ms)
	case 0xb4:
		instr_0xb4_ORA(ms)
	case 0xb5:
		instr_0xb5_ORA(ms)
	case 0xb6:
		instr_0xb6_ORA(ms)
	case 0xb7:
		instr_0xb7_ORA(ms)
	case 0xb8:
		instr_0xb8_CMP(ms)
	case 0xb9:
		instr_0xb9_CMP(ms)
	case 0xba:
		instr_0xba_CMP(ms)
	case 0xbb:
		instr_0xbb_CMP(ms)
	case 0xbc:
		instr_0xbc_CMP(ms)
	case 0xbd:
		instr_0xbd_CMP(ms)
	case 0xbe:
		instr_0xbe_CMP(ms)
	case 0xbf:
		instr_0xbf_CMP(ms)
	case 0xc0:
		instr_0xc0_RNZ(ms)
	case 0xc1:
		instr_0xc1_POP_B(ms)
	case 0xc2:
		instr_0xc2_JNZ_adr(ms)
	case 0xc3:
		instr_0xc3_JMP_adr(ms)
	case 0xc4:
		instr_0xc4_CNZ_adr(ms)
	case 0xc5:
		instr_0xc5_PUSH_B(ms)
	case 0xc6:
		instr_0xc6_ADI_D8(ms)
	case 0xc7:
		instr_0xc7_RST_0(ms)
	case 0xc8:
		instr_0xc8_RZ(ms)
	case 0xc9:
		instr_0xc9_RET(ms)
	case 0xca:
		instr_0xca_JZ_adr(ms)
	case 0xcb:
		invalid(0xcb)
	case 0xcc:
		instr_0xcc_CZ_adr(ms)
	case 0xcd:
		instr_0xcd_CALL_adr(ms)
	case 0xce:
		instr_0xce_ACI_D8(ms)
	case 0xcf:
		instr_0xcf_RST_1(ms)
	case 0xd0:
		instr_0xd0_RNC(ms)
	case 0xd1:
		instr_0xd1_POP_D(ms)
	case 0xd2:
		instr_0xd2_JNC_adr(ms)
	case 0xd3:
		instr_0xd3_OUT_D8(ms)
	case 0xd4:
		instr_0xd4_CNC_adr(ms)
	case 0xd5:
		instr_0xd5_PUSH_D(ms)
	case 0xd6:
		instr_0xd6_SUI_D8(ms)
	case 0xd7:
		instr_0xd7_RST_2(ms)
	case 0xd8:
		instr_0xd8_RC(ms)
	case 0xd9:
		invalid(0xd9)
	case 0xda:
		instr_0xda_JC_adr(ms)
	case 0xdb:
		instr_0xdb_IN(ms)
	case 0xdc:
		instr_0xdc_CC_adr(ms)
	case 0xdd:
		invalid(0xdd)
	case 0xde:
		instr_0xde_SBI_D8(ms)
	case 0xdf:
		instr_0xdf_RST_3(ms)
	case 0xe0:
		instr_0xe0_RPO(ms)
	case 0xe1:
		instr_0xe1_POP_H(ms)
	case 0xe2:
		instr_0xe2_JPO_adr(ms)
	case 0xe3:
		instr_0xe3_XTHL(ms)
	case 0xe4:
		instr_0xe4_CPO_adr(ms)
	case 0xe5:
		instr_0xe5_PUSH_H(ms)
	case 0xe6:
		instr_0xe6_ANI_D8(ms)
	case 0xe7:
		instr_0xe7_RST_4(ms)
	case 0xe8:
		instr_0xe8_RPE(ms)
	case 0xe9:
		instr_0xe9_PCHL(ms)
	case 0xea:
		instr_0xea_JPE_adr(ms)
	case 0xeb:
		instr_0xeb_XCHG(ms)
	case 0xec:
		instr_0xec_CPE_adr(ms)
	case 0xed:
		invalid(0xed)
	case 0xee:
		instr_0xee_XRI_D8(ms)
	case 0xef:
		instr_0xef_RST_5(ms)
	case 0xf0:
		instr_0xf0_RP(ms)
	case 0xf1:
		instr_0xf1_POP(ms)
	case 0xf2:
		instr_0xf2_JP_adr(ms)
	case 0xf3:
		instr_0xf3_DI(ms)
	case 0xf4:
		instr_0xf4_CP_adr(ms)
	case 0xf5:
		instr_0xf5_PUSH(ms)
	case 0xf6:
		instr_0xf6_ORI_D8(ms)
	case 0xf7:
		instr_0xf7_RST_6(ms)
	case 0xf8:
		instr_0xf8_RM(ms)
	case 0xf9:
		instr_0xf9_SPHL(ms)
	case 0xfa:
		instr_0xfa_JM_adr(ms)
	case 0xfb:
		instr_0xfb_EI(ms)
	case 0xfc:
		instr_0xfc_CM_adr(ms)
	case 0xfd:
		invalid(0xfd)
	case 0xfe:
		instr_0xfe_CPI_D8(ms)
	case 0xff:
		instr_0xff_RST_7(ms)
	}
}

func invalid(opcode uint8) {
	panic(fmt.Sprintf("Invalid opcode: 0x%02x\n", opcode))
}
