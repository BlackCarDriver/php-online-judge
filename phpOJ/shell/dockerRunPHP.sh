PHP=$(pwd)
	sudo docker run \
		--rm \
		-i \
		-v $PHP/UserCode:/UserCode \
		-v $PHP/SysTmpCode:/SysTmpCode \
		php \
		bin/bash /SysTmpCode/runPHPCode.sh
