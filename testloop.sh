

#!/bin/bash
set -e

for i in {1..50}
do 
echo "Test run $i"
go test -failfast -count 1 -run TestScripts
done