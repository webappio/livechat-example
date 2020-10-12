FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    curl -sL https://deb.nodesource.com/setup_14.x | bash && \
    apt-get install -y nodejs &&  \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install --production

COPY . .
ENV REACT_APP_BROWSER=none 
EXPOSE 80
ENTRYPOINT ["./node_modules/.bin/react-scripts", "start"]
