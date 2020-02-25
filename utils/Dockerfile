FROM telegraf
COPY telegraf.conf /etc/telegraf/telegraf.conf
COPY co2.test /tmp/co2-levels.log
COPY temp.test /tmp/ambient-temp.log
CMD ["telegraf"]
