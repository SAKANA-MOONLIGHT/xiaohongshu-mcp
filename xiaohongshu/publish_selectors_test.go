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
