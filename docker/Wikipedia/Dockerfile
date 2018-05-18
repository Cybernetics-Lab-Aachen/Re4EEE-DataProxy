FROM ubuntu:16.04

RUN apt-get update && \
	apt-get upgrade -y && \
	apt-get install -y zip git wget curl nano mc

# Load a Wikipedia dump from https://dumps.wikimedia.org/enwiki/
RUN cd /home && \
    curl -o enwiki-multistream.xml.bz2 https://dumps.wikimedia.org/enwiki/20180501/enwiki-20180501-pages-articles-multistream.xml.bz2 && \
    curl -o enwiki-multistream-index.txt.bz2 https://dumps.wikimedia.org/enwiki/20180501/enwiki-20180501-pages-articles-multistream-index.txt.bz2

CMD ["echo"]