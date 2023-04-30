package designerpdfviewer

func designerPdfViewer(h []int32, word string) int32 {
	var maxH int32
	for _, c := range word {
		if height := h[c-97]; height > maxH {
			maxH = height
		}
	}
	return maxH * int32(len(word))
}
