package promtail

import "maps"

func copyLabels(src map[string]string) map[string]string {
	dst := make(map[string]string, len(src))
	maps.Copy(dst, src)
	return dst
}

func copyAndMergeLabels(srcs ...map[string]string) map[string]string {
	//
	// We do assume, that gathering map's sizes in a single loop is faster
	// then additional allocations in loop fot the target map
	// 	TODO: verify :)
	//
	predictedLength := 0
	for i := range srcs {
		predictedLength += len(srcs[i])
	}
	dst := make(map[string]string, predictedLength)

	for i := range srcs {
		for key := range srcs[i] {
			dst[key] = srcs[i][key]
		}
	}

	return dst
}
