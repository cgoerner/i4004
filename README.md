# Intel 4004 emulator

An emulator for the Intel 4004, the first commercially produced microprocessor.

Build:
```
go build ./cmd/i4004
go build ./cmd/i4004-server
```

Run:
```
./i4004 examples/memtest.rom
# or
./i4004-server examples/fulltest.rom
```

Start the web server:
```
./i4004-server
```

Install:
```
go install github.com/cgoerner/i4004/cmd/i4004
i4004 examples/fulltest.rom

# or

go install github.com/cgoerner/i4004/cmd/i4004-server
i4004-server
```