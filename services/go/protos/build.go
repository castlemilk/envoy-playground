// Calculator Service
//go:generate protoc ./calculator/calculator.proto -I calculator -I third_party/ --gofast_out=plugins=grpc:./ --validate_out=lang=go:./ --grpc-gateway_out=logtostderr=false:./calculator --swagger_out=logtostderr=false:./calculator
package protos 
