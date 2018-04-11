package common

import "time"

func GetGeneratorTactTime(generationSpeed float64) time.Duration {
	return time.Duration((float64(1)/float64(generationSpeed)) * float64(time.Second))
}