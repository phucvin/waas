module hello

go 1.19

replace "waas/km" => ../../km
require waas/km v0.0.0

require (
	github.com/wapc/wapc-guest-tinygo v0.3.3
	karmem.org v1.2.9
)

require (
	golang.org/x/crypto v0.0.0-20220513210258-46612604a0f9 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
)
