package aggregator

import "testing"

func TestNew(t *testing.T) {
	aggr := New("xendit", &Config{
		ApiKey: "",
	})
	aggr.Url = ""
	aggr.RegisterCallback()
	// aggr.GetCurrentUrl()
}
