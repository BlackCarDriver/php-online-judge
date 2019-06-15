userResult=$(php /UserCode/test-UserCode.php)
result=$(echo "$userResult" | grep "error:")
if [ "$result" != "" ];then
    echo $userResult
else
    echo $userResult > /UserCode/UserResult.txt
    sysResult=$(php /SysTmpCode/demo/test-SystemCode.php)
    echo $sysResult > /SysTmpCode/SystemResult.txt
fi