FROM ubuntu:bionic

WORKDIR /app

ADD . /app

RUN mkdir /var/run/sshd && \
    apt-get update && \
    apt-get install -y \
    openssh-server && \
    sed -i 's/PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config && \
    sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd

RUN echo 'root:MyPass' | chpasswd &&
ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

#RUN cat /app/id_rsa.pub | ssh user@$( hostname ) 'cat >> .ssh/authorized_keys'

EXPOSE 22

CMD ["/usr/sbin/sshd", "-D"]
