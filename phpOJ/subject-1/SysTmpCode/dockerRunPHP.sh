PHP=$(pwd)
	sudo docker run \
		--rm \
		-i \
		-v $PHP/UserCode:/UserCode \
		-v $PHP/phpOJ/subject-1/SysTmpCode:/SysTmpCode \
		php:alpine \
		bin/sh /SysTmpCode/runPHPCode.sh
