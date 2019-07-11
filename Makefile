PROJECT_PATH  = github.com/tommy-sho/microservice-sample
PTARGET = api channel post room user video grpc/health/v1

protoc: clean-proto
	@echo $(PTARGET); \
	for t in $(PTARGET); do \
		protoc  \
			-Iproto \
			--proto_path=${GOPATH}/src \
			--proto_path=. \
			--proto_path=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
			--proto_path=${GOPATH}/src/github.com/tommy-sho/microservice-sample/ \
			--go_out=plugins=grpc:./proto \
			./proto/$$t/*.proto; \
	done

clean-proto:
	find .  |grep .pb.go | xargs rm