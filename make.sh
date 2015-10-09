go get github.com/codegangsta/cli
go get gopkg.in/pipe.v2
go get gopkg.in/yaml.v2

for GOOS in darwin linux; do
    export GOOS=$GOOS
    for GOARCH in 386 amd64; do
        export GOARCH=$GOARCH
        go build -v -o bin/android-update-$GOOS-$GOARCH
    done
done
