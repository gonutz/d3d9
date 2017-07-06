REM see README.md for how to use this script
set FXC="C:\Program Files (x86)\Windows Kits\8.1\bin\x86\fxc.exe"
%FXC% /WX /T vs_2_0 /Fo vs.object vs.code
%FXC% /WX /T ps_2_0 /Fo ps.object ps.code
go get github.com/gonutz/bin2go
bin2go -s -v vso -o vso.go vs.object
bin2go -s -v pso -o pso.go ps.object
go build -ldflags -H=windowsgui