FROM phusion/baseimage:18.04-1.0.0

RUN apt-get update && apt-get install -y apt-transport-https gnupg2 curl
RUN curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
RUN echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list
RUN apt-get update
RUN apt-get install -y kubectl

RUN mkdir /wheelhouse
ADD ./dist /wheelhouse
RUN chmod +x /wheelhouse/wheelhouse

WORKDIR /wheelhouse

CMD ["/wheelhouse/wheelhouse"]
