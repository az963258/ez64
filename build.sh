#!/bin/bash
cd src
go build main.go
mv main ez64

echo "do you want to move ez64 to '/usr/bin'? [y/n]"
read answer

case $answer in
    y|Y ) echo "okay, moving to '/usr/bin'" ;;
    n|N ) echo "okay, ez64 will stay in src folder" ; exit 1;;
    * ) echo "wrong answer, y/n only" ; exit 2;;
esac

sudo mv ez64 /usr/bin
