# devcontainer-video

Dette eksempelet bygger på https://www.youtube.com/watch?v=SDa3v4Quj7Y
Veldig bra.
 
* Installer Remote Development extension pack
* Opprett repo før du bruker visual studio code
* Åpne VS Code
* Åpne kommandopalletten (Ctrl+Shift+P) og velg
* devcontainer:clone repository in container volume
* velg configfile: default linux unviversal (Maskin den er stor)
  Den inneholder flere utlvikler verktøy som er nyttige for utvikling

  # nuts and bolts
  Noen ganger hvis git repo er satt opp med signering av commits vil devcontainer ikke finne sertifikate som er registert på host maskinen.
  Brukt da
  ```
    git config commit.gpgsign false
  ```
  
