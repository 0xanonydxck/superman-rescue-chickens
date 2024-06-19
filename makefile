test:
	go test -cover -v ./module/**
bench:
	go test -count=5 -bench=. ./module