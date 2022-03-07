FROM golang:1.17

WORKDIR $home\source\repos\bindicator

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get github.com/Jacobbrewer1/bindicator

RUN go build -a -v -work -o /bindicatorexe

CMD [ "/bindicatorexe" ]
