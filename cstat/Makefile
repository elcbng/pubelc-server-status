CC = gcc
OBJ := main.o httpd.o netstat.o load.o
TOP_DIR := $(PWD)
OBJ_DIR := $(TOP_DIR)/obj
BIN_DIR := $(TOP_DIR)/bin
BIN := cstat
SUB_DIR := libjson \
			httpd \
			netstat \
			load

export CC OBJ TOP_DIR OBJ_DIR BIN_DIR BIN SUB_DIR

all: CHECKDIR main.o $(SUB_DIR) cstat

CHECKDIR: 
	mkdir -p $(OBJ_DIR) $(BIN_DIR)

$(SUB_DIR):ECHO
	make -C $@

cstat:
	make -C obj

ECHO:
	@echo $(SUB_DIR)
	@echo begin compile

clean:
	@rm -rf $(OBJ_DIR)/*.o $(BIN_DIR)

main.o: main.c
	$(CC) -c -o $(OBJ_DIR)/main.o main.c
