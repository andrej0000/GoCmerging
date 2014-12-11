
__global__ void kernel(int * vals, int size){
	int tid = blockDim.x * blockIdx.x + threadIdx.x;
	int stride = blockDim.x * gridDim.x;
	
	for (; tid < size; tid += stride) {
		vals[tid] *= 2;
	}
}

extern "C"
int foo(int size){
	int * vals;
	int * devVals;

	cudaMallocHost((void**)&vals, size * sizeof(vals[0]));
	cudaMalloc((void**)&devVals, size * sizeof(devVals[0]));

	int i = 0;
	for (; i < size; i++) {
		vals[i] = i;
	}
	cudaMemcpy(devVals, vals, size * sizeof(vals[0]), cudaMemcpyHostToDevice);
	kernel<<<1, 16, 16>>>(devVals, size);
	cudaMemcpy(vals, devVals, size * sizeof(devVals[0]), cudaMemcpyDeviceToHost);
	i = 0;
	int tmp = 0;
	for (; i < size; i++) {
		tmp += vals[i];
	}
	return tmp;
}
