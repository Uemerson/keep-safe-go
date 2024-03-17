#!/bin/bash
function trap_ctrlc ()
{
    # perform cleanup here
    docker compose -f docker-compose.yml down
 
    # exit shell script with error code 2
    # if omitted, shell script will continue execution
    exit 2
}

trap "trap_ctrlc" 2
docker compose -f docker-compose.yml --env-file .env up -d --build --remove-orphans
docker compose -f docker-compose.yml logs -f --tail=15 api