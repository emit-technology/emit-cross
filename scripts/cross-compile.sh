#!/bin/sh

workspace="$PWD/build/_workspace"

root="$PWD"

GOBIN="$PWD/build/bin"

workdir="$workspace/src/github.com/emit-technology"

projectName="emit-cross"

go get -u github.com/goware/modvendor

#
rm -rf "$workspace"
rm -rf "$GOBIN"

go mod vendor

modvendor -copy="**/*.c **/*.h **/*.proto **/*.so **/*.dyld **/*.A **/*.B0 **/*.B1 **/*.H **/*.L **/*.bin" -v

#
##
mkdir -p "$workdir/$projectName"
##
cp -R ` find .   \( -path "./build" -o -path "./.git" -o -path "./contract" -o -path "./.idea" -o -path "./images" -o -path "./go.mod" \) -prune -o -type d -depth 1 -print ` "$workdir/$projectName"

# Set up the environment to use the workspace.

GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$workdir/$projectName"

#
cp -r "$workdir/$projectName/vendor/github.com/sero-cash/go-czero-import/czero/lib_LINUX_AMD64_V3/"* "$workdir/$projectName/vendor/github.com/sero-cash/go-czero-import/czero/lib/"
cp -r "$workdir/$projectName/vendor/github.com/sero-cash/go-czero-import/czero/lib_DARWIN_AMD64/"* "$workdir/$projectName/vendor/github.com/sero-cash/go-czero-import/czero/lib/"
cp -r "$workdir/$projectName/vendor/github.com/sero-cash/go-czero-import/czero/lib_WINDOWS_AMD64/"* "$workdir/$projectName/vendor/github.com/sero-cash/go-czero-import/czero/lib/"
#


#
xgo  --targets=$1 -v  --dest=$GOBIN ./cmd/cross
#
cd "$root"

#
#rm -rf "$root/vendor"
#
#rm -rf "$workspace"

