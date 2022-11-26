module waas

go 1.19

replace waas/km => ./km
replace waas/capabilities/wait => ./capabilities/wait

require waas/km v0.0.0
require waas/capabilities/wait v0.0.0

require (
	github.com/wapc/wapc-go v0.5.5
	karmem.org v1.2.9
)

require (
	github.com/Workiva/go-datastructures v1.0.53 // indirect
	github.com/tetratelabs/wazero v1.0.0-pre.3 // indirect
)
