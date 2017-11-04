# go-eccodes
Go wrapper for [ecCodes](https://software.ecmwf.int/wiki/display/ECC/ecCodes+Home)

- [go-eccodes](#go-eccodes)
    - [Build and install ecCodes C library](#build-and-install-eccodes-c-library)
        - [Build and install ecCodes C library for Linux](#build-and-install-eccodes-c-library-for-linux)
            - [Install development tools](#install-development-tools)
            - [Build and install `zlib`](#build-and-install-zlib)
            - [Build and install `libpng`](#build-and-install-libpng)
            - [Build and install `libaec`](#build-and-install-libaec)
            - [Build end install `libjpeg`](#build-end-install-libjpeg)
            - [Build end install `libopenjpeg2`](#build-end-install-libopenjpeg2)
            - [Build end install `libjasper`](#build-end-install-libjasper)
            - [Build end install `libeccodes`](#build-end-install-libeccodes)
        - [Build and install ecCodes C library for Windows (using MSYS2/MINGW64)](#build-and-install-eccodes-c-library-for-windows-using-msys2mingw64)
            - [Install `MSYS2`](#install-msys2)
            - [Install Mingw-w64 (run `MSYS2 MinGW 64-bit` shell)](#install-mingw-w64-run-msys2-mingw-64-bit-shell)
            - [Install ecCodes dependencies (run `MSYS2 MinGW 64-bit` shell)](#install-eccodes-dependencies-run-msys2-mingw-64-bit-shell)
            - [Build and install `libaec`](#build-and-install-libaec)
            - [Build end install `libeccodes`](#build-end-install-libeccodes)

## Build and install ecCodes C library

### Build and install ecCodes C library for Linux

#### Install development tools

```bash
sudo apt-get install gcc make cmake libtool
```

#### Build and install `zlib`

source: [zlib](https://zlib.net/)

```bash
cd ./contrib
tar -xzf zlib-1.2.11.tar.gz
cd zlib-1.2.11
make distclean
./configure --static
make
sudo make install
cd ..
rm -r ./zlib-1.2.11
cd ..
```

#### Build and install `libpng`

source: [libpng](https://libpng.sourceforge.io/index.html)

```bash
cd ./contrib
tar -xzf libpng-1.6.34.tar.gz
cd libpng-1.6.34
./configure --disable-shared
make check
sudo make install
cd ..
rm -r ./libpng-1.6.34
cd ..
```

#### Build and install `libaec`

source: [libaec](https://gitlab.dkrz.de/k202009/libaec)

```bash
cd ./contrib
tar -xzf libaec-1.0.2.tar.gz
cd libaec-1.0.2
mkdir build
cd build
../configure --disable-shared
make check
sudo make install
cd ../..
rm -r ./libaec-1.0.2
cd ..
```

#### Build end install `libjpeg`

source: [libjpeg](http://www.ijg.org/)

```bash
cd ./contrib
tar -xzf jpegsrc.v9b.tar.gz
cd jpeg-9b
./configure --disable-shared
make
make test
sudo make install
cd ..
rm -r ./jpeg-9b
cd ..
```

#### Build end install `libopenjpeg2`

source: [libopenjpeg2](http://www.openjpeg.org/)

```bash
cd ./contrib
tar -xzf openjpeg-2.1.2.tar.gz
cd openjpeg-2.1.2
mkdir build
cd build
cmake -DBUILD_SHARED_LIBS:bool=OFF -DBUILD_THIRDPARTY:bool=ON ..
make
sudo make install
cd ../..
rm -r ./openjpeg-2.1.2
cd ..
```

#### Build end install `libjasper`

source: [libjasper](https://www.ece.uvic.ca/~frodo/jasper/)

```bash
cd ./contrib
tar -xzf jasper-version-2.0.14.tar.gz
mkdir build
cd build
cmake -DJAS_ENABLE_SHARED=false ../jasper-version-2.0.14
make clean all
make test
sudo make install
cd ..
rm -r ./build
rm -r ./jasper-version-2.0.14
cd ..
```

#### Build end install `libeccodes`

source: [libeccodes](https://software.ecmwf.int/wiki/display/ECC/ecCodes+Home)

```bash
cd ./contrib
tar -xzf eccodes-2.5.0-Source.tar.gz
mkdir build
cd build
cmake -DBUILD_SHARED_LIBS=OFF -DENABLE_NETCDF=OFF -DENABLE_JPG=ON -DENABLE_PNG=ON -DENABLE_AEC=ON \
    -DENABLE_PYTHON=OFF -DENABLE_FORTRAN=OFF -DENABLE_MEMFS=ON ../eccodes-2.5.0-Source
make
ctest
sudo make install
cd ..
rm -r ./build
rm -r ./eccodes-2.5.0-Source
cd ..
```

### Build and install ecCodes C library for Windows (using MSYS2/MINGW64)

#### Install `MSYS2`

source: [MSYS2](http://www.msys2.org/)

- download [installer](http://repo.msys2.org/distrib/x86_64/msys2-x86_64-20161025.exe) for x86_64

- install MSYS2 following the [guide](http://www.msys2.org/)

- uncomment `MSYS=winsymlinks:nativestrict` everywhere to enable symbol links

#### Install Mingw-w64 (run `MSYS2 MinGW 64-bit` shell)

```bash
pacman -S --needed base-devel git mingw-w64-x86_64-toolchain mingw-w64-x86_64-cmake
```

#### Install ecCodes dependencies (run `MSYS2 MinGW 64-bit` shell)

```bash
pacman -S mingw-w64-x86_64-zlib
curl -O http://repo.msys2.org/mingw/x86_64/mingw-w64-x86_64-openjpeg2-2.1.2-2-any.pkg.tar.xz
pacman -U mingw-w64-x86_64-openjpeg2-2.1.2-2-any.pkg.tar.xz
pacman -S mingw-w64-x86_64-jasper
```

#### Build and install `libaec`

source: [libaec](https://gitlab.dkrz.de/k202009/libaec)

- extract `libaec-1.0.2.tar.gz` to MSYS2 user home directory

- replace original files by files from `contrib\MSYS2\patches\libaec-1.0.2`

- run `MSYS2 MinGW 64-bit` shell and execute commands to build and install:

```bash
tar -xzf libaec-1.0.2.tar.gz
cd libaec-1.0.2
mkdir build
cd build
../configure
make check
make install
```

#### Build end install `libeccodes`

source: [libeccodes](https://software.ecmwf.int/wiki/display/ECC/ecCodes+Home)

- extract `eccodes-2.5.0-Source.tar.gz` to MSYS2 user home directory

- replace original files by files from `contrib\MSYS2\patches\eccodes-2.5.0-Source`

- run `MSYS2 MinGW 64-bit` shell and execute commands to build and install:

```bash
mkdir build
cd build
cmake -G "MSYS Makefiles" -DCMAKE_INSTALL_PREFIX=/mingw64 -DDISABLE_OS_CHECK=ON -DENABLE_NETCDF=OFF -DENABLE_JPG=ON -DENABLE_PNG=ON -DENABLE_AEC=ON \
    -DENABLE_PYTHON=OFF -DENABLE_FORTRAN=OFF -DENABLE_MEMFS=OFF ../eccodes-2.5.0-Source
make
make install
```

- set `ECCODES_DEFINITION_PATH` environment variable to `C:\Bin\msys64\mingw64\share\eccodes\definitions`
