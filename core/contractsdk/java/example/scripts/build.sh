for contract in `ls`;do   [[ -f $contract/pom.xml ]] && mvn -f $contract/pom.xml package;done