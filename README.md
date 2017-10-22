# go-eccodes
Go wrapper for [ecCodes](https://software.ecmwf.int/wiki/display/ECC/ecCodes+Home)

## Build and install ecCodes C library

### Build and install dependencies

#### Install development tools

```bash
sudo apt-get install gcc make cmake libtool
```

#### Build and install [zlib](https://zlib.net/)

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

#### Build and install [libpng](https://libpng.sourceforge.io/index.html)

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

#### Build and install [libaec](https://gitlab.dkrz.de/k202009/libaec)

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

#### Build end install [libjpeg](http://www.ijg.org/)

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

#### Build end install [libopenjpeg2](http://www.openjpeg.org/)

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

#### Build end install [libjasper](https://www.ece.uvic.ca/~frodo/jasper/)

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

#### Build end install [libeccodes](https://software.ecmwf.int/wiki/display/ECC/ecCodes+Home)

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
