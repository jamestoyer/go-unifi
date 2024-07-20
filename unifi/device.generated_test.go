package unifi

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDevicePortOverrides_UnmarshalJSON(t *testing.T) {
	sampleOverride := DevicePortOverrides{
		Autoneg:                nil,
		PortKeepaliveEnabled:   boolPtr(false),
		Name:                   stringPtr(""),
		SettingPreference:      nil,
		ExcludedNetworkIDs:     nil,
		PortSecurityMACAddress: &[]string{},
	}

	out, err := json.Marshal(sampleOverride)
	require.NoError(t, err)
	assert.Equal(t, `{"name":"","port_keepalive_enabled":false,"port_security_mac_address":[]}`, string(out))

	var output DevicePortOverrides
	err = json.Unmarshal([]byte(`{"name":"","port_keepalive_enabled":false,"port_security_mac_address":[]}`), &output)
	require.NoError(t, err)
	// assert.Equal(t, sampleOverride, output)

	var te *[]string
	var ns *string
	assert.Nil(t, output.Autoneg)
	assert.Equal(t, stringPtr(""), output.Name)
	assert.Equal(t, boolPtr(false), output.PortKeepaliveEnabled)
	assert.Equal(t, &[]string{}, output.PortSecurityMACAddress)
	assert.Equal(t, te, output.ExcludedNetworkIDs)
	assert.Equal(t, ns, output.PortProfileID)

	// type fields struct {
	// 	AggregateNumPorts            int
	// 	Autoneg                      bool
	// 	Dot1XCtrl                    string
	// 	Dot1XIDleTimeout             int
	// 	EgressRateLimitKbps          int
	// 	EgressRateLimitKbpsEnabled   bool
	// 	ExcludedNetworkIDs           []string
	// 	FecMode                      string
	// 	Forward                      string
	// 	FullDuplex                   bool
	// 	Isolation                    bool
	// 	LldpmedEnabled               bool
	// 	LldpmedNotifyEnabled         bool
	// 	MirrorPortIDX                int
	// 	NATiveNetworkID              string
	// 	Name                         string
	// 	OpMode                       string
	// 	PoeMode                      string
	// 	PortIDX                      int
	// 	PortKeepaliveEnabled         bool
	// 	PortProfileID                string
	// 	PortSecurityEnabled          bool
	// 	PortSecurityMACAddress       []string
	// 	PriorityQueue1Level          int
	// 	PriorityQueue2Level          int
	// 	PriorityQueue3Level          int
	// 	PriorityQueue4Level          int
	// 	QOSProfile                   *DeviceQOSProfile
	// 	SettingPreference            string
	// 	Speed                        int
	// 	StormctrlBroadcastastEnabled bool
	// 	StormctrlBroadcastastLevel   int
	// 	StormctrlBroadcastastRate    int
	// 	StormctrlMcastEnabled        bool
	// 	StormctrlMcastLevel          int
	// 	StormctrlMcastRate           int
	// 	StormctrlType                string
	// 	StormctrlUcastEnabled        bool
	// 	StormctrlUcastLevel          int
	// 	StormctrlUcastRate           int
	// 	StpPortMode                  bool
	// 	TaggedVLANMgmt               string
	// 	VoiceNetworkID               string
	// }
	// type args struct {
	// 	b []byte
	// }
	// tests := []struct {
	// 	name    string
	// 	fields  fields
	// 	args    args
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		dst := &DevicePortOverrides{
	// 			AggregateNumPorts:            tt.fields.AggregateNumPorts,
	// 			Autoneg:                      tt.fields.Autoneg,
	// 			Dot1XCtrl:                    tt.fields.Dot1XCtrl,
	// 			Dot1XIDleTimeout:             tt.fields.Dot1XIDleTimeout,
	// 			EgressRateLimitKbps:          tt.fields.EgressRateLimitKbps,
	// 			EgressRateLimitKbpsEnabled:   tt.fields.EgressRateLimitKbpsEnabled,
	// 			ExcludedNetworkIDs:           tt.fields.ExcludedNetworkIDs,
	// 			FecMode:                      tt.fields.FecMode,
	// 			Forward:                      tt.fields.Forward,
	// 			FullDuplex:                   tt.fields.FullDuplex,
	// 			Isolation:                    tt.fields.Isolation,
	// 			LldpmedEnabled:               tt.fields.LldpmedEnabled,
	// 			LldpmedNotifyEnabled:         tt.fields.LldpmedNotifyEnabled,
	// 			MirrorPortIDX:                tt.fields.MirrorPortIDX,
	// 			NATiveNetworkID:              tt.fields.NATiveNetworkID,
	// 			Name:                         tt.fields.Name,
	// 			OpMode:                       tt.fields.OpMode,
	// 			PoeMode:                      tt.fields.PoeMode,
	// 			PortIDX:                      tt.fields.PortIDX,
	// 			PortKeepaliveEnabled:         tt.fields.PortKeepaliveEnabled,
	// 			PortProfileID:                tt.fields.PortProfileID,
	// 			PortSecurityEnabled:          tt.fields.PortSecurityEnabled,
	// 			PortSecurityMACAddress:       tt.fields.PortSecurityMACAddress,
	// 			PriorityQueue1Level:          tt.fields.PriorityQueue1Level,
	// 			PriorityQueue2Level:          tt.fields.PriorityQueue2Level,
	// 			PriorityQueue3Level:          tt.fields.PriorityQueue3Level,
	// 			PriorityQueue4Level:          tt.fields.PriorityQueue4Level,
	// 			QOSProfile:                   tt.fields.QOSProfile,
	// 			SettingPreference:            tt.fields.SettingPreference,
	// 			Speed:                        tt.fields.Speed,
	// 			StormctrlBroadcastastEnabled: tt.fields.StormctrlBroadcastastEnabled,
	// 			StormctrlBroadcastastLevel:   tt.fields.StormctrlBroadcastastLevel,
	// 			StormctrlBroadcastastRate:    tt.fields.StormctrlBroadcastastRate,
	// 			StormctrlMcastEnabled:        tt.fields.StormctrlMcastEnabled,
	// 			StormctrlMcastLevel:          tt.fields.StormctrlMcastLevel,
	// 			StormctrlMcastRate:           tt.fields.StormctrlMcastRate,
	// 			StormctrlType:                tt.fields.StormctrlType,
	// 			StormctrlUcastEnabled:        tt.fields.StormctrlUcastEnabled,
	// 			StormctrlUcastLevel:          tt.fields.StormctrlUcastLevel,
	// 			StormctrlUcastRate:           tt.fields.StormctrlUcastRate,
	// 			StpPortMode:                  tt.fields.StpPortMode,
	// 			TaggedVLANMgmt:               tt.fields.TaggedVLANMgmt,
	// 			VoiceNetworkID:               tt.fields.VoiceNetworkID,
	// 		}
	// 		if err := dst.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
	// 			t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
	// 		}
	// 	})
	// }
}

func boolPtr(b bool) *bool {
	return &b
}

func stringPtr(s string) *string {
	return &s
}

func stringSlicePtr(s []string) *[]string {
	return &s
}
