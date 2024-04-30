// tfgen
// Écrit par J.F.Gratton (jean-francois@famillegratton.net)
// setup.go, jfgratton : 2024-04-29

package pool_resource

import "github.com/jeanfrancoisgratton/customError"

func (p PoolResource) fetchResourceInfo() *customError.CustomError {
	return nil
}

/*
func GetPoolName() (string, *customError.CustomError) {
	//ce := &customError.CustomError{}
	// Call your function from the helperFunctions package
	promptValue := helperFunctions.GetValueFromPrompt("Please enter the pool name: ")

	// Perform a type assertion to convert interface{} to string
	poolname, ok := promptValue.(string)
	if !ok {
		// Return a TypeAssertionError in case of type assertion failure
		ce := &customError.CustomError{Title: "Type assertion error",
			Message:  "Unable to convert value to string",
			Fatality: customError.Fatal}
		return "", ce
	}

	return poolname, nil
}
*/
