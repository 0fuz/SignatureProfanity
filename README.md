# SignatureProfanity

Tool to find signature from given solidity functionName(address,uint) using all CPU-cores.

Suitable for finding exact signature to reduce gas usage a bit.

### Compile from source code
```shell
# install golang
sudo apt install golang -y
# or
sudo snap install go
# or visit https://go.dev/doc/install
# or enter into golang docker container

# clone repository
git clone https://github.com/0fuz/SignatureProfanity

# install dependencies 
go mod tidy

# compile
go build
```

Run
```shell
./SignatureProfanity "<fnName(address)>" <signature1,signature2,signature3,...>

./SignatureProfanity "fn_example_name_1872aff099(address)" 0x00000001,0x00000002,0x00000003,0x00000004,0x00000005,0x00000006,0x00000007,0x00000008,0x00000009,0x00000010
# output
# fn_example_name_1872aff099(address) 0x00000003
# can be proved here https://emn178.github.io/online-tools/keccak_256.html
```
