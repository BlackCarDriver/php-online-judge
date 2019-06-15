userResult=$(php ./code/test-UserCode.php)
result=$(echo "$userResult" | grep "^Error:")
if [ "$result" != "" ];then
    echo "$userReusult"
else
    userResult > 