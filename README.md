# tfgen

A package to generate Terraform config files (".tf files"), used in provisioning VMs
___

## How to install 

**fixme** document this section
## How does it work

There are two operating modes: you either provide the whole data to the package, or wish to be prompted for all the needed values.
<br><br>
### Mode 1: data is provided
The data must be in JSON form, as shown below <<== ** fixme : add info **

### Mode 2: you get prompted
The library will prompt for all the various variables needed to generate a tf file.


Whatever the selected mode, the library will generate a TF file and store it at the place of your choosing (a parameter you pass to the package)


## Package layout and dependencies

### Dependencies

This package relies on some other packages I've written, on top of more other "mainstream" packages:
- https://github.com/jeanfrancoisgratton/customErrors: customizable error messages; you could easily get rid of this if needed
- https://github.com/jeanfrancoisgratton/helperFunctions: misc functions sprinkled across the package

### Package layout

This package leans heavily on the Terraform API, and especially dmacvicar's Terraform libvirt provider (https://github.com/dmacvicar/terraform-provider-libvirt/tree/main)

Each resource defined in that provider is mapped in a "subpackage":
- cloud-init
- domain
- network <-- this one is not yet fully supported
- volume
- pool

Within each subpackage (subdir) you will always find the following files :
- *Config : the necessary structures to that resource
- *Setup : all the functions/interfaces needed to process information in and out of that structure

## How to use
**fime** document this section