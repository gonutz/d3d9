REM you have to change this path to where your shader compiler (fxc.exe) is
"C:\Program Files (x86)\Windows Kits\8.1\bin\x86\fxc.exe" /WX /T vs_2_0 /Fo vso vs.txt
"C:\Program Files (x86)\Windows Kits\8.1\bin\x86\fxc.exe" /WX /T ps_2_0 /Fo pso ps.txt
go get github.com/gonutz/bin2go
bin2go -s vso
bin2go -s pso