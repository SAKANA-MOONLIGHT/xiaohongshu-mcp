package xiaohongshu

import "testing"

// TestSelectorsNotEmpty 确保所有选择器字段不为空
func TestSelectorsNotEmpty(t *testing.T) {
	fields := map[string]string{
		"UploadContentArea":   selectors.UploadContentArea,
		"PublishTab":          selectors.PublishTab,
		"PopoverCover":        selectors.PopoverCover,
		"UploadInput":         selectors.UploadInput,
		"UploadInputFallback": selectors.UploadInputFallback,
		"ImagePreview":        selectors.ImagePreview,
		"TitleInput":          selectors.TitleInput,
		"ContentEditorQuill":  selectors.ContentEditorQuill,
		"ContentPlaceholder":  selectors.ContentPlaceholder,
		"TitleMaxLength":      selectors.TitleMaxLength,
		"ContentMaxLength":    selectors.ContentMaxLength,
		"PublishButton":       selectors.PublishButton,
		"TopicContainer":      selectors.TopicContainer,
		"TopicItem":           selectors.TopicItem,
		"VisibilityDropdown":  selectors.VisibilityDropdown,
		"VisibilityOptions":   selectors.VisibilityOptions,
		"ScheduleSwitch":      selectors.ScheduleSwitch,
		"ScheduleDateInput":   selectors.ScheduleDateInput,
		"OriginalSwitchCards": selectors.OriginalSwitchCards,
		"OriginalSwitch":      selectors.OriginalSwitch,
		"AddProductSpans":     selectors.AddProductSpans,
		"ProductModal":        selectors.ProductModal,
		"GoodsSearchInput":    selectors.GoodsSearchInput,
		"GoodsLoading":        selectors.GoodsLoading,
		"GoodsListItem":       selectors.GoodsListItem,
		"GoodsCheckbox":       selectors.GoodsCheckbox,
		"GoodsSaveButton":     selectors.GoodsSaveButton,
		"GoodsSaveFallback":   selectors.GoodsSaveFallback,
	}

	for name, value := range fields {
		if value == "" {
			t.Errorf("selectors.%s 不能为空", name)
		}
	}
}

// TestCriticalSelectorsNotEmpty 确保所有 critical 选择器字段不为空
func TestCriticalSelectorsNotEmpty(t *testing.T) {
	critical := CriticalSelectors()
	if len(critical) == 0 {
		t.Fatal("CriticalSelectors() 返回空 map")
	}
	expectedKeys := []string{
		"UploadContentArea",
		"UploadInput",
		"TitleInput",
		"ContentEditorQuill",
		"PublishButton",
	}
	for _, key := range expectedKeys {
		val, ok := critical[key]
		if !ok {
			t.Errorf("CriticalSelectors() 缺少字段: %s", key)
			continue
		}
		if val == "" {
			t.Errorf("CriticalSelectors().%s 不能为空", key)
		}
	}
}

// TestAllSelectorsCount 确保 AllSelectors() 返回全部 28 个字段
func TestAllSelectorsCount(t *testing.T) {
	all := AllSelectors()
	if len(all) != 28 {
		t.Errorf("AllSelectors() 期望 28 个字段，实际 %d 个", len(all))
	}
	// 验证所有值非空
	for name, val := range all {
		if val == "" {
			t.Errorf("AllSelectors().%s 不能为空", name)
		}
	}
}
