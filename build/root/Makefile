GO := go
MKDIR := mkdir -p
RM := rm
OUT_DIR ?= target
BIN_DIR := $(OUT_DIR)/bin

all: bins

bins: $BIN_DIR/containers-ps

$BIN_DIR:
	${MKDIR} ${BIN_DIR}

$BIN_DIR/containers-ps: $BIN_DIR
	${GO} build -o ${BIN_DIR}/containers-ps cmd/containersps/containersps.go

clean:
	${RM} -f ${BIN_DIR}/containers-ps
