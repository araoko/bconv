# bconv
Simple Bandwidth Unit Conversion (Windows App)


this program uses the walk library (github.com/lxn/walk)  for its GUI


to get rid of the console use:
go build -ldflags -H=windowsgui

the icon in the assets folder is no longer required to be deployed with the executable. it is houever required when creating the rsrc.syso file to set the icon for the executable.

to generate the syso file use:

rsrc -manifest bconv.manifest -ico assets\bconv.ico

