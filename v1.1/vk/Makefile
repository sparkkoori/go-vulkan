all:
	c-for-go -ccdefs -out .. vk.yml

clean:
	rm -f doc.go types.go const.go
	rm -f cgo_helpers.go cgo_helpers.c cgo_helpers.h
	rm -f vk.go

test:
	go build
