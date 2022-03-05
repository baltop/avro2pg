echo avro2pg kill
PID=`ps -ef | grep "avro2pg" | grep -v 'grep' | awk '{print $2}'`
if [ -z $PID ]; then
	echo "no process is running"
else
	echo "kill process"
	kill -9 $PID
fi


