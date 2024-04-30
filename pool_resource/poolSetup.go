// tfgen
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// volumeSetup.go, jfgratton : 2024-04-29

package pool_resource

import (
	"fmt"
	"github.com/jeanfrancoisgratton/customError"
	"github.com/jeanfrancoisgratton/helperFunctions"
)

func (p PoolResource) fetchResourceInfo() *customError.CustomError {
	p.Name = helperFunctions.GetStringValFromPrompt("What is the pool name ? ")
	helperFunctions.Yellow("Please note:")
	fmt.Println("            Whatever you answer to the next question, the type will be 'dir'")
	fmt.Println("            As this is the only type currently supported.")
	p.Type = helperFunctions.GetStringValFromPrompt("What is the pool type ? ")
	p.Type = "dir"
	p.Path = helperFunctions.GetStringValFromPrompt("What is the pool path ? ")
	return nil
}
