all: cstat

clean:
	@rm -rf *.o
	@rm -rf cstat

cstat: main.o httpd.o netstat.o load.o
	gcc -o cstat $^

main.o: main.c
	gcc -c -o main.o main.c

httpd.o: httpd/httpd.c httpd/httpd.h
	gcc -c -o httpd.o httpd/httpd.c

netstat.o : netstat/netstat.c netstat/netstat.h
	gcc -c -o netstat.o netstat/netstat.c

load.o : load/load.c load/load.h
	gcc -c -o load.o load/load.c