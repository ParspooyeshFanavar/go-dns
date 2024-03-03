package iohandlers

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

func init() {
	Marshalers["json"] = toJson
}

func toJson(d *DnsSchema) {
	jsonData, err := json.Marshal(d)
	if err != nil {
		zap.L().Sugar().Warnf("Error converting to JSON: %v", err)
	}
	fmt.Printf("%s\n", jsonData)
}
