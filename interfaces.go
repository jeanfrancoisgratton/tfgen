// tfgen
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// interfaces.go, jfgratton : 2024-04-30

package tfgen

import "github.com/jeanfrancoisgratton/customError"

// This will be all interfaces will be hosted... eventually

type ResourceInfoFetcher interface {
	fetchResourceInfo() *customError.CustomError
}
