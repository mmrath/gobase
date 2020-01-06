# base image
FROM node:lts

ARG APP

ENV APP=${APP}

# install chrome for protractor tests
RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
RUN sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google.list'
RUN apt-get update && apt-get install -yq google-chrome-stable

# set working directory
RUN mkdir -p /build/web
WORKDIR /build/web

# add `/app/node_modules/.bin` to $PATH
ENV PATH /build/web/node_modules/.bin:$PATH

# install and cache app dependencies
COPY . /build/web
RUN yarn



CMD ng serve $APP --host 0.0.0.0
