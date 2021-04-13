# Projekat kulturna dobra - NTP

<img alt="Pharo" src="https://img.shields.io/badge/pharo%20-%234f0599.svg?&style=for-the-badge"/> <img alt="Go" src="https://img.shields.io/badge/go-%2300ADD8.svg?&style=for-the-badge&logo=go&logoColor=white"/> <img alt="Firebase" src="https://img.shields.io/badge/firebase%20-%23039BE5.svg?&style=for-the-badge&logo=firebase"/>


## Opis aplikacije
- Sistem omogucuje neulogovanim korisnicima da vide kutlurna dobra koja se nalaze u ponudi sistema. 
- Korisnici mogu da pretrazuju i sortiraju kulturna dobra po raznim kriterijumima kao sto su naziv, lokacija (drzava, grad, ulica),prosecna ocena(rejting), broj recenzija (komentara), tip kulturnog dobra, broju lajkova, broju dislajkova.
- Sistem omogucuje neulogovanim korisnicima da lajkuju i dislajkuju svako kulturno dobro koje se nalazi u ponudi sistema.
- Komentari za recenzije od strane neautentifikovanog korisnika
- ocenjivanje recenzije da li je korisna
- ocenjivanje korisnika da li je azuran ili nije

## Arhitektura sistema
Sistem se sastoji iz tri tehnologije: Pharo, Golang i Firebase.

### Pharo
- GUI sistema 
- Omogucava graficki prikaz, sortiranje, filtriranje, lajkovanje i dislajkovanje kulturnih dobara.

### Golang
- Bekend funkcionalnosti
- REST servisi koji komuniciraju sa klijentom, odnosno Pharo
- U ovom servisu je implementirana sva logika sistema
- Komunicira sa Firestore bazom preko Firebase Admin SDK 

### Firebase
- Koristi se Firestore nosql baza podataka
- Ova baza je inicijalno popunjena podacima i podaci se ucitavaju pri pokretanju aplikacije
