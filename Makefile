run: 
	go build -o ./concept ./vendor/lux-concept/ && ./concept


goinstall:
	go install -u github.com/luxengine/lux/...
	go install -u github.com/luxengine/gl


goget:
	go get github.com/luxengine/lux/...
	go get github.com/luxengine/tornago