FROM scratch

COPY httpechoes /httpechoes

EXPOSE 8080

ENTRYPOINT ["/httpechoes"]
