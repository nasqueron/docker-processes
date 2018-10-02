GO := go
MKDIR := mkdir -p
RM := rm
OUT_DIR ?= target
BIN_DIR := $(OUT_DIR)/bin
BINS=${BIN_DIR}/containers-ps ${BIN_DIR}/what-container

all: bins

bins: ${BINS}

${BIN_DIR}:
	${MKDIR} ${BIN_DIR}

${BIN_DIR}/containers-ps: ${BIN_DIR}
	${GO} build -o ${BIN_DIR}/containers-ps cmd/containersps/containersps.go

${BIN_DIR}/what-container: ${BIN_DIR}
	${GO} build -o ${BIN_DIR}/what-container cmd/whatcontainer/whatcontainer.go

clean:
	${RM} -f ${BINS}
