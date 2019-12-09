cd src
gox -osarch="linux/amd64" -output="..\bin\service-{{.OS}}-{{.Arch}}"
cd ..