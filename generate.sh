#!/bin/sh

if [ $# -ne 1 ];then
	echo "Usage: $0 DIR-NAME"
	exit 1
fi

mkdir $1
cp template/* $1/

echo create $1
