// Code generated by "stringer -type=GameCategory,GameStatus"; DO NOT EDIT.

package igdb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MainGame-0]
	_ = x[DLCAddon-1]
	_ = x[Expansion-2]
	_ = x[Bundle-3]
	_ = x[StandaloneExpansion-4]
	_ = x[Mod-5]
	_ = x[Episode-6]
	_ = x[Season-7]
	_ = x[Remake-8]
	_ = x[Remaster-9]
	_ = x[ExpandedGame-10]
	_ = x[Port-11]
	_ = x[Fork-12]
}

const _GameCategory_name = "MainGameDLCAddonExpansionBundleStandaloneExpansionModEpisodeSeasonRemakeRemasterExpandedGamePortFork"

var _GameCategory_index = [...]uint8{0, 8, 16, 25, 31, 50, 53, 60, 66, 72, 80, 92, 96, 100}

func (i GameCategory) String() string {
	if i < 0 || i >= GameCategory(len(_GameCategory_index)-1) {
		return "GameCategory(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GameCategory_name[_GameCategory_index[i]:_GameCategory_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StatusReleased-0]
	_ = x[StatusAlpha-2]
	_ = x[StatusBeta-3]
	_ = x[StatusEarlyAccess-4]
	_ = x[StatusOffline-5]
	_ = x[StatusCancelled-6]
	_ = x[StatusRumored-7]
	_ = x[StatusDelisted-8]
}

const (
	_GameStatus_name_0 = "StatusReleased"
	_GameStatus_name_1 = "StatusAlphaStatusBetaStatusEarlyAccessStatusOfflineStatusCancelledStatusRumoredStatusDelisted"
)

var (
	_GameStatus_index_1 = [...]uint8{0, 11, 21, 38, 51, 66, 79, 93}
)

func (i GameStatus) String() string {
	switch {
	case i == 0:
		return _GameStatus_name_0
	case 2 <= i && i <= 8:
		i -= 2
		return _GameStatus_name_1[_GameStatus_index_1[i]:_GameStatus_index_1[i+1]]
	default:
		return "GameStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
