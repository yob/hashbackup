# hashbackup

An experiment in saving various hash values (md5, sha1, etc) for a directory
structure.

Most of my day job time is spent in ruby and javascript, and I'm interested in
exploring languages with a better concurrency story.

## Installation

    cd $GOPATH
    go get github.com/yob/hashbackup
    go install github.com/yob/hashbackup

## Usage

The executable takes a single argument - a directory path - and prints out md5, sha1, byes, path

    $ ./bin/hashbackup bin/
    c8ac9a646729a8ee4999d045c068be58        08f6d79b503a1e31de55cb6e6837f87fa9448302        6942576 /home/jh/go/bin/cfssl
    da2069a579423789d76048e39bf22f4e        719f97182e31ff89a8e2e39a3b083d3da6a44c43        9039056 /home/jh/go/bin/godep
    02d2a96fa584b74e4bafa1caf5b62766        0fc9345ad0f14956a0b69c9640ca48ef1a645c95        2175419 /home/jh/go/bin/hashbackup
