package util

import (
	"math"
)

const (
	decayFactory = 0.95
)

// CalculateNormalizationRelevanceScore 计算归一化相似度分数
func CalculateNormalizationRelevanceScore(queryEmbedding, memoryEmbedding []float32) float64 {
	// 计算余弦相似度
	vectorLength := len(queryEmbedding)
	if len(queryEmbedding) > len(memoryEmbedding) {
		vectorLength = len(memoryEmbedding)
	}
	var sum, magnitudeQuery, magnitudeMemory float64
	for i := 0; i < vectorLength; i++ {
		sum += float64(queryEmbedding[i] * memoryEmbedding[i])
		magnitudeQuery += math.Pow(float64(queryEmbedding[i]), 2)
		magnitudeMemory += math.Pow(float64(memoryEmbedding[i]), 2)
	}
	if magnitudeQuery == 0 || magnitudeMemory == 0 {
		return 0
	}
	relevance := sum / (math.Sqrt(magnitudeQuery) * math.Sqrt(magnitudeMemory))
	// 结果归一化到[0,1]范围
	return (relevance + 1) / 2
}

// CalculateNormalizationRecencyScore 计算归一化相近度分数
func CalculateNormalizationRecencyScore(lastVisitTime, currentTime int64) float64 {
	// 时间记录为秒，按小时衰减
	decayCount := (currentTime - lastVisitTime) / 3600
	recency := math.Pow(decayFactory, float64(decayCount))
	// 结果归一化到[0,1]范围
	minRecency := math.Pow(decayFactory, math.MaxInt32)
	return (recency - minRecency) / (1 - minRecency)
}

// CalculateNormalizationImportanceScore 计算归一化重要度得分
func CalculateNormalizationImportanceScore(importance int) float64 {
	if importance < 1 || importance > 10 {
		return 0
	}
	// 结果归一化到[0,1]范围
	return float64(importance / 10)
}
