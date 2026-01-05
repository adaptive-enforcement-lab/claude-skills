package domain

import "testing"

func TestPluginConfig_GetMarketplaceName(t *testing.T) {
	tests := []struct {
		name      string
		config    *PluginConfig
		pluginKey string
		want      string
	}{
		{
			name: "uses MarketplaceName override when set",
			config: &PluginConfig{
				MarketplaceName: "enforcement",
			},
			pluginKey: "enforce",
			want:      "enforcement",
		},
		{
			name: "falls back to pluginKey when MarketplaceName empty",
			config: &PluginConfig{
				MarketplaceName: "",
			},
			pluginKey: "patterns",
			want:      "patterns",
		},
		{
			name:      "falls back to pluginKey when MarketplaceName not set",
			config:    &PluginConfig{},
			pluginKey: "secure",
			want:      "secure",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.config.GetMarketplaceName(tt.pluginKey)
			if got != tt.want {
				t.Errorf("GetMarketplaceName() = %v, want %v", got, tt.want)
			}
		})
	}
}
