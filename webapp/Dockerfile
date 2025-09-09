FROM node:20
WORKDIR /app
COPY package.json package-lock.json* pnpm-lock.yaml* yarn.lock* ./
RUN npm install
COPY . .
EXPOSE 5173
ENV VITE_API_BASE=http://localhost:8080
CMD ["npm","run","dev","--","--host","0.0.0.0"]
