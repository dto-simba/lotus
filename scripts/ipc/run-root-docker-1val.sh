#!/usr/bin/env bash
set -e

if [ $# -ne 2 ]
then
    echo "Provide the port where lotus and the validator libp2p host will be listening as first and second arguments"
    echo "Args: [lotus_port] [validator_libp2p_port]"
    exit 1
fi

PORT=$1
VAL_PORT=$2

echo "[*] Running docker container for root in port $PORT"
img=`docker run -dit -p $PORT:1234 -p $VAL_PORT:1347 --name ipc_root_$PORT --entrypoint "/scripts/ipc/entrypoints/eudico-root-single.sh" eudico`
echo "[*] Waiting for the daemon to start"
docker exec -it $img  eudico wait-api --timeout 3500s
sleep 10
name=`docker ps --format "{{.Names}}" --filter "id=$img"`
echo ">>> Root daemon running in container: $img (friendly name: $name)"
token=`docker exec -it $img  eudico auth create-token --perm admin`
echo ">>> Token to /root daemon: $token"
wallet=`docker exec -it $img  eudico wallet default`
echo ">>> Default wallet: $wallet"
