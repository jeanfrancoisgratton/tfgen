// tfgen
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename : network_resource.go
// Original timestamp : 2024/04/27 15:17

package tfgen

// **************************************************************************************************
// IMPORTANT NOTE:
// This resource is currently not supported. The following is a stub for whenever I get to work on it
// **************************************************************************************************

// Network :
// https://registry.terraform.io/providers/dmacvicar/libvirt/latest/docs/resources/network
type NetworkResource struct {
	Name   string `json:"name"`
	Domain string `json:"domain,omitempty"`
}
