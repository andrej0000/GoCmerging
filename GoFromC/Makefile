all: goCode

libccode.so: cCode.c
	gcc -fPIC -shared cCode.c -o libccode.so

goCode: libccode.so goCode.go cCode.h
	go build -ldflags="-r ./" goCode.go

clean:
	rm libccode.so
	rm goCode
