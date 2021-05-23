FROM scratch

ENV FILE=/opt/bexs/input-route.csv
ENV HOST=0.0.0.0
ENV PORT=3000

COPY storage/input-route.csv /opt/bexs/input-route.csv
COPY bexs /opt/bexs/bexs

ENTRYPOINT ["/opt/bexs/bexs"]
