# tfgen

A library to generate terraform config files (".tf files"), used in provisioning VMs
___
## How does it work

There are two operating modes: you either provide the whole data to the package, or wish to be prompted for all the needed values.
<br><br>
### Mode 1: data is provided
The data must be in JSON form, as shown below

### Mode 2: you get prompted
The library will prompt for all the various variables needed to generate a tf file.


Whatever the selected mode, the library will generate a TF file and store it at the place of your choosing (a parameter you pass to the package)