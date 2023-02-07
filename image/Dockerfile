FROM alpine as builder
RUN apk add curl
RUN curl -o direwolf20-installer https://api.modpacks.ch/public/modpack/101/2312/server/arm/linux
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "awscliv2.zip"
RUN unzip awscliv2.zip
COPY scripts scripts

FROM openjdk:18
WORKDIR /minecraft
COPY --from=builder direwolf20-installer direwolf20-installer
RUN chmod +x direwolf20-installer && \
    ./direwolf20-installer 101 2312 --nojava && \
    rm direwolf20-installer && \
    echo "eula=true" > eula.txt
RUN echo "-Xmx7500M" > user_jvm_args.txt
WORKDIR /
COPY --from=builder aws aws
RUN ./aws/install && rm -rf aws
COPY --from=builder scripts scripts
CMD ["scripts/wrapper.sh"]