#!/bin/bash

for file in [0-9]*; do
  if [ -d "$file" ] ; then
		echo Testing $file...
		cd $file
		gotestdox
		cd ..
	fi
done
