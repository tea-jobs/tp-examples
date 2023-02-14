############################################################
# Build cmds
############################################################

# Build jwt-tool cmd 
.PHONY: build-jwt-tool
build-jwt-tool:
	go build -o bin/jwt-tool cmd/jwt-tool/main.go

############################################################
# Run cmds
############################################################
.PHONY: run-jwt-tool
run-jwt-tool:
	bin/jwt-tool $(MODE) $(DATA) $(KEY)
