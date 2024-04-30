// tfgen
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : volume_resource.go
// Original timestamp : 2024/04/27 15:16

package volume_resource

// Volume :
// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs/resources/volume
type VolumeResource struct {
	Name        string `json:"name"`
	Pool        string `json:"pool,omitempty"`
	Source      string `json:"source,omitempty"`
	Size        uint64 `json:"size,omitempty"`
	BaseVolID   string `json:"basevolumeid,omitempty"`
	BaseVolName string `json:"basevolumename,omitempty"`
	BaseVolPool string `json:"basevolumepool,omitempty"`
}
