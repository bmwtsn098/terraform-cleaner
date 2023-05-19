FROM golang:1.19 as builder

COPY . /terraform-cleaner
WORKDIR /terraform-cleaner/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzz_terraform_cleaner.a && \
    clang -fsanitize=fuzzer fuzz_terraform_cleaner.a -o fuzz_terraform_cleaner.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /terraform-cleaner/mayhem/fuzz_terraform_cleaner.libfuzzer /