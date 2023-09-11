FROM python:3.11-alpine3.18

COPY . /app

WORKDIR /app

RUN chmod +x ./entrypoint.sh && \
    apk add git && \
    pip install cookiecutter
    
ENTRYPOINT ["./entrypoint.sh"]