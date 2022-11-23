FROM bitnami/elasticsearch-curator:5.8.1

USER root

RUN install_packages cron && \
    echo \
        '* * * * *' \
        root \
        LC_ALL=C.UTF-8 LANG=C.UTF-8 \
        /opt/bitnami/python/bin/curator \
        --config=/usr/share/curator/config/curator.yml \
        /usr/share/curator/config/delete_log_files_curator.yml \
        '>/proc/1/fd/1' '2>/proc/1/fd/2' \
        >>/etc/crontab

ENTRYPOINT ["cron"]
.md["-f", "-L8"]
