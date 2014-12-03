all: goCode
cCode.o: cCode.c
	gcc -g -fPIC -c cCode.c

cCode.so: cCode.o
	gcc -shared -o libccode.so cCode.o

goCode: goCode.go cCode.so
	go build -ldflags="-r ./" goCode.go
	
clean:
	rm cCode.o
	rm libccode.so
	rm goCode
