FROM docker-mirror.artifactory.at.ivi.ru/postgres-ru:15.6
RUN  /bin/ln -fs /usr/share/zoneinfo/Europe/Moscow /etc/localtime
ADD docker-entrypoint-initdb.d/. /docker-entrypoint-initdb.d/