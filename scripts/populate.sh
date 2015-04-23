#!/usr/bin/env bash

# create genre
horror=`curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Horror</Name></Actor>' http://localhost:8080/genre | grep ID`
family=`curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Family</Name></Actor>' http://localhost:8080/genre | grep ID`
romcom=`curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Romantic Comedy</Name></Actor>' http://localhost:8080/genre | grep ID`


# create actors
bridges=`curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Jeff Bridges</Name></Actor>' http://localhost:8080/actors | grep ID`
goldblum=`curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Jeff Goldblum</Name></Actor>' http://localhost:8080/actors | grep ID`
garlin=`curl -X POST -H 'Content-type: application/xml' -d '<Actor><Name>Jeff Garlin</Name></Actor>' http://localhost:8080/actors | grep ID`

horrorId=${horror:9:-2}
familyId=${family:9:-2}
romcomId=${romcom:9:-2}
bridgesId=${bridges:9:-2}
goldblumId=${goldblum:9:-2}
garlinId=${garlin:9:-2}

set -x

# create movies
curl -X POST -H 'Content-type: application/xml'  -d "<Movie><Title>Air Bud</Title><Released>August 1, 1997</Released><Description>Kid invents magic dog, terrorizes classmates.</Description><GenreID>$familyId</GenreID><CastIDs><CastID>$bridgesId</CastID><CastID>$goldblumId</CastID></CastIDs></Movie>" http://localhost:8080/movies
curl -X POST -H 'Content-type: application/xml'  -d "<Movie><Title>The Double</Title><Released>April 4, 2014</Released><Description>Sad guy from Facebook movie makes you feel depressed.</Description><GenreID>$romcomId</GenreID><CastIDs><CastID>$garlinId</CastID><CastID>$goldblumId</CastID></CastIDs></Movie>" http://localhost:8080/movies
curl -X POST -H 'Content-type: application/xml'  -d "<Movie><Title>Air Bud: Golden Receiver</Title><Released>August 14, 1998</Released><Description>Everyone&apos;s favorite wizard is back for more action under the Friday Night Lights (trademark).</Description><GenreID>$horrorId</GenreID><CastIDs><CastID>$bridgesId</CastID></CastIDs></Movie>" http://localhost:8080/movies
curl -X POST -H 'Content-type: application/xml'  -d "<Movie><Title>Jurassic Park</Title><Released>June 11, 1993</Released><Description>God creates Dinosaurs. God destroys Dinosaurs. God Creates Man. Man destroys God. Man creates Dinosaurs.</Description><GenreID>$romcomId</GenreID><CastIDs><CastID>$goldblumId</CastID></CastIDs></Movie>" http://localhost:8080/movies
curl -X POST -H 'Content-type: application/xml'  -d "<Movie><Title>Air Bud X: Wrath of Bud</Title><Released>May 6, 2015</Released><Description>Air Bud does it again in this whimsical take on Death of a Salesman.</Description><GenreID>$familyId</GenreID><CastIDs><CastID>$bridgesId</CastID><CastID>$garlinId</CastID></CastIDs></Movie>" http://localhost:8080/movies
