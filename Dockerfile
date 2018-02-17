FROM horseville/core
# runner image
FROM xena/alpine
COPY --from=0 /root/go/src/github.com/horseville/horseville/bin/ /usr/local/bin/
CMD /usr/local/bin/horsevilled
