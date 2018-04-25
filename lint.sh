echo "+ gofmt"
(! go list -f "{{.Dir}}" ./... | xargs gofmt -s -d | grep '^') || exit 1
echo "+ gometalinter"
gometalinter ./... || exit 1
echo "+ staticcheck (failure is ignored)"
staticcheck ./... || echo "staticcheck failure is ignored"
