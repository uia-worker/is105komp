# is105komp
Kompendiet (utgått)

# Oversettelsesbuffer eller buffer for oversettelser
Siden adressene i det virtuelle minne må forløpende oversettes til adressene i maskinvaren, er det rom for optimalisering, hvis noen av disse oversettelsene (mappingene VA0xOEF0 -> MA0x0AAB) blir mellomlagret, slik at de kan hentes fra dette mellomlageret som på engelsk blir kalt "translation lookaside buffer" eller TLB. 

TLB blir også kalt for adresse-oversettelses cache. Den er en del av minne-administrasjons enhet, som har direkte forbindelse med (eller er en del av) den sentrale prosesseringsenheten, - CPU (eng. Central Processing Unit). Den kan være pakket i den samme boksen som CPU. 

CPU ---> MMU ---> bus (mellom CPU og RAM)
         ^ |				|
         | |                |
         | v                v
         TLB               RAM

Det er virtuell adresse som sendes fra CPU register til MMU, som oversetter adressen til den fysiske adressen og mellomlagrer denne "mappingen" i TBL. 

MMU sporer minne i blokker (eng. pages).

Se bilde for worklow 

Steps_In_a_Translation_Lookaside_Buffer.png (PNG Image, 592 × 531 pixels). (2016, October 06). Retrieved from https://upload.wikimedia.org/wikipedia/commons/c/c1/Steps_In_a_Translation_Lookaside_Buffer.png

qemu-system-x86_64 -m 2048 -boot d -cdrom path_to_android_x86.iso -hda android_disk.img
eller
Anbox (Android in a Box)
