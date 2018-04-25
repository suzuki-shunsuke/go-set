set -e

echo "+ gofmt"
! git ls-files | grep ".go$" | xargs gofmt -s -d | grep '^'
echo "+ gometalinter"
gometalinter ./...
echo "+ staticcheck (failure is ignored)"
staticcheck ./... || echo "staticcheck failure is ignored"
