FROM alpine:3.4

RUN apk --no-cache add python py-pip
RUN pip install Flask

COPY . /app

WORKDIR "/app"
CMD ["python", "app.py"]