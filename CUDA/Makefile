all: goCode

libkernel.so: kernel.cu
	nvcc -ccbin=g++-4.8 --shared --compiler-options -fPIC -arch=sm_35 kernel.cu -o libkernel.so

goCode: libkernel.so goCode.go
	go build -ldflags="-r ./" goCode.go

clean:
	rm -f goCode
	rm -f libkernel.so
