# vim: noet:sw=2:ts=2



build:
	make -C Go build 
	make -C Docker build

push: build
	make -C Docker push


clean:
	make -C Go clean
