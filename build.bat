@echo off

go install GopherGL/src/window
go install GopherGL/src/gfx
go install GopherGL/src/camera
go install GopherGL/src/input
go build -o build/GopherGL.exe src/main.go

pushd build
start GopherGL.exe
popd