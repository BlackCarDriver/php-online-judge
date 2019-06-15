PHP=$(pwd)
	sudo docker run \
		--rm \
		-i \
		-v $PHP/UserCode:/UserCode \
		-v $PHP/SysTmpCode:/SysTmpCode \
		php:alpine \
		bin/sh /SysTmpCode/runPHPCode.sh
