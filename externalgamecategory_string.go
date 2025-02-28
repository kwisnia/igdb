// Code generated by "stringer -type=ExternalGameCategory"; DO NOT EDIT.

package igdb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ExternalSteam-1]
	_ = x[ExternalGOG-5]
	_ = x[ExternalYoutube-10]
	_ = x[ExternalMicrosoft-11]
	_ = x[ExternalApple-13]
	_ = x[ExternalTwitch-14]
	_ = x[ExternalAndroid-15]
	_ = x[ExternalAmazonAsin-20]
	_ = x[ExternalAmazonLuna-22]
	_ = x[ExternalAmazonAdg-23]
	_ = x[ExternalEpicGameStore-26]
	_ = x[ExternalOculus-28]
	_ = x[ExternalUtomik-29]
	_ = x[ExternalItchIo-30]
	_ = x[ExternalXboxMarketplace-31]
	_ = x[ExternalKartridge-32]
	_ = x[ExternalPlaystationStoreUS-36]
	_ = x[ExternalFocusEntertainment-37]
}

const (
	_ExternalGameCategory_name_0 = "ExternalSteam"
	_ExternalGameCategory_name_1 = "ExternalGOG"
	_ExternalGameCategory_name_2 = "ExternalYoutubeExternalMicrosoft"
	_ExternalGameCategory_name_3 = "ExternalAppleExternalTwitchExternalAndroid"
	_ExternalGameCategory_name_4 = "ExternalAmazonAsin"
	_ExternalGameCategory_name_5 = "ExternalAmazonLunaExternalAmazonAdg"
	_ExternalGameCategory_name_6 = "ExternalEpicGameStore"
	_ExternalGameCategory_name_7 = "ExternalOculusExternalUtomikExternalItchIoExternalXboxMarketplaceExternalKartridge"
	_ExternalGameCategory_name_8 = "ExternalPlaystationStoreUSExternalFocusEntertainment"
)

var (
	_ExternalGameCategory_index_2 = [...]uint8{0, 15, 32}
	_ExternalGameCategory_index_3 = [...]uint8{0, 13, 27, 42}
	_ExternalGameCategory_index_5 = [...]uint8{0, 18, 35}
	_ExternalGameCategory_index_7 = [...]uint8{0, 14, 28, 42, 65, 82}
	_ExternalGameCategory_index_8 = [...]uint8{0, 26, 52}
)

func (i ExternalGameCategory) String() string {
	switch {
	case i == 1:
		return _ExternalGameCategory_name_0
	case i == 5:
		return _ExternalGameCategory_name_1
	case 10 <= i && i <= 11:
		i -= 10
		return _ExternalGameCategory_name_2[_ExternalGameCategory_index_2[i]:_ExternalGameCategory_index_2[i+1]]
	case 13 <= i && i <= 15:
		i -= 13
		return _ExternalGameCategory_name_3[_ExternalGameCategory_index_3[i]:_ExternalGameCategory_index_3[i+1]]
	case i == 20:
		return _ExternalGameCategory_name_4
	case 22 <= i && i <= 23:
		i -= 22
		return _ExternalGameCategory_name_5[_ExternalGameCategory_index_5[i]:_ExternalGameCategory_index_5[i+1]]
	case i == 26:
		return _ExternalGameCategory_name_6
	case 28 <= i && i <= 32:
		i -= 28
		return _ExternalGameCategory_name_7[_ExternalGameCategory_index_7[i]:_ExternalGameCategory_index_7[i+1]]
	case 36 <= i && i <= 37:
		i -= 36
		return _ExternalGameCategory_name_8[_ExternalGameCategory_index_8[i]:_ExternalGameCategory_index_8[i+1]]
	default:
		return "ExternalGameCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
