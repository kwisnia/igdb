// Code generated by "stringer -type=AgeRatingCategory,AgeRatingEnum"; DO NOT EDIT.

package igdb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AgeRatingESRB-1]
	_ = x[AgeRatingPEGI-2]
	_ = x[AgeRatingCERO-3]
	_ = x[AgeRatingUSK-4]
	_ = x[AgeRatingGRAC-5]
	_ = x[AgeRatingCLASSIND-6]
	_ = x[AgeRatingACB-7]
}

const _AgeRatingCategory_name = "AgeRatingESRBAgeRatingPEGIAgeRatingCEROAgeRatingUSKAgeRatingGRACAgeRatingCLASSINDAgeRatingACB"

var _AgeRatingCategory_index = [...]uint8{0, 13, 26, 39, 51, 64, 81, 93}

func (i AgeRatingCategory) String() string {
	i -= 1
	if i < 0 || i >= AgeRatingCategory(len(_AgeRatingCategory_index)-1) {
		return "AgeRatingCategory(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AgeRatingCategory_name[_AgeRatingCategory_index[i]:_AgeRatingCategory_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[AgeRatingThree-1]
	_ = x[AgeRatingSeven-2]
	_ = x[AgeRatingTwelve-3]
	_ = x[AgeRatingSixteen-4]
	_ = x[AgeRatingEighteen-5]
	_ = x[AgeRatingRP-6]
	_ = x[AgeRatingEC-7]
	_ = x[AgeRatingE-8]
	_ = x[AgeRatingE10-9]
	_ = x[AgeRatingT-10]
	_ = x[AgeRatingM-11]
	_ = x[AgeRatingAO-12]
	_ = x[AgeRatingCERO_A-13]
	_ = x[AgeRatingCERO_B-14]
	_ = x[AgeRatingCERO_C-15]
	_ = x[AgeRatingCERO_D-16]
	_ = x[AgeRatingCERO_Z-17]
	_ = x[AgeRatingUSK_0-18]
	_ = x[AgeRatingUSK_6-19]
	_ = x[AgeRatingUSK_12-20]
	_ = x[AgeRatingUSK_18-21]
	_ = x[AgeRatingGRAC_ALL-22]
	_ = x[AgeRatingGRAC_Twelve-23]
	_ = x[AgeRatingGRAC_Fifteen-24]
	_ = x[AgeRatingGRAC_Eighteen-25]
	_ = x[AgeRatingGRAC_TESTING-26]
	_ = x[AgeRatingCLASS_IND_L-27]
	_ = x[AgeRatingCLASS_IND_Ten-28]
	_ = x[AgeRatingCLASS_IND_Twelve-29]
	_ = x[AgeRatingCLASS_IND_Fourteen-30]
	_ = x[AgeRatingCLASS_IND_Sixteen-31]
	_ = x[AgeRatingCLASS_IND_Eighteen-32]
	_ = x[AgeRatingACB_G-33]
	_ = x[AgeRatingACB_PG-34]
	_ = x[AgeRatingACB_M-35]
	_ = x[AgeRatingACB_MA15-36]
	_ = x[AgeRatingACB_R18-37]
	_ = x[AgeRatingACB_RC-38]
}

const _AgeRatingEnum_name = "AgeRatingThreeAgeRatingSevenAgeRatingTwelveAgeRatingSixteenAgeRatingEighteenAgeRatingRPAgeRatingECAgeRatingEAgeRatingE10AgeRatingTAgeRatingMAgeRatingAOAgeRatingCERO_AAgeRatingCERO_BAgeRatingCERO_CAgeRatingCERO_DAgeRatingCERO_ZAgeRatingUSK_0AgeRatingUSK_6AgeRatingUSK_12AgeRatingUSK_18AgeRatingGRAC_ALLAgeRatingGRAC_TwelveAgeRatingGRAC_FifteenAgeRatingGRAC_EighteenAgeRatingGRAC_TESTINGAgeRatingCLASS_IND_LAgeRatingCLASS_IND_TenAgeRatingCLASS_IND_TwelveAgeRatingCLASS_IND_FourteenAgeRatingCLASS_IND_SixteenAgeRatingCLASS_IND_EighteenAgeRatingACB_GAgeRatingACB_PGAgeRatingACB_MAgeRatingACB_MA15AgeRatingACB_R18AgeRatingACB_RC"

var _AgeRatingEnum_index = [...]uint16{0, 14, 28, 43, 59, 76, 87, 98, 108, 120, 130, 140, 151, 166, 181, 196, 211, 226, 240, 254, 269, 284, 301, 321, 342, 364, 385, 405, 427, 452, 479, 505, 532, 546, 561, 575, 592, 608, 623}

func (i AgeRatingEnum) String() string {
	i -= 1
	if i < 0 || i >= AgeRatingEnum(len(_AgeRatingEnum_index)-1) {
		return "AgeRatingEnum(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AgeRatingEnum_name[_AgeRatingEnum_index[i]:_AgeRatingEnum_index[i+1]]
}
