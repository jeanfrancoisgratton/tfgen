// tfgen
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original directory and filename : pool_resource.go
// Original timestamp : 2024/04/27 15:14

package tfgen

// Pool :
// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs/resources/pool
// Note : the only supported pool type for now is "dir"
type PoolResource struct {
	Name string `json:"name"`
	Type string `json:"type"`
	// Note: this one should be optional for non-dir types, but since only dir types are supported
	// for now, we won't add the "omitempty" directive
	Path string `json:"path"`
}
