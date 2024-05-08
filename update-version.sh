#!/usr/bin/bash

sed -i "s/$1/$2/g" PKGBUILD
git commit -am "build(arch): update pkgbuild for $2"
git tag $2
git push
git push origin $2

