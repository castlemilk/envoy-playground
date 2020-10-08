
build:
	docker build -t envoy .
start: build
	docker run -v ${PWD}/config/envoy.yaml:/etc/envoy/envoy.yaml -p 9901:9901 envoy 