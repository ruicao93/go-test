package hns

import (
	"encoding/json"
	"testing"

	"github.com/Microsoft/hcsshim"
	//"github.com/Microsoft/hcsshim/internal/hns"
)

type ExtendedHNSEndpoint struct {
	hcsshim.HNSEndpoint
	AdditionalParams map[string]string `json:"AdditionalParams,omitempty"`
}

func TestSharmalExtendedHNSEndoint(t *testing.T) {
	ep := ExtendedHNSEndpoint{
		HNSEndpoint: hcsshim.HNSEndpoint{
			Name: "test",
		},
		AdditionalParams: map[string]string{"1": "1"},
	}
	res, _ := json.Marshal(ep)
	t.Logf("test: %s", res)
}
