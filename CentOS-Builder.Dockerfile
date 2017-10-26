FROM centos:latest

RUN yum -y install git gcc make libtool && \
    yum -y install epel-release && \
    yum -y install cmake3 && \
    yum clean all

ENV GOLANG_VERSION 1.9.2
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 de874549d9a8d8d8062be05808509c09a88a248e77ec14eb77453530829ac02b
RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz
ENV PATH /usr/local/go/bin:$PATH

WORKDIR /root

# build and install zlib
ADD ./contrib/zlib-1.2.11.tar.gz ./
RUN cd ./zlib-1.2.11 && \
    make distclean && \
    CFLAGS=-std=c99 ./configure --static && \
    make install && \
    cd .. && \
    rm -r ./zlib-1.2.11

# build and install libpng
ADD ./contrib/libpng-1.6.34.tar.gz ./
RUN cd libpng-1.6.34 && \
    CFLAGS=-std=c99 ./configure --disable-shared && \
    make check install && \
    cd .. && \
    rm -r ./libpng-1.6.34

# build and install libaec
ADD ./contrib/libaec-1.0.2.tar.gz ./
RUN cd libaec-1.0.2 && \
    mkdir build && \
    cd build && \
    CFLAGS=-std=c99 ../configure --disable-shared && \
    make check install && \
    cd ../.. && \
    rm -r ./libaec-1.0.2

# build and install libjpeg
ADD ./contrib/jpegsrc.v9b.tar.gz ./
RUN cd jpeg-9b && \
    CFLAGS=-std=c99 ./configure --disable-shared && \
    make && \
    make test install && \
    cd .. && \
    rm -r ./jpeg-9b

# build and install libopenjpeg2 -DCFLAGS=-std=c99
ADD ./contrib/openjpeg-2.1.2.tar.gz ./
RUN cd openjpeg-2.1.2 && \
    mkdir build && \
    cd build && \
    cmake3 -DBUILD_SHARED_LIBS:bool=OFF -DBUILD_THIRDPARTY:bool=ON .. && \
    make install && \
    cd ../.. && \
    rm -r ./openjpeg-2.1.2

# build and install libjasper
ADD ./contrib/jasper-version-2.0.14.tar.gz ./
RUN mkdir build && \
    cd build && \
    cmake3 -DJAS_ENABLE_SHARED=false ../jasper-version-2.0.14 && \
    make clean all && \
    make install && \
    cd .. && \
    rm -r ./build && \
    rm -r ./jasper-version-2.0.14

# build and install libeccodes
ADD ./contrib/eccodes-2.5.0-Source.tar.gz ./
RUN mkdir build && \
    cd build && \
    cmake3 -DBUILD_SHARED_LIBS=OFF -DENABLE_NETCDF=OFF -DENABLE_JPG=ON -DENABLE_PNG=ON -DENABLE_AEC=ON \
        -DENABLE_PYTHON=OFF -DENABLE_FORTRAN=OFF -DENABLE_MEMFS=ON ../eccodes-2.5.0-Source && \
    make && \
    ctest3 && \
    make install && \
    cd .. && \
    rm -r ./build && \
    rm -r ./eccodes-2.5.0-Source

ENV GOPATH /root/go
ENV PATH /root/go/bin:$PATH

RUN go get -u github.com/golang/dep/cmd/dep
