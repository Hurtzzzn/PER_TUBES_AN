#ifndef TUBES_H_INCLUDED
#define TUBES_H_INCLUDED

#include <iostream>
#include <conio.h>
#include <stdlib.h>

using namespace std;

struct Penyewa {
    string id_penyewa;
    string nama;
    int nik;
};

struct Buku {
    string id_buku, judul, penulis;
    int harga_sewa;
};

typedef Penyewa infotypePenyewa;

typedef Buku infotypeBuku;

typedef struct elmPenyewa *adrPenyewa;

typedef struct elmBuku *adrBuku;

typedef struct elmRelasi *adrRelasi;

struct elmPenyewa {
    infotypePenyewa info;
    adrPenyewa next;
};

struct elmBuku {
    infotypeBuku info;
    adrBuku next;
    adrBuku prev;
};

struct elmRelasi {
    adrPenyewa nextPenyewa;
    adrBuku nextBuku;
    adrRelasi nextRelasi;
};

struct ListPenyewa {
    adrPenyewa First;
};

struct ListBuku {
    adrBuku First;
    adrBuku Last;
};

struct ListRelasi {
    adrRelasi First;
};


void createListPenyewa(ListPenyewa &LP);

void createListBuku(ListBuku &LB);

void createListRelasi(ListRelasi &LR);

adrPenyewa createElmPenyewa(infotypePenyewa dataP);

adrBuku createElmBuku(infotypeBuku dataB);

adrRelasi createElmRelasi();

void insertLastPenyewa(ListPenyewa &LP, adrPenyewa P);

void insertFirstBuku(ListBuku &LB, adrBuku B);

void addRelasi(ListRelasi &LR, ListPenyewa LP, ListBuku LB, string id_penyewa, string id_buku);

void deletePenyewa(ListPenyewa &LP, ListRelasi &LR, string id_penyewa, adrPenyewa &P, adrRelasi &R);

void deleteBuku(ListBuku &LB, ListRelasi &LR, string id_buku, adrBuku &B, adrRelasi &R);

void deleteRelasi(ListRelasi &LR, string id_penyewa, string id_buku, adrRelasi &R);

adrPenyewa findPenyewa(ListPenyewa LP, string id_penyewa);

adrBuku findBuku(ListBuku LB, string id_buku);

adrRelasi findRelasiPenyewaAndBuku(ListRelasi LR, string id_penyewa, string id_buku);

void showAllPenyewa(ListPenyewa LP);

void showAllBuku(ListBuku LB);

void showBukuFromPenyewa(ListRelasi LR, ListPenyewa LP, string id_penyewa);

void showPenyewaFromBuku(ListRelasi LR, ListBuku LB, string id_buku);

void showAllPenyewaWithBuku(ListRelasi LR, ListPenyewa LP);

void showwAllBukuWithPenyewa(ListRelasi LR, ListBuku LB);

void countBukuDariPenyewa(ListRelasi LR, ListPenyewa LP, string id_penyewa);

void countPenyewaDariBuku(ListRelasi LR, ListBuku LB, string id_buku);

void countBukuWithNoPenyewa(ListRelasi LR, ListBuku LB);

void countPenyewaWithNoBuku(ListRelasi LR, ListPenyewa LP);

void editBukuDariPenyewa(ListRelasi &LR, ListBuku LB, string id_penyewa, string id_buku, string newId_Buku);

void editPenyewaDariBuku(ListRelasi &LR, ListPenyewa LP, string id_buku, string id_penyewa, string newId_Penyewa);

void insertLastRelasi(ListRelasi &LR, adrRelasi R);

void deleteFirstPenyewa(ListPenyewa &LP, adrPenyewa &P);

void deleteLastPenyewa(ListPenyewa &LP, adrPenyewa &P);

void deleteAfterPenyewa(ListPenyewa &LP, adrPenyewa prec, adrPenyewa &P);

void deleteFirstBuku(ListBuku &LB, adrBuku &B);

void deleteLastBuku(ListBuku &LB, adrBuku &B);

void deleteAfterBuku(ListBuku &LB, adrBuku prec, adrBuku &B);

void deleteFirstRelasi(ListRelasi &LR, adrRelasi &R);

void deleteLastRelasi(ListRelasi &LR, adrRelasi &R);

void deleteAfterRelasi(ListRelasi &LR, adrRelasi prec, adrRelasi &R);

void showOneDataPenyewa(adrPenyewa P);

void showOneDataBuku(adrBuku B);

void showOneDataRelasi(adrRelasi R);

void validIdPenyewa(ListPenyewa LP, string &id_penyewa);

void validIdBuku(ListBuku LB, string &id_buku);

int menuUtama();

int menuPenyewa();

int menuBuku();

int menuRelasi();





#endif // TUBES_H_INCLUDED
