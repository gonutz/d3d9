REM This script assumes that the environment variable FXC_COMPILER points to your fxc.exe.
REM If you have the Windows SDK installed the path should be something like:
REM C:\Program Files (x86)\Windows Kits\<version>\bin\x86\fxc.exe
"FXC_COMPILER" /WX /T vs_2_0 /Fo vso vs.txt
"FXC_COMPILER" /WX /T ps_2_0 /Fo pso ps.txt

REM bin2go converts binary files to Go byte slices
go get github.com/gonutz/bin2go
bin2go -s vso
bin2go -s pso