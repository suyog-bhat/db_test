# db_test
For 10000000 records (ec2 instance)

RocksDB----
time took for write: 4.342664708s
Num of records: 69905 17.378042ms
LevelDb-----
time took for write 2m37.513682112s
Num of records: 69905 11.478167ms

rocksDb setup guide in ubuntu
dependencies:
Upgrade your gcc to version at least 4.8 to get C++11 support.
Install gflags. First, try: sudo apt-get install libgflags-dev 
If this doesn't work and you're using Ubuntu, here's a nice tutorial: (http://askubuntu.com/questions/312173/installing-gflags-12-04)

Install snappy. This is usually as easy as: sudo apt-get install libsnappy-dev

Install zlib. Try: sudo apt-get install zlib1g-dev

Install bzip2: sudo apt-get install libbz2-dev

Install lz4: sudo apt-get install liblz4-dev

Install zstandard: sudo apt-get install libzstd-dev

installing rocksDb:
git clone https://github.com/facebook/rocksdb.git
cd rocksdb
make static_lib

setup wrapper grocksdb
CGO_CFLAGS="-I/path/to/rocksdb/include" \
CGO_LDFLAGS="-L/path/to/rocksdb -lrocksdb -lstdc++ -lm -lz -lsnappy -llz4 -lzstd" \
  go build
  
if go build did not work try : go build -tags builtin_static
