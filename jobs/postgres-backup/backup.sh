#!/bin/bash

export DUMP_FILE=/backup_${PRODUCT_NAME}_`date +%Y%m%d_%H%M%S`.pgdump
PGPASSWORD=$POSTGRES_PASSWORD pg_dump -d $POSTGRES_DB -U $POSTGRES_USER -h $POSTGRES_HOST -f $DUMP_FILE
bzip2 $DUMP_FILE
aws s3 cp ${DUMP_FILE}.bz2 $S3_BACKUP_PATH
