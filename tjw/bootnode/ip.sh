ip=$(ifconfig|grep inet|grep -v inet6|grep broadcast|awk '{print $2}')

echo $ip

bootnode_addr=enode://"$(grep enode bootnode.log|tail -n 1|awk -F '://' '{print $2}'|awk -F '@' '{print $1}')""@$ip:30301"
echo $bootnode_addr