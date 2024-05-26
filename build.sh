#!/bin/bash
project=$(go list -m)
echo "Building project $project"
if [ -d "build" ] 
then
    rm -rf build/*
else
	mkdir build/
fi
go build -ldflags "-s -w"
echo "Copying files"
mv "$project" build/
mkdir -p ~/bin
cp -r build/* "$HOME"/bin/
echo "Add PATH="\$PATH:/\$HOME/bin" to .bashrc"
