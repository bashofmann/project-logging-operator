FROM registry.suse.com/bci/bci-micro:15.3
RUN echo 'project-logging-operator:x:1000:1000::/home/project-logging-operator:/bin/bash' >> /etc/passwd && \
    echo 'project-logging-operator:x:1000:' >> /etc/group && \
    mkdir /home/project-logging-operator && \
    chown -R project-logging-operator:project-logging-operator /home/project-logging-operator
COPY bin/project-logging-operator /usr/bin/
USER project-logging-operator
CMD ["project-logging-operator"]
