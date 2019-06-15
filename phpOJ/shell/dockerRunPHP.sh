PHP=$(pwd)
	sudo docker run --name php \
		--rm \
		-i \
		-v $PHP/code:/code \
		php \
		php ./code/test-UserCode.php
