#!/bin/bash

current_path=`pwd`
case "`uname`" in
    Linux)
		bin_abs_path=$(readlink -f $(dirname $0))
		;;
	*)
		bin_abs_path=`cd $(dirname $0); pwd`
		;;
esac
base=${bin_abs_path}/..
otterNodeIdFile=$base/conf/nid
logback_configurationFile=$base/conf/logback.xml
export LANG=en_US.UTF-8

if [ ! -d $base/logs/node ] ; then
	mkdir -p $base/logs/node
fi

ARIA2C=$base/aria2c

## set java path
if [ -z "$JAVA" ] ; then
  JAVA=$(which java)
fi

if [ -z "$JAVA" ]; then
	echo "Cannot find a Java JDK. Please set either set JAVA or put java (>=1.5) in your PATH." 2>&2
	exit 1
fi

case "$#"
in
0 )
	;;
1 )
	var=$*
	if [ -d $var ]
	then
		otterNodeIdFile=$var
        logback_configurationFile=$base/conf/logback.xml
	elif [ -f $var ] ; then
		otterNodeIdFile=$base/conf/nid
        logback_configurationFile=$var
	else
		echo "THE PARAMETER IS NOT CORRECT.PLEASE CHECK AGAIN."
        exit
	fi;;
2 )
	var1=$1
	var2=$2
	if [ -d $var1 -a -f $var2 ] ; then
		otterNodeIdFile=$var1
		logback_configurationFile=$var2
	elif [ -d $var2 -a -f $var1 ] ; then
		otterNodeIdFile=$var2
		logback_configurationFile=$var1
	else
		if [ "$1" = "debug" ]; then
			DEBUG_PORT=$2
			DEBUG_SUSPEND="n"
			JAVA_DEBUG_OPT="-Xdebug -Xnoagent -Djava.compiler=NONE -Xrunjdwp:transport=dt_socket,address=$DEBUG_PORT,server=y,suspend=$DEBUG_SUSPEND"
		fi
     fi;;
* )
	echo "THE PARAMETERS MUST BE TWO OR LESS.PLEASE CHECK AGAIN."
	exit;;
esac

# str=`file $JAVA | grep 64-bit`
# if [ -n "$str" ]; then
# 	JAVA_OPTS="-server -Xms2048m -Xmx3072m -Xmn1024m -XX:SurvivorRatio=2 -XX:PermSize=96m -XX:MaxPermSize=256m -Xss256k -XX:-UseAdaptiveSizePolicy -XX:MaxTenuringThreshold=15 -XX:+DisableExplicitGC -XX:+UseConcMarkSweepGC -XX:+CMSParallelRemarkEnabled -XX:+UseCMSCompactAtFullCollection -XX:+UseFastAccessorMethods -XX:+UseCMSInitiatingOccupancyOnly -XX:+HeapDumpOnOutOfMemoryError"
# else
	JAVA_OPTS="-server -Xms1024m -Xmx1024m -XX:NewSize=256m -XX:MaxNewSize=256m -XX:MaxPermSize=128m "
# fi

JAVA_OPTS=" $JAVA_OPTS -Djava.awt.headless=true -Djava.net.preferIPv4Stack=true -Dfile.encoding=UTF-8"
OTTER_OPTS="-DappName=otter-node -Ddubbo.application.logger=slf4j -Dlogback.configurationFile=$logback_configurationFile -Dnid=$(cat $otterNodeIdFile)"

if [ -e $otterNodeIdFile -a -e $logback_configurationFile ]
then
	for i in $base/lib/*;
	do CLASSPATH=$i:"$CLASSPATH";
	done
	CLASSPATH="$base/conf:$CLASSPATH";

	echo LOG CONFIGURATION : $logback_configurationFile
	echo Otter nodeId file : $otterNodeIdFile
	echo CLASSPATH :$CLASSPATH

	echo "cd to $bin_abs_path for workaround relative path"
	cd $bin_abs_path

	nohup $JAVA $JAVA_OPTS $JAVA_DEBUG_OPT $OTTER_OPTS -classpath .:$CLASSPATH com.alibaba.otter.node.deployer.OtterLauncher 1>>$base/logs/node/node.log 2>&1 &
	echo $! > $base/bin/otter.pid
	echo "cd to $current_path for continue"
	cd $current_path
else
	echo "otterNodeIdFile file("$otterNodeIdFile") OR log configration file($logback_configurationFile) is not exist,please create then first!"
fi
