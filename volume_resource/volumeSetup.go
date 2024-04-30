// tfgen
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// volumeSetup.go, jfgratton : 2024-04-29

package volume_resource

import (
	"github.com/jeanfrancoisgratton/customError"
	"github.com/jeanfrancoisgratton/helperFunctions"
)

// This might seem a bit long-winded...
// We check at each string prompt if the value is empty. If not, we assign the variable's value
// To the appropriate struct member.
// We could directly return the string prompt for vallue to the strict member, but since
// The struct has a lot of "omitempty", this is a sure way to keep the JSON as short as possible
func (v VolumeResource) fetchResourceInfo() *customError.CustomError {
	v.Name = helperFunctions.GetStringValFromPrompt("What is the volume name ? ")
	helperFunctions.Blue("Note that all the other parameters are optional")
	v.Pool = helperFunctions.GetStringValFromPrompt("In which pool should this volume be hosted ? ")
	if v.Pool == "" {
		v.Pool = "default"
	}
	if s := helperFunctions.GetStringValFromPrompt("What is the volume source ? "); s != "" {
		v.Source = s
	}
	s := helperFunctions.GetValueFromPrompt("Please enter the volume size ? ")
	vs, ok := s.(uint64)
	if !ok {
		// Return a TypeAssertionError in case of type assertion failure
		ce := &customError.CustomError{Title: "Type assertion error",
			Message:  "Unable to convert value to uint64",
			Fatality: customError.Fatal}
		return ce
	} else {
		v.Size = vs
	}
	if bvid := helperFunctions.GetStringValFromPrompt("What is the base volume ID ?"); bvid != "" {
		v.BaseVolID = bvid
	}
	if bvn := helperFunctions.GetStringValFromPrompt("What is the base volume name ?"); bvn != "" {
		v.BaseVolID = bvn
	}
	if bvp := helperFunctions.GetStringValFromPrompt("What is the base volume pool ?"); bvp != "" {
		v.BaseVolID = bvp
	}
	return nil
}
