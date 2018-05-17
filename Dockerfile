FROM ubuntu:16.04
ENV GOPATH /go

# Update the operating system and install base tools:
RUN apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y zip git wget curl && \
	# Create the Go workspace:
	mkdir /go && \
    mkdir /go/src && \
    mkdir /go/bin && \
    mkdir /go/pkg && \
    cd /go && \
	wget --no-check-certificate -O go.tar.gz https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz && \
	tar -C /usr/local -xzf go.tar.gz && \
	rm go.tar.gz && \
    /usr/local/go/bin/go get -u github.com/golang/dep/cmd/dep

# Insert all files from the repo (from the current directory, not from Git):
ADD . /go/src/github.com/Cybernetics-Lab-Aachen/Re4EEE-DataProxy/

# Compile and Setup
RUN export PATH=$PATH:/usr/local/go/bin:/go/bin && \
	cd /go/src/github.com/Cybernetics-Lab-Aachen/Re4EEE-DataProxy && \
	# Compile:
	dep ensure && \
    go install && \
	# Copy the final binary and the runtime scripts to the home folder:
	cp /go/bin/Re4EEE-DataProxy /home && \
	cp /go/src/github.com/Cybernetics-Lab-Aachen/Re4EEE-DataProxy/run.sh /home/run.sh && \
	# Make the scripts executable:
	chmod 0777 /home/run.sh && \
	chmod 0777 /home/Re4EEE-DataProxy

ENV Re4EEEDataProxy_TwitterDBHostname="tweet-collector-elearning-db"
ENV Re4EEEDataProxy_TwitterDBDatabaseName="TweetCollector"
ENV Re4EEEDataProxy_TwitterDBUsername="TweetCollector"
ENV Re4EEEDataProxy_TwitterDBPassword="PASSWORD"
ENV Re4EEEDataProxy_TwitterDBCollectionName="Tweets"
ENV Re4EEEDataProxy_ServerIfacePort="0.0.0.0:80"

# Run anything below as nobody:
USER nobody

# TweetCollector provides the admin interface on port 50000. There is no public interface!
EXPOSE 50000

# Define the working directory:
WORKDIR /home

# The default command to run, if a container starts:
CMD ["./run.sh"]