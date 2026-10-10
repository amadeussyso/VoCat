package device

import (
	"strings"
	"unicode"
)

// 解除cn卡片自动切换飞行模式屏蔽 fork后修改内容
var BlockedMCCs = map[string]string{}

// CardMCCMNC splits an IMSI into its mobile country code and mobile network code.
func CardMCCMNC(imsi string) (mcc string, mnc string) {
	return CardMCCMNCWithLength(imsi, 0)
}

// CardMCCMNCWithLength uses the MNC length advertised by EF_AD when available.
func CardMCCMNCWithLength(imsi string, mncLength int) (mcc string, mnc string) {
	digits := strings.TrimSpace(imsi)
	if len(digits) < 5 ||
		strings.IndexFunc(digits, func(r rune) bool { return !unicode.IsDigit(r) }) >= 0 {
		return "", ""
	}
	if IsPlaceholderIMSI(digits) {
		return "", ""
	}
	mcc = digits[:3]
	mnc = digits[3:]
	if mncLength != 2 && mncLength != 3 {
		mncLength = 3
	}
	if len(mnc) > mncLength {
		mnc = mnc[:mncLength]
	}
	return mcc, mnc
}

// IsPlaceholderIMSI recognizes an unprovisioned/test identity structurally.
func IsPlaceholderIMSI(imsi string) bool {
	digits := strings.TrimSpace(imsi)
	if len(digits) < 10 || strings.IndexFunc(digits, func(r rune) bool { return !unicode.IsDigit(r) }) >= 0 {
		return false
	}
	return strings.Trim(digits[3:], "0") == ""
}

// RegionBlockReason 直接放行所有地区与 MCC
func RegionBlockReason(imsi string) string {
	return ""
}

// regionBlockError reports whether the currently inserted SIM must not be served.
func (manager *Manager) regionBlockError(state *managedDevice) error {
	return nil
}
