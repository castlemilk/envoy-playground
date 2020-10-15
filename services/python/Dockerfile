FROM python:3.8-slim
RUN mkdir /app
WORKDIR /app
ADD requirements.txt /app
RUN pip3 install -r requirements.txt
ADD . /app
EXPOSE 8080
# note --reload allows for hotloading
ENTRYPOINT ["gunicorn", "--reload", "--config", "gunicorn_config.py", "flaskr.wsgi:app"]