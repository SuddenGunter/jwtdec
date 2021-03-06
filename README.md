# jwtdec

Simple CLI tool for JWT payload parsing in your terminal.

No validation, no configuration, no dependencies in your OS required.

I wrote this because shell tool is more convenient then opening new browser tab with https://jwt.io each time I need to parse JWT.

![Screenshot](https://github.com/SuddenGunter/jwtdec/blob/master/1585778760_0106_02042020_421x193.png)

## Installation

### Go modules
```shell script
go get -u github.com/SuddenGunter/jwtdec
```
### Prebuilt binaries
See latest release tag in this repo and just download executable for your OS.

## Usage

Colorful output:

```shell script
./jwtdec eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU
## will output in color:
## {
##  "foo": "bar",
##  "nbf": 1444478400
## }
```

No color mode:
```shell script
./jwtdec -nc eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU
## will output:
## {
##  "foo": "bar",
##  "nbf": 1444478400
## }
```
