#It shell run php code and save the result in text file
#username should append to the command when run it shell, such as: $./run-commit.sh demo

########################################################

codePath="/home/ubuntu/Workplace/php-online-judge/tmpCode/"
userName=$1
workPlace="${codePath}${userName}"

sudo docker run -it --rm \
-v $workPlace:/workplace \
php:alpine \
sh -c 'cd workplace && php commit.php > commit-result.txt'
