FROM ubuntu:18.04 AS build

RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    curl -sL https://deb.nodesource.com/setup_14.x | bash && \
    apt-get install -y nodejs &&  \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install --production

COPY . .
RUN npm run build


FROM ubuntu:18.04

RUN apt-get update && \
    apt-get install -y curl ca-certificates && \
    curl -sL https://deb.nodesource.com/setup_14.x | bash && \
    apt-get install -y nodejs &&  \
    rm -rf /var/lib/apt/lists/*

RUN npm install -g http-server-spa

WORKDIR /app
COPY --from=build /app/build .
CMD ["http-server-spa", ".", "index.html", "80"]