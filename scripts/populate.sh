#!/usr/bin/env bash
set -x

# create genre
curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Horror</Name></Actor>' http://localhost:8080/genre
curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Family</Name></Actor>' http://localhost:8080/genre
curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Romantic Comedy</Name></Actor>' http://localhost:8080/genre

# create actors
curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Jeff Bridges</Name></Actor>' http://localhost:8080/actors
curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Jeff Goldblum</Name></Actor>' http://localhost:8080/actors
curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Jeff Garlin</Name></Actor>' http://localhost:8080/actors

# create movies
curl -X POST -H 'Content-type: application/xml'  -d '<Movie><Title>Air Bud</Title><Released>August 1, 1997</Released><Description>Kid invents magic dog, terrorizes classmates.</Description></Movie>' http://localhost:8080/movies
curl -X POST -H 'Content-type: application/xml'  -d '<Movie><Title>The Double</Title><Released>April 4, 2014</Released><Description>Sad guy from Facebook movie makes you feel depressed.</Description></Movie>' http://localhost:8080/movies
curl -X POST -H 'Content-type: application/xml'  -d '<Movie><Title>Air Bud: Golden Receiver</Title><Released>August 14, 1998</Released><Description>Everyone&apos;s favorite wizard is back for more action under the Friday Night Lights (trademark).</Description></Movie>' http://localhost:8080/movies
curl -X POST -H 'Content-type: application/xml'  -d '<Movie><Title>Jurassic Park</Title><Released>June 11, 1993</Released><Description>God creates Dinosaurs. God destroys Dinosaurs. God Creates Man. Man destroys God. Man creates Dinosaurs.</Description></Movie>' http://localhost:8080/movies
curl -X POST -H 'Content-type: application/xml'  -d '<Movie><Title>Air Bud X: Wrath of Bud</Title><Released>May 6, 2015</Released><Description>Air Bud does it again in this whimsical take on Death of a Salesman.</Description></Movie>' http://localhost:8080/movies
