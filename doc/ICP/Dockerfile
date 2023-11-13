FROM rust:1.71

RUN apt update && apt install -y curl libunwind-dev

RUN sh -ci "$(curl -fsSL https://internetcomputer.org/install.sh)"

EXPOSE 4939

CMD ["/usr/local/bin/dfx","start","--host","0.0.0.0:4939"]
