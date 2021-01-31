FROM ubuntu:18.04
RUN apt-get update && \
    apt-get install gcc g++ make libgtk2.0-0 libgtk-3-0 libgbm-dev libnotify-dev libgconf-2-4 \
    systemd libnss3 libxss1 libasound2 libxtst6 xauth xvfb curl -y
RUN curl -fSsL https://deb.nodesource.com/setup_12.x | bash
RUN apt-get install nodejs -y
RUN npm install -g cypress http-server --unsafe-perm --silent
RUN cypress install
RUN cypress verify
COPY . /app
WORKDIR /app/public/cypress/videos
CMD ["http-server", "-p", "80", "-i", "false"]
