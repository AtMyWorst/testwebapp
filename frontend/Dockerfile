FROM node

ENV NUXT_ENV_BACKEND2_URL=https://staging.webml.app/backend2
ENV NUXT_ENV_BACKEND1_URL=https://staging.webml.app/backend

COPY . ./
RUN yarn install
RUN yarn build

CMD ["yarn", "start"]
