FROM envoyproxy/envoy:v1.14.1
COPY ./proxy/envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml
EXPOSE 8080