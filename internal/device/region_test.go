package device

import "testing"

func TestRegionBlockReason(t *testing.T) {
	// 已解除大陆 460/461 限制，断言不再触发屏蔽
	if reason := RegionBlockReason("460001234567890"); reason != "" {
		t.Errorf("RegionBlockReason should not block, got: %s", reason)
	}
	if reason := RegionBlockReason("461001234567890"); reason != "" {
		t.Errorf("RegionBlockReason should not block, got: %s", reason)
	}
}
