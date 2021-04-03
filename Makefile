sonar:
	go test -coverprofile=./coverage.out ./...; 
	sonar-scanner;