userResult=$(php /UserCode/zzm/test-UserCode.php)
result=$(echo "$userResult" | grep "error:")
if [ "$result" != "" ];then
    echo $userResult
else
    echo "$userResult" > /SysTmpCode/zzm/UserResult.txt
    sysResult=$(php /SysTmpCode/SystemCode.php)
    echo "$sysResult" > /SysTmpCode/SystemResult.txt
fi