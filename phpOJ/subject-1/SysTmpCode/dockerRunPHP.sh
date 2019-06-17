PHP=$(pwd)
	sudo docker run \
		--rm \
		-i \
		-v $PHP/phpOJ/subject-1/UserCode:/UserCode \
		-v $PHP/phpOJ/subject-1/SysTmpCode:/SysTmpCode \
		php:alpine \
		bin/sh /SysTmpCode/runPHPCode.sh
