#!/bin/bash

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color
PADDING="    "

echo -e "\n${BLUE}====================[ Database Setup Initiated ]====================${NC}"
echo -e "${PADDING}${BLUE}PGUSER: ${YELLOW}$PGUSER${NC}, ${BLUE}PGDATABASE: ${YELLOW}$PGDATABASE${NC}, ${BLUE}POSTGRES_PASSWORD: ${YELLOW}$POSTGRES_PASSWORD.${NC}"
echo -e "${BLUE}====================================================================${NC}"
echo -e "${PADDING}\n"

# Check if the database already exists
if psql -v ON_ERROR_STOP=1 --username "$PGUSER" --dbname "$PGDATABASE" -lqt | cut -d \| -f 1 | grep -qw "$PGDATABASE"; then
    echo -e "${PADDING}${YELLOW}Database:${NC} [$PGDATABASE] already exists.${NC}\n"

    psql -v ON_ERROR_STOP=1 --username "$PGUSER" --dbname "$PGDATABASE" <<-EOSQL
        GRANT ALL PRIVILEGES ON DATABASE "$PGDATABASE" TO "$PGUSER";
EOSQL
    echo -e "${PADDING}${GREEN}All privileges granted to ${YELLOW}[$PGUSER]${NC} in ${YELLOW}[$PGDATABASE]${NC}\n"

else

    echo -e "${PADDING}${GREEN} Creating database.${NC}\n"
    echo -e "${PADDING}${RED}The given database ${YELLOW}[$PGDATABASE] does not exist.${NC}\n"

  psql -v ON_ERROR_STOP=1 --username "$PGUSER" --dbname "$PGDATABASE" <<-EOSQL
     CREATE DATABASE "$PGDATABASE";
     GRANT ALL PRIVILEGES ON DATABASE "$PGDATABASE" TO "$PGUSER";
EOSQL

  echo -e "${PADDING}${GREEN}The given database${NC} ${YELLOW}[$PGDATABASE]${NC} created successfully.\n"
  echo -e "${PADDING}${GREEN}All privileges granted to ${YELLOW}[$PGUSER]${NC} in ${YELLOW}[$PGDATABASE]${NC}\n"
fi

echo -e "${PADDING}"
echo -e "${BLUE}====================[ Database Setup Finished ]====================${NC}\n"
