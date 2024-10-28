FROM scratch
ENV AWS_REGION="us-east-1"
# Se instala app
ADD main /
ADD .env /

# Se habilita puerto
EXPOSE 80

# Se ejecuta solo cuando se corre "docker run"
ENTRYPOINT ["/main"]